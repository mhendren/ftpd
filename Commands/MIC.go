package Commands

import (
	"FTPserver/Configuration"
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

func (cmd MIC) IsExtendedCommand(cs *Connection.Status, _ Configuration.FTPConfig) bool {
	return cs.Security.SecurityMechanism.DiscloseMIC()
}

func (cmd MIC) AcceptedArguments(_ *Connection.Status, _ Configuration.FTPConfig) []string {
	return nil
}
