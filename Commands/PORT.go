package Commands

import (
	"FTPserver/Connection"
	"FTPserver/Replies"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PORT struct {
	cs *Connection.Status
}

func (cmd PORT) Execute(args string) Replies.FTPReply {
	if cmd.cs.EPSVAll {
		return Replies.CreateReplyBadCommandSequence()
	}
	cmd.cs.Type = Connection.TransferType(Connection.Active)
	segments := strings.Split(args, ",")
	if len(segments) != 6 {
		return Replies.CreateReplySyntaxErrorInParameters()
	}
	hiPort, err := strconv.Atoi(segments[4])
	if err != nil {
		return Replies.CreateReplySyntaxErrorInParameters()
	}
	loPort, err := strconv.Atoi(segments[5])
	if err != nil {
		return Replies.CreateReplySyntaxErrorInParameters()
	}
	port := uint16(hiPort*256) + uint16(loPort)

	remoteAddr := fmt.Sprintf("%v.%v.%v.%v:%v", segments[0], segments[1], segments[2], segments[3], port)
	localIP := strings.Split(cmd.cs.CommandConnection.LocalAddr().String(), ":")
	localPort, _ := strconv.Atoi(localIP[1])
	localAddr := fmt.Sprintf("%v:%v", localIP[0], (uint16)(localPort-1))

	dataConnection := Connection.DataConnection{
		TransferType:    Connection.TransferType(Connection.Active),
		Protocol:        "tcp4",
		LocalAddress:    localAddr,
		RemoteAddress:   remoteAddr,
		Connection:      nil,
		PassivePort:     0,
		PassiveListener: nil,
		Security:        cmd.cs.Security,
		IsSetup:         false,
	}
	err = dataConnection.Setup()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error opening data connection: %v\n", err)
		return Replies.CreateReplyCantOpenDataConnection()
	}
	cmd.cs.DataConnection = dataConnection
	return Replies.CreateReplyCommandOkay()
}
