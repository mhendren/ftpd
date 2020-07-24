package Commands

import (
	"FTPserver/Configuration"
	"FTPserver/Connection"
	"FTPserver/Replies"
	"math"
	"strconv"
)

type PBSZ struct {
	cs *Connection.Status
}

func (cmd PBSZ) Execute(args string) Replies.FTPReply {
	// only TLS, a streaming system, is supported so only size 0 will work
	if !cmd.cs.Security.SecureCommandChannel {
		return Replies.CreateReplyBadCommandSequence()
	}
	size, err := strconv.Atoi(args)
	if err != nil || size < 0 || size > math.MaxInt32 {
		return Replies.CreateReplySyntaxErrorInParameters()
	}
	cmd.cs.Security.ProtectedBSize = 0
	reply := Replies.CreateReplyCommandOkay()
	if size > 0 {
		reply.Message = "PBSZ=0"
	}
	return reply
}

func (cmd PBSZ) Name() string {
	return "PBSZ"
}

func (cmd PBSZ) IsExtendedCommand(_ *Connection.Status, _ Configuration.FTPConfig) bool {
	return true
}

func (cmd PBSZ) AcceptedArguments(_ *Connection.Status, _ Configuration.FTPConfig) []string {
	return nil
}
