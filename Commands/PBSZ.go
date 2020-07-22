package Commands

import (
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
	if !cmd.cs.Security.IsSecure {
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
