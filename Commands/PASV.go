package Commands

import (
	"FTPserver/Connection"
	"FTPserver/Replies"
	"fmt"
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
	dataConnection := Connection.DataConnection{
		TransferType:    Connection.TransferType(Connection.Passive),
		Protocol:        "tcp",
		LocalAddress:    fmt.Sprintf("%v:0", localIP[0]),
		RemoteAddress:   "",
		Connection:      nil,
		PassivePort:     0,
		PassiveListener: nil,
		Security:        cmd.cs.Security,
		IsSetup:         false,
	}
	err := dataConnection.Setup()
	if err != nil {
		return disconnect("Disconnecting (error creating PASV listening socket)", err)
	}
	cmd.cs.DataConnection = dataConnection
	return Replies.CreateReplyEnteringPassiveMode(cmd.cs.DataConnection.PassiveListener.Addr())
}
