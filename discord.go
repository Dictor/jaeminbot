package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/dictor/jaeminbot/ent"
	"github.com/sirupsen/logrus"
)

type (
	editMessageContext struct {
		OriginalMessage *discordgo.MessageCreate
		TriggerMessage  *discordgo.Message
		Keyword         string
	}

	vmMessageContext struct {
		Session *discordgo.Session
		Message *discordgo.MessageCreate
		Command *ent.Command
	}

	UserNotFound error
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
		Title:       ":dizzy_face: 이런! 오류가 발생했어요 ㅠㅠ",
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
				{Name: "사용법", Value: "예를 들어 `안녕`이라는 명령어를 추가해볼까요? 첫번째로, `재민쿤 명령어 추가 안녕`을 입력하여 명령어를 추가해주세요. 두번째로, `재민쿤 명령어 코드 안녕`을 입력해 `안녕` 명령어를 호출하면 실행할 코드를 입력해주세요. 이때 코드는 아래 사진과 같이 제 메세지에 답장 버튼을 누르고 입력해주셔야해요. 마지막으로, `재민쿤 명령어 저장`을 입력하면 명령어 등록이 끝나요. 명령어를 사용하려면 `재민쿤 안녕`이라고 불러주시면 되요.", Inline: false},
				{Name: "명령어", Value: "- `재민쿤 명령어 추가 <명령어>` : 명령어 추가\n -`재민쿤 명령어 코드 <명령어>` : 코드 작성\n - `재민쿤 명령어 저장 <python 또는 javascript>` : 명령어 저장", Inline: false},
			},
			Description: "안녕하세요! 저는 어떤 코드라도 척척 실행해내는 석박사통합과정생 김재민이에요. ",
		}))
		return
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
			return
		case "오류":
			logDiscordSendResult(s.ChannelMessageSend(m.ChannelID, "오류 확인 기능은 아직 준비중이에요!"))
			return
		}
	case 3:
		switch msg[1] {
		case "명령어":
			switch msg[2] {
			case "추가":
				logDiscordSendResult(s.ChannelMessageSend(m.ChannelID, "명령어 추가는 `재민쿤 명령어 추가 <명령어 호출 키워드>`로 가능합니다."))
				return
			case "저장":
				logDiscordSendResult(s.ChannelMessageSend(m.ChannelID, "명령어 저장은 코드를 답글로 입력한 후 `재민쿤 명령어 저장 <javascript 또는 python>`로 가능합니다."))
				return
			}
		}
	case 4:
		switch msg[1] {
		case "명령어":
			switch msg[2] {
			case "추가":
				cmd, err := Client.Command.Create().
					SetID(fmt.Sprintf("%s-%s", msg[3], m.GuildID)).
					SetCreator(m.Author.ID).
					SetKeyword(msg[3]).
					SetCode("").
					SetServer(m.GuildID).
					Save(ClientContext)
				if err != nil {
					discordErrorHandler(s, m, err)
					return
				}
				logDiscordSendResult(s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("`%s` 호출 키워드의 명령어가 `%s` 사용자에 의해 추가되었습니다. 명령어 코드 작성을 위해서는 `재민쿤 명령어 코드 %s` 명령어를 사용해주세요.", cmd.Keyword, cmd.Creator, cmd.Keyword)))
				return
			case "코드":
				_, cmd, err := getCommandByKeyword(msg[3], m.GuildID)
				if err != nil {
					if _, ok := err.(*ent.NotFoundError); ok {
						logDiscordSendResult(s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("`%s` 호출 키워드의 명령어를 찾을 수 없습니다!", msg[3])))
					} else {
						discordErrorHandler(s, m, err)
					}
					return
				}

				trigMsg, err := s.ChannelMessageSend(m.ChannelID, "이 메세지의 답장으로 수정할 코드를 보내고 `재민쿤 명령어 저장`을 호출해주세요")
				logDiscordSendResult(trigMsg, err)
				logDiscordSendResult(s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
					Type:        discordgo.EmbedTypeArticle,
					Title:       fmt.Sprintf("`%s` 명령어의 코드", msg[3]),
					Description: fmt.Sprintf("```\n%s\n```", cmd.Code),
				}))
				if err != nil {
					discordErrorHandler(s, m, err)
				}
				editMessage[m.Author.ID] = editMessageContext{
					OriginalMessage: m,
					TriggerMessage:  trigMsg,
					Keyword:         msg[3],
				}
				return
			case "저장":
				editCtx, ok := editMessage[m.Author.ID]
				if !ok {
					logDiscordSendResult(s.ChannelMessageSend(m.ChannelID, "편집 중인 명령어가 없습니다!"))
					return
				}
				exist, cmd, err := getCommandByKeyword(editCtx.Keyword, m.GuildID)
				if !exist {
					logDiscordSendResult(s.ChannelMessageSend(m.ChannelID, "어라? 편집 중이던 명령어를 찾을 수 없습니다!"))
					discordErrorHandler(s, m, err)
					delete(editMessage, m.Author.ID)
					return
				}
				rmsg, err := findCodeReply(s, m, editCtx)
				if err != nil {
					discordErrorHandler(s, m, err)
					return
				}

				if msg[3] != "javascript" && msg[3] != "python" {
					logDiscordSendResult(s.ChannelMessageSend(m.ChannelID, "언어는 `python` 또는 `javascript`여야 합니다."))
					return
				}
				newCmd, err := cmd.Update().SetCode(rmsg.Content).SetLanguage(msg[3]).Save(ClientContext)
				if err != nil {
					discordErrorHandler(s, m, err)
					return
				}
				logDiscordSendResult(s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%d 바이트의 용량의 코드로 `%s` 명령어를 업데이트 했습니다.", len(newCmd.Code), newCmd.Keyword)))
				return
			}
		}
	}
	exist, cmd, _ := getCommandByKeyword(msg[1], m.GuildID)
	var err error
	if exist {
		arg := msg[2:]
		switch cmd.Language {
		case "javascript":
			err = runJavascriptCode(vmMessageContext{
				Session: s,
				Message: m,
				Command: cmd,
			}, arg)
		case "python":
			err = runPythonCode(vmMessageContext{
				Session: s,
				Message: m,
				Command: cmd,
			}, arg)
		}
		if err != nil {
			discordErrorHandler(s, m, err)
		}
	} else {
		logDiscordSendResult(s.ChannelMessageSend(m.ChannelID, "등록되지 않은 명령어에요!:sob:"))
	}
}

func findCodeReply(s *discordgo.Session, m *discordgo.MessageCreate, ctx editMessageContext) (*discordgo.Message, error) {
	msgs, err := s.ChannelMessages(m.ChannelID, 100, "", ctx.TriggerMessage.ID, "")
	if err != nil {
		return nil, err
	}

	me, err := s.User("@me")
	if err != nil {
		return nil, err
	}
	for _, m := range msgs {
		if m.Author.ID == ctx.OriginalMessage.Author.ID && m.Type == discordgo.MessageTypeReply && m.Mentions[0].ID == me.ID {
			return m, nil
		}
	}
	return nil, fmt.Errorf("답장 메세지를 찾을 수 없습니다")
}
