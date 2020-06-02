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
		// Standard RFC-959 Command Set
		"USER": USER{cs: cs},
		"NOOP": NOOP{},

		"PASS": PASS{cs: cs, config: config},
		"ACCT": NotImplemented{},
		"CWD":  CWD{cs: cs, isCDUP: false},
		"CDUP": CWD{cs: cs, isCDUP: true},
		"SMNT": NotImplemented{},
		"QUIT": QUIT{cs: cs},
		"REIN": NotImplemented{},
		"PORT": PORT{cs: cs},
		"PASV": NotImplemented{},
		"TYPE": TYPE{cs: cs},
		"STRU": NotImplemented{},
		"MODE": NotImplemented{},
		"RETR": NotImplemented{},
		"STOR": STOR{cs: cs},
		"STOU": NotImplemented{},
		"APPE": NotImplemented{},
		"ALLO": NotImplemented{},
		// Updated in 3659 "REST": NotImplemented{},
		"RNFR": NotImplemented{},
		"RNTO": NotImplemented{},
		"ABOR": NotImplemented{},
		"DELE": NotImplemented{},
		"RMD":  NotImplemented{},
		"MKD":  NotImplemented{},
		"PWD":  PWD{cs: cs},
		"LIST": LIST{cs: cs},
		"NLST": NotImplemented{},
		"SITE": NotImplemented{},
		"SYST": NotImplemented{},
		"STAT": NotImplemented{},
		"HELP": NotImplemented{},

		// RFC-1639 Command Set
		"LPRT": NotImplemented{},
		"LRSV": NotImplemented{},

		// RFC-2228 Command Set
		"ADAT": NotImplemented{},
		"AUTH": NotImplemented{},
		"CCC":  NotImplemented{},
		"CONF": NotImplemented{},
		"ENC":  NotImplemented{},
		"MIC":  NotImplemented{},
		"PBSZ": NotImplemented{},
		"PROT": NotImplemented{},

		// RFC-2389 Command Set
		"FEAT": NotImplemented{},
		"OPTS": NotImplemented{},

		// RFC-2428 Command Set
		"EPRT": NotImplemented{},
		"EPSV": NotImplemented{},

		// RFC-2640 Command Set
		"LANG": NotImplemented{},

		// RFC-3659 Command Set
		"MDTM": NotImplemented{},
		"MLSD": NotImplemented{},
		"MLST": NotImplemented{},
		"REST": NotImplemented{},
		"SIZE": NotImplemented{},

		// RFC-7151 Command Set
		"HOST": NotImplemented{},

		// RFC-775,743,737 We do not do this Command Set
		"XCUP": NotImplemented{},
		"XMKD": NotImplemented{},
		"XPWD": NotImplemented{},
		"XRCP": NotImplemented{},
		"XRMD": NotImplemented{},
		"XRSQ": NotImplemented{},
		"XSEM": NotImplemented{},
		"XSEN": NotImplemented{},
	}
	cmd, ok := commandMap[command]
	return cmd, ok
}
