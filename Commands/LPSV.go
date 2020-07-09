package Commands

import (
	"FTPserver/Connection"
	"FTPserver/Replies"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

type LPSV struct {
	cs            *Connection.Status
	localListener net.Listener
}

func (cmd LPSV) String() string {
	localIP := cmd.localListener.Addr().(*net.TCPAddr).IP
	localSegments := strings.Split(cmd.localListener.Addr().String(), ":")
	localPort, _ := strconv.Atoi(localSegments[len(localSegments)-1])
	outString := "6,16,"
	for _, val := range localIP {
		outString += fmt.Sprintf("%v,", val)
	}
	outString += "2,"
	netPort := make([]byte, 2)
	binary.BigEndian.PutUint16(netPort, uint16(localPort))
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

	cmd.cs.Type = Connection.TransferType(Connection.Passive)

	if cmd.cs.DataConnected {
		return Replies.CreateReplyBadCommandSequence()
	}
	localIP := strings.Split(cmd.cs.CommandConnection.LocalAddr().String(), "]:")
	listenSocket, err := net.Listen("tcp6", fmt.Sprintf("%v]:0", localIP[0]))
	if err != nil {
		return disconnect("Disconnecting (error creating PASV listening socket)", err)
	}
	localListenIP := strings.Split(listenSocket.Addr().String(), "]:")
	if len(localListenIP) < 2 {
		return disconnect("Disconnecting, could not read port number for local socket", nil)
	}
	cmd.localListener = listenSocket
	cmd.cs.SendFTPReply(Replies.CreateReplyEnteringLongPassiveMode(cmd.String()))
	connectionSocket, err := listenSocket.Accept()
	if err != nil {
		return disconnect("Disconnecting (error getting local port number)", err)
	}
	cmd.cs.DataConnect(connectionSocket)
	return Replies.CreateReplyNoAction()
}
