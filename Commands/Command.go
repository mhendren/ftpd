package Commands

import (
	"FTPserver/Configuration"
	"FTPserver/Connection"
	"FTPserver/Replies"
)

type FTPCommand interface {
	Execute(args string) Replies.FTPReply
}

func GetCommandFunction(command string, cs *Connection.Status, config Configuration.FTPConfig) (FTPCommand, bool) {
	commandMap := map[string]FTPCommand{
		"USER": USER{cs: cs},
		"NOOP": NOOP{},

		"PASS": NotImplemented{},
		"ACCT": NotImplemented{},
		"CWD":  NotImplemented{},
		"CDUP": NotImplemented{},
		"SMNT": NotImplemented{},
		"QUIT": QUIT{cs: cs},
		"REIN": NotImplemented{},
		"PORT": NotImplemented{},
		"PASV": NotImplemented{},
		"TYPE": NotImplemented{},
		"STRU": NotImplemented{},
		"MODE": NotImplemented{},
		"RETR": NotImplemented{},
		"STOR": NotImplemented{},
		"STOU": NotImplemented{},
		"APPE": NotImplemented{},
		"ALLO": NotImplemented{},
		"REST": NotImplemented{},
		"RNFR": NotImplemented{},
		"RNTO": NotImplemented{},
		"ABOR": NotImplemented{},
		"DELE": NotImplemented{},
		"RMD":  NotImplemented{},
		"MKD":  NotImplemented{},
		"PWD":  NotImplemented{},
		"LIST": NotImplemented{},
		"NLST": NotImplemented{},
		"SITE": NotImplemented{},
		"SYST": NotImplemented{},
		"STAT": NotImplemented{},
		"HELP": NotImplemented{},
	}
	cmd, ok := commandMap[command]
	return cmd, ok
}
