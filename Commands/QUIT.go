package Commands

import (
	"FTPserver/Connection"
	"FTPserver/Replies"
	"fmt"
	"os"
)

type QUIT struct {
	cs *Connection.Status
}

func (cmd QUIT) Execute(args string) Replies.FTPReply {
	cmd.cs.Disconnect()
	_, _ = fmt.Fprintln(os.Stderr, "Disconnecting")
	return Replies.CreateReplyClosingControlConnection("Goodbye")
}
