package Commands

import (
	"FTPserver/Configuration"
	"FTPserver/Connection"
	"FTPserver/Replies"
	"encoding/base64"
)

type ADAT struct {
	cs *Connection.Status
}

func (cmd ADAT) Execute(args string) Replies.FTPReply {
	if !cmd.cs.Security.SecureCommandChannel {
		return Replies.CreateReplyBadCommandSequence()
	}
	decodedArgs, err := base64.StdEncoding.DecodeString(args)
	if err != nil {
		return Replies.CreateReplySyntaxErrorInParameters()
	}
	return cmd.cs.Security.SecurityMechanism.ADAT(decodedArgs)
}

func (cmd ADAT) Name() string {
	return "ADAT"
}

func (cmd ADAT) IsExtendedCommand(cs *Connection.Status, _ Configuration.FTPConfig) bool {
	return cs.Security.SecurityMechanism.DiscloseADAT()
}

func (cmd ADAT) AcceptedArguments(_ *Connection.Status, _ Configuration.FTPConfig) []string {
	return nil
}
