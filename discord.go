package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

var currentSession *discordgo.Session

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

/*
재민쿤 : help 1
재민쿤 명령어 : 명령어 목록 2
재민쿤 명령어 추가 3
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
			Client.
		case "오류":
		default:
		}
	case 3:
	}
}
