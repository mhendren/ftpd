package Commands

import (
	"FTPserver/Connection"
	"FTPserver/Replies"
	"encoding/base64"
)

type MIC struct {
	cs *Connection.Status
}

func (cmd MIC) Execute(args string) Replies.FTPReply {
	if !cmd.cs.Security.SecureCommandChannel {
		return Replies.CreateReplyBadCommandSequence()
	}
	decodedArgs, err := base64.StdEncoding.DecodeString(args)
	if err != nil {
		return Replies.CreateReplySyntaxErrorInParameters()
	}
	return cmd.cs.Security.SecurityMechanism.MIC(decodedArgs)
}

func (cmd MIC) Name() string {
	return "MIC"
}
