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
		"PASV": PASV{cs: cs},
		"TYPE": TYPE{cs: cs},
		"STRU": NotImplemented{},
		"MODE": NotImplemented{},
		"RETR": RETR{cs: cs},
		"STOR": STOR{cs: cs, isUnique: false},
		"STOU": STOR{cs: cs, isUnique: true},
		"APPE": NotImplemented{},
		"ALLO": NotImplemented{},
		// Updated in 3659 "REST": NotImplemented{},
		"RNFR": RNFR{cs: cs},
		"RNTO": RNTO{cs: cs},
		"ABOR": NotImplemented{},
		"DELE": DELE{cs: cs},
		"RMD":  RMD{cs: cs},
		"MKD":  MKD{cs: cs},
		"PWD":  PWD{cs: cs},
		"LIST": LIST{cs: cs, abbreviated: false},
		"NLST": LIST{cs: cs, abbreviated: true},
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
