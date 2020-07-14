package Commands

import (
	"FTPserver/Connection"
	"FTPserver/Replies"
	"fmt"
	"net"
	"os"
	"strconv"
)

type EPSV struct {
	cs *Connection.Status
}

func (cmd EPSV) String() string {
	return ""
}

func (cmd EPSV) Execute(args string) Replies.FTPReply {
	disconnect := func(message string, err error) Replies.FTPReply {
		cmd.cs.Disconnect()
		_, _ = fmt.Fprintln(os.Stderr, message)
		_, _ = fmt.Fprintln(os.Stderr, fmt.Errorf("error: %s", err))
		return Replies.CreateReplyNotAvailableClosingConnection()
	}

	if args != "" {
		if args == "ALL" {
			cmd.cs.EPSVAll = true
		} else {
			val, err := strconv.Atoi(args)
			if err != nil {
				return Replies.CreateReplySyntaxErrorInParameters()
			}
			if val < 1 || val > 2 {
				return Replies.CreateReplyUnsupportedExtendedPortProtocol("1, 2")
			}
			cmd.cs.PreferredEProtocol = val
		}
	}

	cmd.cs.Type = Connection.TransferType(Connection.Passive)

	if cmd.cs.DataConnected {
		return Replies.CreateReplyBadCommandSequence()
	}
	tcpAddr, err := net.ResolveTCPAddr("tcp", cmd.cs.CommandConnection.LocalAddr().String())
	if err != nil {
		return disconnect("Disconnecting (error creating EPSV listening socket)", err)
	}

	var lAddr string
	if tcpAddr.IP.To4() == nil && cmd.cs.PreferredEProtocol == 0 {
		lAddr = fmt.Sprintf("[%v]:0", tcpAddr.IP)
	} else if cmd.cs.PreferredEProtocol == 0 {
		lAddr = fmt.Sprintf("%v:0", tcpAddr.IP)
	}
	if cmd.cs.PreferredEProtocol != 0 {
		lAddr = ":0"
	}

	listenSocket, err := net.Listen("tcp", lAddr)
	if err != nil {
		return disconnect("Disconnecting (error creating EPSV listening socket)", err)
	}

	cmd.cs.SendFTPReply(Replies.CreateReplyEnteringExtendedPassiveMode(listenSocket.Addr()))
	connectionSocket, err := listenSocket.Accept()
	if err != nil {
		return disconnect("Disconnecting (error getting local port number)", err)
	}
	cmd.cs.DataConnect(connectionSocket)
	return Replies.CreateReplyNoAction()
}
