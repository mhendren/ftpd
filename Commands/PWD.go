package Commands

import (
	"FTPserver/Connection"
	"FTPserver/Replies"
)

type PWD struct {
	cs *Connection.Status
}

func (cmd PWD) Execute(_ string) Replies.FTPReply {
	return Replies.CreateReplyPathnameCreated(cmd.cs.CurrentPath, false)
}

func (cmd PWD) Name() string {
	return "PWD"
}
