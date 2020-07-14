package Commands

import (
	"FTPserver/Connection"
	"FTPserver/Replies"
	"fmt"
	"net"
	"os"
	"strings"
)

type PASV struct {
	cs *Connection.Status
}

func (cmd PASV) Execute(_ string) Replies.FTPReply {
	disconnect := func(message string, err error) Replies.FTPReply {
		cmd.cs.Disconnect()
		_, _ = fmt.Fprintln(os.Stderr, message)
		_, _ = fmt.Fprintln(os.Stderr, fmt.Errorf("error: %s", err))
		return Replies.CreateReplyNotAvailableClosingConnection()
	}

	if cmd.cs.EPSVAll {
		Replies.CreateReplyBadCommandSequence()
	}

	cmd.cs.Type = Connection.TransferType(Connection.Passive)

	if cmd.cs.DataConnected {
		return Replies.CreateReplyBadCommandSequence()
	}
	localIP := strings.Split(cmd.cs.CommandConnection.LocalAddr().String(), ":")
	listenSocket, err := net.Listen("tcp", fmt.Sprintf("%v:0", localIP[0]))
	if err != nil {
		return disconnect("Disconnecting (error creating PASV listening socket)", err)
	}
	localListenIP := strings.Split(listenSocket.Addr().String(), ":")
	if len(localListenIP) < 2 {
		return disconnect("Disconnecting, could not read port number for local socket", nil)
	}
	cmd.cs.SendFTPReply(Replies.CreateReplyEnteringPassiveMode(listenSocket.Addr()))
	connectionSocket, err := listenSocket.Accept()
	if err != nil {
		return disconnect("Disconnecting (error getting local port number)", err)
	}
	cmd.cs.DataConnect(connectionSocket)
	return Replies.CreateReplyNoAction()
}
