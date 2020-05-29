package Commands

import (
	"FTPserver/Configuration"
	"FTPserver/Connection"
	"FTPserver/Replies"
	"strings"
)

type PASS struct {
	cs     *Connection.Status
	config Configuration.FTPConfig
}

func (cmd PASS) Authenticate(_ string) bool {
	if strings.ToLower(cmd.cs.User) == "anonymous" {
		return true
	}
	return false
}

func (cmd PASS) Execute(args string) Replies.FTPReply {
	if args == "" {
		return Replies.CreateReplyNotLoggedIn()
	}
	if cmd.Authenticate(args) {
		cmd.cs.Authenticated = true
		cmd.cs.SetCommands(nil)
		return Replies.CreateReplyCommandOkay()
	}
	return Replies.CreateReplyNotLoggedIn()
}
