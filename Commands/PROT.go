package Commands

import (
	"FTPserver/Configuration"
	"FTPserver/Connection"
	"FTPserver/Replies"
)

type PROT struct {
	cs *Connection.Status
}

func (cmd PROT) Execute(args string) Replies.FTPReply {
	if cmd.cs.LastCommand != "PBSZ" || !cmd.cs.Security.SecureCommandChannel {
		return Replies.CreateReplyBadCommandSequence()
	}
	var dataLevelProtection Connection.DataChannelProtectionLevel
	switch args {
	case "C":
		dataLevelProtection = Connection.DataChannelProtectionLevel(Connection.Clear)
	case "P":
		dataLevelProtection = Connection.DataChannelProtectionLevel(Connection.Private)
	case "S":
		dataLevelProtection = Connection.DataChannelProtectionLevel(Connection.Safe)
	case "E":
		dataLevelProtection = Connection.DataChannelProtectionLevel(Connection.Confidential)
	default:
		return Replies.CreateReplyCommandNotImplementedForParameter()
	}
	isSupported := false
	for _, b := range cmd.cs.Security.SecurityMechanism.SupportedProtections() {
		if dataLevelProtection == b {
			isSupported = true
			break
		}
	}
	if !isSupported {
		return Replies.CreateReplyRequestedPROTNotSupported()
	}
	if cmd.cs.Security.ProtectedBSize == -1 {
		return Replies.CreateReplyBadCommandSequence()
	}
	cmd.cs.Security.ProtectionLevel = dataLevelProtection
	return Replies.CreateReplyCommandOkay()
}

func (cmd PROT) Name() string {
	return "PROT"
}

func (cmd PROT) IsExtendedCommand(_ *Connection.Status, _ Configuration.FTPConfig) bool {
	return true
}

func (cmd PROT) AcceptedArguments(_ *Connection.Status, _ Configuration.FTPConfig) []string {
	return nil
}
