package main

import (
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

	currentSession.AddHandler(messageHandler)
	err = currentSession.Open()
	return err
}

func discordMessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID { // Ignore bot's itself message
		return
	}
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
