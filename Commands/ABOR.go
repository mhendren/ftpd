package Commands

import (
	"FTPserver/Connection"
	"FTPserver/Replies"
)

type ABOR struct {
	cs *Connection.Status
}

func (cmd ABOR) Execute(_ string) Replies.FTPReply {
	if cmd.cs.DataConnected {
		defer cmd.cs.DataDisconnect()
	}
	return Replies.CreateReplyClosingDataConnection()
}

func (cmd ABOR) Name() string {
	return "ABOR"
}
