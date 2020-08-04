package Commands

import (
	"FTPserver/Commands/OptionsCommands"
	"FTPserver/Configuration"
	"FTPserver/Connection"
	"FTPserver/Replies"
	"strings"
)

type OPTS struct {
	cs     *Connection.Status
	config Configuration.FTPConfig
}

func (cmd OPTS) Execute(args string) Replies.FTPReply {
	splitArgs := strings.SplitN(args, " ", 1)
	if len(splitArgs) < 1 {
		return Replies.CreateReplySyntaxErrorInParameters()
	}
	command := splitArgs[0]
	rest := ""
	if len(splitArgs) > 1 {
		rest = splitArgs[1]
	}
	oCmd, ok := OptionsCommands.GetCommandFunction(command, cmd.cs, cmd.config)
	if !ok {
		return Replies.CreateReplySyntaxErrorInParameters()
	}
	err := oCmd.OptionsSet(cmd.cs, command, rest)
	if err != nil {
		return Replies.CreateReplySyntaxErrorInParameters()
	}
	return Replies.CreateReplyCommandOkay()
}

func (cmd OPTS) Name() string {
	return "OPTS"
}
