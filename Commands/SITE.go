package Commands

import (
	"FTPserver/Commands/SiteCommands"
	"FTPserver/Configuration"
	"FTPserver/Connection"
	"FTPserver/Replies"
	"strings"
)

type SITE struct {
	cs     *Connection.Status
	config Configuration.FTPConfig
}

func (cmd SITE) Execute(argString string) Replies.FTPReply {
	args := strings.Split(argString, " ")
	if len(args) < 1 {
		return Replies.CreateReplySyntaxErrorInParameters()
	}
	commandName := strings.TrimSpace(args[0])
	siteCommand, ok := SiteCommands.GetCommandFunction(commandName, cmd.cs, cmd.config)
	if !ok {
		return Replies.CreateReplySyntaxErrorInParameters()
	}
	return siteCommand.Execute(argString)
}

func (cmd SITE) Name() string {
	return "SITE"
}
