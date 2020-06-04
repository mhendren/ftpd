package Commands

import (
	"FTPserver/Connection"
	"FTPserver/Replies"
	"os"
	"path/filepath"
)

type RMD struct {
	cs *Connection.Status
}

func (cmd RMD) Execute(args string) Replies.FTPReply {
	if args == "" {
		return Replies.CreateReplySyntaxErrorInParameters()
	}
	fullPath := filepath.Join(cmd.cs.CurrentPath, args)
	err := os.Remove(fullPath)
	if err != nil {
		return Replies.CreateReplyRequestedActionNotTakenError()
	}
	return Replies.CreateReplyCommandOkay()
}
