package SiteCommands

import (
	"FTPserver/Connection"
	"FTPserver/Replies"
	"fmt"
	"strconv"
	"strings"
)

type IDLE struct {
	cs *Connection.Status
}

func (siteCommand IDLE) Execute(argString string) Replies.FTPReply {
	args := strings.Split(argString, " ")
	for _, val := range args {
		strings.TrimSpace(val)
	}
	if len(args) == 1 {
		return Replies.CreateReplyCommandOkayCustom(fmt.Sprintf("Current idle timeout is %v second",
			siteCommand.cs.IdleTimeout))
	}
	if len(args) != 2 {
		return Replies.CreateReplySyntaxErrorInParameters()
	}
	idle, err := strconv.Atoi(args[1])
	if err != nil {
		return Replies.CreateReplySyntaxErrorInParameters()
	}
	siteCommand.cs.IdleTimeout = idle
	return Replies.CreateReplyCommandOkay()
}

func (siteCommand IDLE) Help(_ string) string {
	return "Syntax: SITE IDLE [<sp> idle-timeout]"
}
