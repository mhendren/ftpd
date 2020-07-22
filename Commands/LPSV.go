package Commands

import (
	"FTPserver/Connection"
	"FTPserver/Replies"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strings"
)

type LPSV struct {
	cs            *Connection.Status
	localListener net.Listener
}

func (cmd LPSV) String() string {
	localIP := cmd.cs.DataConnection.PassiveListener.Addr().(*net.TCPAddr).IP
	outString := "6,16,"
	for _, val := range localIP {
		outString += fmt.Sprintf("%v,", val)
	}
	outString += "2,"
	netPort := make([]byte, 2)
	binary.BigEndian.PutUint16(netPort, uint16(cmd.cs.DataConnection.PassivePort))
	outString += fmt.Sprintf("%v,%v", netPort[0], netPort[1])
	return outString
}

func (cmd LPSV) Execute(_ string) Replies.FTPReply {
	disconnect := func(message string, err error) Replies.FTPReply {
		cmd.cs.Disconnect()
		_, _ = fmt.Fprintln(os.Stderr, message)
		_, _ = fmt.Fprintln(os.Stderr, fmt.Errorf("error: %s", err))
		return Replies.CreateReplyNotAvailableClosingConnection()
	}

	if cmd.cs.EPSVAll {
		return Replies.CreateReplyBadCommandSequence()
	}

	cmd.cs.Type = Connection.TransferType(Connection.Passive)

	if cmd.cs.DataConnected {
		return Replies.CreateReplyBadCommandSequence()
	}
	localAddr := strings.Split(cmd.cs.CommandConnection.LocalAddr().String(), "]:")
	dataConnection := Connection.DataConnection{
		TransferType:    Connection.TransferType(Connection.Passive),
		Protocol:        "tcp",
		LocalAddress:    localAddr[0] + "]:0",
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
	return Replies.CreateReplyEnteringLongPassiveMode(cmd.String())
}
