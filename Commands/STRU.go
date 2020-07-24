package Commands

import (
	"FTPserver/Configuration"
	"FTPserver/Connection"
	"FTPserver/Replies"
	"strings"
)

type STRU struct {
	cs     *Connection.Status
	config Configuration.FTPConfig
}

func (cmd STRU) Execute(args string) Replies.FTPReply {
	var toStructure Connection.Structure
	switch strings.ToLower(args) {
	case "f":
		toStructure = Connection.Structure(Connection.File)
	case "r":
		toStructure = Connection.Structure(Connection.Record)
	case "p":
		toStructure = Connection.Structure(Connection.Page)
	default:
		return Replies.CreateReplySyntaxErrorInParameters()
	}
	acceptableStructure := false
	for i := 0; i < len(cmd.config.SupportedStructures); i++ {
		if cmd.config.SupportedStructures[i] == toStructure {
			acceptableStructure = true
			break
		}
	}
	if acceptableStructure {
		cmd.cs.Structure = toStructure
		return Replies.CreateReplyCommandOkay()
	}
	return Replies.CreateReplyCommandNotImplementedForParameter()
}

func (cmd STRU) Name() string {
	return "STRU"
}
