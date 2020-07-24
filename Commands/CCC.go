package Commands

import (
	"FTPserver/Configuration"
	"FTPserver/Connection"
	"FTPserver/Replies"
)

type CCC struct {
	cs *Connection.Status
}

// Note: absolutely no one thinks supporting CCC is a good idea. In fact the very idea of starting a channel
// for encryption, and then backing off of that is bad idea. The problem that it solves is way better handled
// with Passive mode. But I want a complete solution, so I am going to implement this against my better judgement.
// I am making it configurable with a default of Forbid.

// Note: I found that cURL supports CCC by using the --ftp-ssl-ccc flag. This will send the CCC and denegotiate
// TLS on the command channel. Subsequent commands on the command channel come across unencrypted.

func (cmd CCC) Execute(_ string) Replies.FTPReply {
	if cmd.cs.Security.CCCStatus == Connection.CCCStatus(Connection.Forbid) {
		return Replies.CreateReplySyntaxError()
	}
	if !cmd.cs.Security.SecureCommandChannel {
		return Replies.CreateReplyCommandProtectionDeniedPolicy()
	}

	// From a post on Stack Overflow, no one seems to thing there is an official way to close a TLS socket and
	// retain the underlying connection. So the solution is to keep track of the underlying socket, and signal
	// to the TLS connection that you want to quit writing (and then force a read to get the TLS equivalent of
	// FIN wait. Follow this up by resuming the underlying connection as the command connection.

	// Additional note, the REPLY to CCC is still sent encrypted
	cmd.cs.SendFTPReply(Replies.CreateReplyCommandOkay())

	// Best code sample I could find on StackOverflow for removing the TLS from the connection without closing
	// the underlying connection

	_ = cmd.cs.Security.CommandConnection.CloseWrite()
	buffer := make([]byte, 1)
	_, _ = cmd.cs.Security.CommandConnection.Read(buffer)

	// Reset the Command channel components
	cmd.cs.CommandConnection = cmd.cs.StandardConnection
	cmd.cs.TextProto = cmd.cs.StandardTextProto

	cmd.cs.Security.SecureCommandChannel = false

	// Already responded to the CCC so do not allow the caller to respond.
	return Replies.CreateReplyNoAction()
}

func (cmd CCC) Name() string {
	return "CCC"
}

func (cmd CCC) IsExtendedCommand(cs *Connection.Status, _ Configuration.FTPConfig) bool {
	return cs.Security.CCCStatus == Connection.CCCStatus(Connection.Allow)
}

func (cmd CCC) AcceptedArguments(_ *Connection.Status, _ Configuration.FTPConfig) []string {
	return nil
}
