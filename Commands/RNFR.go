package Commands

import (
	"FTPserver/Connection"
	"FTPserver/Replies"
	"os"
	"path/filepath"
)

type RNFR struct {
	cs *Connection.Status
}

func (cmd RNFR) Execute(args string) Replies.FTPReply {
	if args == "" {
		return Replies.CreateReplySyntaxErrorInParameters()
	}
	fullPath := filepath.Join(cmd.cs.CurrentPath, args)
	_, err := os.Stat(fullPath)
	if err != nil {
		return Replies.CreateReplyRequestedFileActionNotAvailable()
	}
	cmd.cs.RenameFrom = fullPath
	cmd.cs.SetCommands(map[string]bool{"RNTO": true})
	return Replies.CreateReplyPendingFurtherInformation()
}
