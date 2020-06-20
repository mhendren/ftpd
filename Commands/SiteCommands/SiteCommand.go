package SiteCommands

import (
	"FTPserver/Configuration"
	"FTPserver/Connection"
	"FTPserver/Replies"
)

type SiteCommand interface {
	Execute(args string) Replies.FTPReply
	Help(args string) string
}

func GetCommandMap(cs *Connection.Status, config Configuration.FTPConfig) map[string]SiteCommand {
	return map[string]SiteCommand{
		"CHMOD": CHMOD{cs: cs},
		"HELP":  HELP{cs: cs, config: config},
		"IDLE":  IDLE{cs: cs},
		"UMASK": UMASK{cs: cs},
	}
}

func GetCommandFunction(command string, cs *Connection.Status, config Configuration.FTPConfig) (SiteCommand, bool) {
	commandMap := GetCommandMap(cs, config)
	cmd, ok := commandMap[command]
	return cmd, ok
}
