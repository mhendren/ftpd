package Commands

import (
	"FTPserver/Connection"
	"FTPserver/Replies"
	"os"
	"path/filepath"
)

type MKD struct {
	cs *Connection.Status
}

func (cmd MKD) Execute(args string) Replies.FTPReply {
	if args == "" {
		return Replies.CreateReplySyntaxErrorInParameters()
	}
	fullPath := filepath.Join(cmd.cs.CurrentPath, args)
	err := os.Mkdir(fullPath, 0755)
	if err != nil {
		return Replies.CreateReplyRequestedActionNotTakenError()
	}
	return Replies.CreateReplyPathnameCreated(fullPath, true)
}
