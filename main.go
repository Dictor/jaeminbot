package main

import (
	"github.com/namsral/flag"
	"github.com/sirupsen/logrus"
)

var(
	Logger *logrus.Logger
	gitTag, gitHash, buildDate string // build flags
)

func mustNoError(action string, err error) {
	if err != nil {
		Logger.WithFields(logrus.Fields{
			"action": action,
			"err":    err,
		}).Panic("necessary action failed")
	}
}

func main() {
	var (
		botToken, storePath string
	)
	Logger = logrus.New()
	Logger.Infof("jaeminbot %s (%s) - %s", gitTag, gitHash, buildDate)


	flag.StringVar(&botToken, "token", "", "Bot's Token string")
	flag.StringVar(&storePath, "store", "./db.db", "Commit store file's path")
	flag.Parse()

	mustNoError("starting bot",startDiscordBot(botToken))
}