package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/dictor/jaeminbot/ent"
	"github.com/dictor/jaeminbot/ent/command"
	"github.com/sirupsen/logrus"
)

type (
	editMessageContext struct {
		Message *discordgo.Message
		Keyword string
	}

	vmMessageContext struct {
		Session *discordgo.Session
		Message *discordgo.MessageCreate
		Command *ent.Command
	}
)

var (
	currentSession *discordgo.Session
	editMessage    map[string]editMessageContext = map[string]editMessageContext{}
)

func startDiscordBot(botToken string) error {
	var err error = nil

	currentSession, err = discordgo.New("Bot " + botToken)
	if err != nil {
		return err
	}

	currentSession.AddHandler(discordMessageHandler)
	err = currentSession.Open()
	return err
}

func logDiscordSendResult(msg *discordgo.Message, err error) {
	if err != nil {
		Logger.WithError(err).Errorln("fail to send message")
		return
	}
	Logger.WithFields(logrus.Fields{
		"content": msg.Content,
		"channel": msg.ChannelID,
	}).Infoln("message sent")
}

func discordErrorHandler(s *discordgo.Session, m *discordgo.MessageCreate, err error) {
	Logger.WithError(err).Error("error caused during processing message")
	logDiscordSendResult(s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Type:        discordgo.EmbedTypeArticle,
		Title:       "이런! 오류가 발생했어요 ㅠㅠ",
		Description: fmt.Sprintf("`%s`", err),
	}))
}

func vmMessageSender(ctx vmMessageContext, msg string) {
	logDiscordSendResult(ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, msg))
}

/*
재민쿤 : help 1
재민쿤 명령어 : 명령어 목록 2
재민쿤 명령어 추가 3
재민쿤 명령어 추가 <keyword> 4
재민쿤 명령어 코드 <keyword> 4
재민쿤 <명령어> 2
재민쿤 <명령어> [인수1] [인수2].... 2+n
재민쿤 오류 : 오류 목록 2
재민쿤 오류 <오류 번호> 3
*/
func discordMessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID { // Ignore bot's itself message
		return
	}

	if strings.HasPrefix(m.Content, "재민쿤") {
		Logger.WithFields(logrus.Fields{
			"server_id":   m.ChannelID,
			"author_id":   m.Author.ID,
			"author_name": m.Author.Username,
			"content":     m.Content,
		}).Infoln("bot is called")
	} else {
		return
	}

	msg := strings.Split(m.Content, " ")
	switch len(msg) {
	case 1:
		logDiscordSendResult(s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Type:  discordgo.EmbedTypeRich,
			Title: "재민쿤 명령어",
			Fields: []*discordgo.MessageEmbedField{
				{Name: "도움말", Value: "`재민쿤`", Inline: true},
				{Name: "명령어 목록", Value: "`재민쿤 명령어`", Inline: true},
				{Name: "명령어 추가", Value: "`재민쿤 명령어 추가`", Inline: true},
				{Name: "명령어 정보", Value: "`재민쿤 <명령어>`", Inline: true},
				{Name: "명령어 실행", Value: "`재민쿤 <명령어> [인수1] [인수2]...`", Inline: true},
			},
			Description: "원하는 기능의 명령어를 채팅방에 입력하면 됩니다. (<>는 필수, []는 선택 입력 항목입니다.)",
		}))
	case 2:
		switch msg[1] {
		case "명령어":
			commands, err := Client.Command.Query().All(ClientContext)
			if err != nil {
				discordErrorHandler(s, m, err)
				return
			}
			detail := ""
			for i, c := range commands {
				if i != 0 {
					detail += ","
				}
				detail += fmt.Sprintf("`%s`", c.Keyword)
			}
			logDiscordSendResult(s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
				Type:        discordgo.EmbedTypeArticle,
				Title:       fmt.Sprintf("총 %d개의 명령어가 있어요.", len(commands)),
				Description: detail,
			}))
		case "오류":
			logDiscordSendResult(s.ChannelMessageSend(m.ChannelID, "오류 확인 기능은 아직 준비중이에요!"))
		default:
			logDiscordSendResult(s.ChannelMessageSend(m.ChannelID, "이해할 수 없는 명령어에요!"))
		}
	case 3:
		switch msg[1] {
		case "명령어":
			switch msg[2] {
			case "추가":
				logDiscordSendResult(s.ChannelMessageSend(m.ChannelID, "명령어 추가는 `재민쿤 명령어 추가 <명령어 호출 키워드>`로 가능합니다."))
			case "저장":
				editCtx, ok := editMessage[m.Author.ID]
				if !ok {
					logDiscordSendResult(s.ChannelMessageSend(m.ChannelID, "편집 중인 명령어가 없습니다!"))
					return
				}
				exist, cmd, err := getCommandByKeyword(editCtx.Keyword)
				if !exist {
					logDiscordSendResult(s.ChannelMessageSend(m.ChannelID, "어라? 편집중이던 명령어를 찾을 수 없습니다!"))
					discordErrorHandler(s, m, err)
					delete(editMessage, m.Author.ID)
					return
				}
				cmd, err = cmd.Update().SetCode(editCtx.Message.Content).Save(ClientContext)
				if err != nil {
					discordErrorHandler(s, m, err)
				}
				logDiscordSendResult(s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%d 바이트의 용량의 코드로 `%s` 명령어를 업데이트 했습니다.", len(cmd.Code), cmd.Keyword)))
			}
		}
	case 4:
		switch msg[1] {
		case "명령어":
			switch msg[2] {
			case "추가":
				cmd, err := Client.Command.Create().
					SetCreator(m.Author.ID).
					SetKeyword(msg[3]).
					SetCode("").
					Save(ClientContext)
				if err != nil {
					discordErrorHandler(s, m, err)
					return
				}
				logDiscordSendResult(s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("`%s` 키워드를 가진 명령어가 `%s` 사용자에 의해 추가되었습니다. 명령어 코드 작성을 위해서는 `재민쿤 명령어 코드` 명령어를 사용해주세요.", cmd.Keyword, cmd.Creator)))
			case "코드":
				cmd, err := Client.Command.Query().Where(command.Keyword(msg[3])).Only(ClientContext)
				if err != nil {
					if _, ok := err.(*ent.NotFoundError); ok {
						logDiscordSendResult(s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("`%s` 명령어를 찾을 수 없습니다!", msg[3])))
					} else {
						discordErrorHandler(s, m, err)
						return
					}
				}
				logDiscordSendResult(s.ChannelMessageSend(m.ChannelID, "제가 보낸 아래 코드 메세지를 수정하고 `재민쿤 명령어 저장`을 호출해주세요"))
				editMsg, err := s.ChannelMessageSend(m.ChannelID, cmd.Code)
				if err != nil {
					discordErrorHandler(s, m, err)
				}
				editMessage[m.Author.ID] = editMessageContext{
					Message: editMsg,
					Keyword: msg[3],
				}
			}
		}
	default:
		exist, cmd, err := getCommandByKeyword(msg[1])
		if exist {
			runCode(vmMessageContext{
				Session: s,
				Message: m,
				Command: cmd,
			})
		} else {
			discordErrorHandler(s, m, err)
		}
	}
}
