package Commands

import (
	"FTPserver/Connection"
	"FTPserver/Replies"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

type PORT struct {
	cs *Connection.Status
}

func (cmd PORT) Execute(args string) Replies.FTPReply {
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
	destination := fmt.Sprintf("%v.%v.%v.%v:%v", segments[0], segments[1], segments[2], segments[3], port)
	localIP := strings.Split(cmd.cs.CommandConnection.LocalAddr().String(), ":")
	localPort, _ := strconv.Atoi(localIP[1])
	local, _ := net.ResolveTCPAddr("tcp4", fmt.Sprintf("%v:%v", localIP[0], (uint16)(localPort-1)))
	remote, _ := net.ResolveTCPAddr("tcp4", destination)
	conn, err := net.DialTCP("tcp4", local, remote)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error opening data connection: %v\n", err)
		return Replies.CreateReplyCantOpenDataConnection()
	}
	cmd.cs.DataConnect(conn)
	return Replies.CreateReplyCommandOkay()
}
