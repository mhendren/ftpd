package Commands

import (
	"FTPserver/Connection"
	"FTPserver/Replies"
	"crypto/tls"
	"fmt"
	"net/textproto"
)

type AUTH struct {
	cs *Connection.Status
}

func (cmd AUTH) Execute(args string) Replies.FTPReply {
	AuthTLSExecute := func() Replies.FTPReply {
		authConn := tls.Server(cmd.cs.CommandConnection, cmd.cs.Security.Config)
		if authConn == nil {
			return Replies.CreateReplyFailedSecurityCheck()
		}
		cmd.cs.SendFTPReply(Replies.CreateReplySecurityDataExchangeComplete())
		cmd.cs.Security.IsSecure = true
		cmd.cs.Security.CommandConnection = authConn
		cmd.cs.CommandConnection = cmd.cs.Security.CommandConnection
		cmd.cs.Security.TextProto = textproto.NewConn(cmd.cs.CommandConnection)
		cmd.cs.TextProto = cmd.cs.Security.TextProto
		return Replies.CreateReplyNoAction()
	}

	AuthSSLExecute := func() Replies.FTPReply {
		return Replies.CreateReplyRequestDeniedForPolicyReasons()
	}

	AuthUnknownExecute := func() Replies.FTPReply {
		reply := Replies.CreateReplyCommandNotImplementedForParameter()
		reply.Message = fmt.Sprintf("Security mechanism \"%v\" is not understood", args)
		return reply
	}

	funcMap := map[string]func() Replies.FTPReply{
		"TLS": AuthTLSExecute,
		"SSL": AuthSSLExecute,
	}

	fn, ok := funcMap[args]
	if !ok {
		fn = AuthUnknownExecute
	}
	return fn()
}
