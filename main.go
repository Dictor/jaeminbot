package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/dictor/jaeminbot/ent"
	_ "github.com/mattn/go-sqlite3"
	"github.com/namsral/flag"
	"github.com/sirupsen/logrus"
)

var (
	Logger                     *logrus.Logger
	gitTag, gitHash, buildDate string // build flags
	Client                     *ent.Client
	ClientContext              context.Context
	ClientContextCancel        context.CancelFunc
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
	flag.StringVar(&storePath, "store", "db.db", "Commit store file's path")
	flag.Parse()

	mustNoError("starting bot", startDiscordBot(botToken))

	Client, err := ent.Open("sqlite3", fmt.Sprintf("file:%s?_fk=1", storePath))
	mustNoError("opening db", err)
	ClientContext, ClientContextCancel = context.WithCancel(context.Background())
	defer Client.Close()
	mustNoError("migrationing db", Client.Schema.Create(context.Background()))

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
