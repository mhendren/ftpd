package Commands

import (
	"FTPserver/Configuration"
	"FTPserver/Connection"
	"FTPserver/Replies"
	"FTPserver/Validation"
	"fmt"
	"strings"
)

type PASS struct {
	cs     *Connection.Status
	config Configuration.FTPConfig
}

func (cmd PASS) Authenticate(_ string) bool {
	if cmd.config.AllowAnonymous &&
		strings.ToLower(cmd.cs.User) == "anonymous" ||
		strings.ToLower(cmd.cs.User) == "ftp" {
		return true
	}
	return false
}

func (cmd PASS) Execute(args string) Replies.FTPReply {
	if args == "" {
		return Replies.CreateReplyNotLoggedIn()
	}
	if cmd.config.PasswordValidator != nil {
		var message string
		auth := cmd.config.PasswordValidator.Authenticate(cmd.cs.User, args)
		if auth == Validation.AuthenticationType(Validation.Authenticated) {
			cmd.cs.SetAuthenticated(true)
			cmd.cs.SetAnonymous(false)
			message = fmt.Sprintf("%v logged in", cmd.cs.User)
		} else if auth == Validation.AuthenticationType(Validation.Anonymous) {
			cmd.cs.SetAuthenticated(false)
			cmd.cs.SetAnonymous(true)
			message = "Logged in anonymously"
		} else {
			return Replies.CreateReplyNotLoggedIn()
		}
		if cmd.config.RequiresAccount {
			cmd.cs.SetCommands(map[string]bool{"ACCT": true})
			return Replies.CreateReplyNeedAccount()
		}
		cmd.cs.SetCommands(nil)
		return Replies.CreateReplyUserLoggedIn(message)
	}
	return Replies.CreateReplyNotLoggedIn()
}
