package Commands

import (
	"FTPserver/Configuration"
	"FTPserver/Connection"
	"FTPserver/Replies"
	"encoding/base64"
)

type CONF struct {
	cs *Connection.Status
}

func (cmd CONF) Execute(args string) Replies.FTPReply {
	if !cmd.cs.Security.SecureCommandChannel {
		return Replies.CreateReplyBadCommandSequence()
	}
	decodedArgs, err := base64.StdEncoding.DecodeString(args)
	if err != nil {
		return Replies.CreateReplySyntaxErrorInParameters()
	}
	return cmd.cs.Security.SecurityMechanism.CONF(decodedArgs)
}

func (cmd CONF) Name() string {
	return "CONF"
}

func (cmd CONF) IsExtendedCommand(cs *Connection.Status, _ Configuration.FTPConfig) bool {
	return cs.Security.SecurityMechanism.DiscloseCONF()
}

func (cmd CONF) AcceptedArguments(_ *Connection.Status, _ Configuration.FTPConfig) []string {
	return nil
}
