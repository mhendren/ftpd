package Commands

import (
	"FTPserver/Configuration"
	"FTPserver/Connection"
	"FTPserver/Replies"
	"fmt"
	"strings"
)

type ACCT struct {
	cs     *Connection.Status
	config Configuration.FTPConfig
}

func (cmd ACCT) Execute(args string) Replies.FTPReply {
	isAnon := func(userName string) bool {
		lcUserName := strings.ToLower(userName)
		for _, anon := range cmd.config.AnonymousUsers {
			if strings.ToLower(anon) == lcUserName {
				return true
			}
		}
		return false
	}
	if !cmd.config.AllowAccount {
		return Replies.CreateReplyCommandNotImplementedSuperfluous()
	}
	if cmd.config.AccountValidator == nil || cmd.config.AccountValidator.Validate(args) {
		var message string
		if isAnon(cmd.cs.User) {
			cmd.cs.SetAuthenticated(false)
			cmd.cs.SetAnonymous(true)
			message = fmt.Sprintf("Logged in anonymously")
		} else {
			cmd.cs.SetAuthenticated(true)
			cmd.cs.SetAnonymous(false)
			message = fmt.Sprintf("%v logged in", cmd.cs.User)
		}
		cmd.cs.Account = args
		cmd.cs.SetCommands(nil)
		return Replies.CreateReplyUserLoggedIn(message)
	}
	return Replies.CreateReplyNeedAccount()
}
