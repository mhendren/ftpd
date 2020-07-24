package Commands

import (
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
