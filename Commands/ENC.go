package Commands

import (
	"FTPserver/Connection"
	"FTPserver/Replies"
	"encoding/base64"
)

type ENC struct {
	cs *Connection.Status
}

func (cmd ENC) Execute(args string) Replies.FTPReply {
	if !cmd.cs.Security.SecureCommandChannel {
		return Replies.CreateReplyBadCommandSequence()
	}
	decodedArgs, err := base64.StdEncoding.DecodeString(args)
	if err != nil {
		return Replies.CreateReplySyntaxErrorInParameters()
	}
	return cmd.cs.Security.SecurityMechanism.ENC(decodedArgs)
}

func (cmd ENC) Name() string {
	return "ENC"
}
