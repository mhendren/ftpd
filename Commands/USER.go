package Commands

import (
	"FTPserver/Configuration"
	"FTPserver/Connection"
	"FTPserver/Replies"
	"fmt"
	"os"
	"strings"
)

type USER struct {
	cs     *Connection.Status
	config Configuration.FTPConfig
}

func (cmd USER) Execute(args string) Replies.FTPReply {
	if args == "" ||
		((strings.ToLower(args) == "anonymous" || strings.ToLower(args) == "ftp") &&
			!cmd.config.AllowAnonymous) {
		return Replies.CreateReplyNeedAccount()
	}
	cmd.cs.SetAnonymous(strings.ToLower(args) == "anonymous")
	cmd.cs.SetUser(args)
	_, _ = fmt.Fprintf(os.Stderr, "User connecting as %v", args)
	if !cmd.config.RequiresPassword {
		if cmd.config.RequiresAccount {
			cmd.cs.SetCommands(map[string]bool{"ACCT": true})
			return Replies.CreateReplyNeedAccount()
		}
		return Replies.CreateReplyUserLoggedIn("User logged in")
	}
	cmd.cs.SetCommands(map[string]bool{"PASS": true})
	return Replies.CreateReplyNeedPassword()
}
