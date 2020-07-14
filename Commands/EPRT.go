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

type EPRT struct {
	cs *Connection.Status
}

type EPRTExecutor interface {
	Execute() Replies.FTPReply
}

type IPv4Executor struct {
	cmd    EPRT
	fields []string
}

type IPv6Executor struct {
	cmd    EPRT
	fields []string
}

func (cmd EPRT) String() string {
	return ""
}

func (executor IPv4Executor) Execute() Replies.FTPReply {
	address := executor.fields[1]
	port := executor.fields[2]
	destination := fmt.Sprintf("%v:%v", address, port)
	localIP := strings.Split(executor.cmd.cs.CommandConnection.LocalAddr().String(), ":")
	localPort, _ := strconv.Atoi(localIP[len(localIP)-1])
	local, _ := net.ResolveTCPAddr("tcp4", fmt.Sprintf("%v:%v", localIP[0], (uint16)(localPort-1)))
	remote, _ := net.ResolveTCPAddr("tcp4", destination)
	conn, err := net.DialTCP("tcp4", local, remote)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error opening data connection: %v\n", err)
		return Replies.CreateReplyCantOpenDataConnection()
	}
	executor.cmd.cs.DataConnect(conn)
	return Replies.CreateReplyCommandOkay()
}

func (executor IPv6Executor) Execute() Replies.FTPReply {
	address := executor.fields[1]
	port := executor.fields[2]

	destinationSocket := fmt.Sprintf("[%v]:%v", address, port)
	localIP := strings.Split(executor.cmd.cs.CommandConnection.LocalAddr().String(), ":")
	localPort, _ := strconv.Atoi(localIP[len(localIP)-1])
	localIPAddress := executor.cmd.cs.CommandConnection.LocalAddr().(*net.TCPAddr).IP
	localSocket := fmt.Sprintf("[%v]:%v", localIPAddress, (uint16)(localPort-1))
	local, _ := net.ResolveTCPAddr("tcp6", localSocket)
	remote, _ := net.ResolveTCPAddr("tcp6", destinationSocket)
	conn, err := net.DialTCP("tcp6", local, remote)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error opening data connection: %v\n", err)
		return Replies.CreateReplyCantOpenDataConnection()
	}
	executor.cmd.cs.DataConnect(conn)
	return Replies.CreateReplyCommandOkay()
}

func (cmd EPRT) Execute(args string) Replies.FTPReply {
	if args == "" {
		return Replies.CreateReplySyntaxErrorInParameters()
	}
	delimiter := args[0]
	if delimiter < '!' || delimiter > '~' { // | is preferred, but any character between ! and ~ is technically valid
		return Replies.CreateReplySyntaxErrorInParameters()
	}
	delimiterString := fmt.Sprintf("%c", delimiter)
	fields := strings.Split(args, delimiterString)
	if len(fields) != 5 {
		return Replies.CreateReplySyntaxErrorInParameters()
	}
	fields = fields[1:4]

	executors := map[int]EPRTExecutor{
		1: IPv4Executor{cmd: cmd, fields: fields},
		2: IPv6Executor{cmd: cmd, fields: fields},
	}

	netProtocol, err := strconv.Atoi(fields[0])
	if err != nil {
		return Replies.CreateReplySyntaxErrorInParameters()
	}

	executor, found := executors[netProtocol]
	if !found {
		outValues := ""
		for k := range executors {
			if outValues != "" {
				outValues += ", "
			}
			outValues += fmt.Sprintf("%v", k)
		}
		return Replies.CreateReplyUnsupportedExtendedPortProtocol(outValues)
	}
	cmd.cs.Type = Connection.TransferType(Connection.Active)
	return executor.Execute()
}
