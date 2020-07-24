package Commands

import (
	"FTPserver/Configuration"
	"FTPserver/Connection"
	"FTPserver/Replies"
	"strings"
)

type MODE struct {
	cs     *Connection.Status
	config Configuration.FTPConfig
}

func (cmd MODE) Execute(args string) Replies.FTPReply {
	var toMode Connection.TransferMode
	switch strings.ToLower(args) {
	case "s":
		toMode = Connection.TransferMode(Connection.Stream)
	case "b":
		toMode = Connection.TransferMode(Connection.Block)
	case "c":
		toMode = Connection.TransferMode(Connection.Compressed)
	default:
		return Replies.CreateReplySyntaxErrorInParameters()
	}
	acceptableMode := false
	for i := 0; i < len(cmd.config.SupportedStructures); i++ {
		if cmd.config.SupportedModes[i] == toMode {
			acceptableMode = true
			break
		}
	}
	if acceptableMode {
		cmd.cs.Mode = toMode
		return Replies.CreateReplyCommandOkay()
	}
	return Replies.CreateReplyCommandNotImplementedForParameter()
}

func (cmd MODE) Name() string {
	return "MODE"
}
