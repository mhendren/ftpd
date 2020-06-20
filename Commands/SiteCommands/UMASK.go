package SiteCommands

import (
	"FTPserver/Connection"
	"FTPserver/Replies"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type UMASK struct {
	cs *Connection.Status
}

func (siteCommand UMASK) Execute(argString string) Replies.FTPReply {
	args := strings.Split(argString, " ")
	for _, val := range args {
		strings.TrimSpace(val)
	}
	if len(args) == 1 {
		return Replies.CreateReplyCommandOkayCustom(fmt.Sprintf("Current umask is %03o", siteCommand.cs.Umask))
	}
	if len(args) != 2 {
		return Replies.CreateReplySyntaxErrorInParameters()
	}

	modeInt, err := strconv.ParseUint(args[1], 8, 32)
	if err != nil {
		return Replies.CreateReplySyntaxErrorInParameters()
	}
	siteCommand.cs.Umask = os.FileMode(modeInt)
	return Replies.CreateReplyCommandOkay()
}

func (siteCommand UMASK) Help(_ string) string {
	return "Syntax: SITE UMASK [ <sp> umask ]"
}
