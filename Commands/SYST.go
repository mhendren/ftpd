package Commands

import (
	"FTPserver/Configuration"
	"FTPserver/Replies"
)

type SYST struct {
	config Configuration.FTPConfig
}

func (cmd SYST) Execute(_ string) Replies.FTPReply {
	system := "UNIX" // Even on windows we are simulating Unix operating
	if cmd.config.SystemType != "" {
		system = cmd.config.SystemType
	}
	return Replies.CreateReplyNameSystemType(system)
}
