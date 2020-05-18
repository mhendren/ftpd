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
	if args == "" {
		return Replies.CreateReplyNeedAccount()
	}
	cmd.cs.SetUser(args)
	cmd.cs.SetAnonymous(strings.ToLower(args) == "anonymous")
	_, _ = fmt.Fprintf(os.Stderr, "User connecting as %v", args)
	cmd.cs.SetCommands(map[string]bool{"PASS": true})
	return Replies.CreateReplyNeedPassword()
}
