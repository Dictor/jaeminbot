package main

import (
	"github.com/dictor/jaeminbot/ent"
	"github.com/dictor/jaeminbot/ent/command"
)

func getCommandByKeyword(keyword string, serverId string) (exist bool, cmd *ent.Command, err error) {
	c, err := Client.Command.Query().
		Where(command.Keyword(keyword), command.Server(serverId)).
		Only(ClientContext)
	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			return false, nil, err
		} else {
			return false, nil, err
		}
	}
	return true, c, nil
}
