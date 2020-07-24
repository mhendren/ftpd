package Commands

import (
	"FTPserver/Configuration"
	"FTPserver/Connection"
	"FTPserver/Replies"
)

type FTPCommand interface {
	Execute(args string) Replies.FTPReply
	Name() string
}

func GetCommandMap(cs *Connection.Status, config Configuration.FTPConfig) map[string]FTPCommand {
	return map[string]FTPCommand{
		// Standard RFC-959 Command Set
		"USER": USER{cs: cs, config: config},
		"NOOP": NOOP{},

		"PASS": PASS{cs: cs, config: config},
		"ACCT": ACCT{cs: cs, config: config},
		"CWD":  CWD{cs: cs, isCDUP: false},
		"CDUP": CWD{cs: cs, isCDUP: true},
		"SMNT": NotImplemented{},
		"QUIT": QUIT{cs: cs},
		"REIN": NotImplemented{},
		"PORT": PORT{cs: cs},
		"PASV": PASV{cs: cs},
		"TYPE": TYPE{cs: cs},
		"STRU": STRU{cs: cs, config: config},
		"MODE": MODE{cs: cs, config: config},
		"RETR": RETR{cs: cs},
		"STOR": STOR{cs: cs, isUnique: false, isAppend: false},
		"STOU": STOR{cs: cs, isUnique: true, isAppend: false},
		"APPE": STOR{cs: cs, isUnique: false, isAppend: true},
		"ALLO": NOOP{},
		// Updated in 3659 "REST": NotImplemented{},
		"RNFR": RNFR{cs: cs},
		"RNTO": RNTO{cs: cs},
		"ABOR": ABOR{cs: cs},
		"DELE": DELE{cs: cs},
		"RMD":  RMD{cs: cs},
		"MKD":  MKD{cs: cs},
		"PWD":  PWD{cs: cs},
		"LIST": LIST{cs: cs, abbreviated: false},
		"NLST": LIST{cs: cs, abbreviated: true},
		"SITE": SITE{cs: cs, config: config},
		"SYST": SYST{config: config},
		"STAT": STAT{cs: cs, config: config},
		"HELP": HELP{cs: cs, config: config},

		// RFC-1639 Command Set
		"LPRT": LPRT{cs: cs},
		"LPSV": LPSV{cs: cs},

		// RFC-2228 Command Set
		"ADAT": ADAT{cs: cs},
		"AUTH": AUTH{cs: cs},
		"CCC":  CCC{cs: cs},
		"CONF": CONF{cs: cs},
		"ENC":  ENC{cs: cs},
		"MIC":  MIC{cs: cs},
		"PBSZ": PBSZ{cs: cs},
		"PROT": PROT{cs: cs},

		// RFC-2389 Command Set
		"FEAT": FEAT{cs: cs, config: config},
		"OPTS": NotImplemented{},

		// RFC-2428 Command Set
		"EPRT": EPRT{cs: cs},
		"EPSV": EPSV{cs: cs},

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
		"XCUP": Undefined{},
		"XMKD": Undefined{},
		"XPWD": Undefined{},
		"XRCP": Undefined{},
		"XRMD": Undefined{},
		"XRSQ": Undefined{},
		"XSEM": Undefined{},
		"XSEN": Undefined{},
	}
}

func GetCommandFunction(command string, cs *Connection.Status, config Configuration.FTPConfig) (FTPCommand, bool) {
	commandMap := GetCommandMap(cs, config)
	cmd, ok := commandMap[command]
	return cmd, ok
}
