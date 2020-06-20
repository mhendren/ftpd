package SiteCommands

import (
	"FTPserver/Connection"
	"FTPserver/Replies"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type CHMOD struct {
	cs *Connection.Status
}

func (siteCommand CHMOD) Execute(argString string) Replies.FTPReply {
	args := strings.Split(argString, " ")
	for _, val := range args {
		strings.TrimSpace(val)
	}
	if len(args) != 3 {
		return Replies.CreateReplySyntaxErrorInParameters()
	}
	mode := args[1]
	fileName := filepath.Join(siteCommand.cs.CurrentPath, args[2])

	modeInt, err := strconv.ParseUint(mode, 8, 32)
	if err != nil {
		return Replies.CreateReplySyntaxErrorInParameters()
	}
	err = os.Chmod(fileName, os.FileMode(modeInt))
	if err != nil {
		return Replies.CreateReplyRequestedFileActionAbortedError()
	}
	return Replies.CreateReplyFileActionCompleted(args[2])
}

func (siteCommand CHMOD) Help(_ string) string {
	return "Syntax: SITE CHMOD <sp> mode <sp> filename"
}
