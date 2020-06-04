package Commands

import (
	"FTPserver/Connection"
	"FTPserver/Replies"
	"os"
	"path/filepath"
)

type RNTO struct {
	cs *Connection.Status
}

func (cmd RNTO) Execute(args string) Replies.FTPReply {
	defer func() {
		cmd.cs.SetCommands(nil)
		cmd.cs.RenameFrom = ""
	}()

	if cmd.cs.AcceptableCommands == nil || !cmd.cs.CanUseCommand("RNTO") {
		// RNTO isn't explicitly listed in available command, so this command is out of sequence
		return Replies.CreateReplyBadCommandSequence()
	}
	if args == "" {
		return Replies.CreateReplySyntaxErrorInParameters()
	}
	fullPath := filepath.Join(cmd.cs.CurrentPath, args)
	_, err := os.Stat(fullPath)
	if err == nil {
		// no error, so the to filename already exists
		return Replies.CreateReplyRequestedActionNotTakenErrorFilename()
	}
	err = os.Rename(cmd.cs.RenameFrom, fullPath)
	if err != nil {
		return Replies.CreateReplyRequestedActionNotTakenErrorFilename()
	}
	return Replies.CreateReplyCommandOkay()
}
