package Commands

import (
	"FTPserver/Connection"
	"FTPserver/Replies"
	"os"
	"path/filepath"
)

type DELE struct {
	cs *Connection.Status
}

func (cmd DELE) Execute(args string) Replies.FTPReply {
	if args == "" {
		return Replies.CreateReplySyntaxErrorInParameters()
	}
	fullPath := filepath.Join(cmd.cs.CurrentPath, args)
	_, err := os.Stat(fullPath)
	if err != nil {
		return Replies.CreateReplyRequestedFileActionNotAvailable()
	}
	err = os.Remove(fullPath)
	if err != nil {
		return Replies.CreateReplyRequestedActionNotTakenError()
	}
	return Replies.CreateReplyCommandOkay()
}
