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

func (cmd EPRT) Execute(args string) Replies.FTPReply {
	TCPExecutor := func(fields []string, protocol string) Replies.FTPReply {
		address := fields[1]
		port := fields[2]

		localIP := strings.Split(cmd.cs.CommandConnection.LocalAddr().String(), ":")
		localPort, _ := strconv.Atoi(localIP[len(localIP)-1])
		localIPAddress := cmd.cs.CommandConnection.LocalAddr().(*net.TCPAddr).IP

		var remoteAddr string
		var localAddr string

		if protocol == "tcp6" {
			remoteAddr = fmt.Sprintf("[%v]:%v", address, port)
			localAddr = fmt.Sprintf("[%v]:%v", localIPAddress, (uint16)(localPort-1))
		} else {
			remoteAddr = fmt.Sprintf("%v:%v", address, port)
			localAddr = fmt.Sprintf("%v:%v", localIPAddress, (uint16)(localPort-1))
		}

		dataConnection := Connection.DataConnection{
			TransferType:    Connection.TransferType(Connection.Active),
			Protocol:        protocol,
			LocalAddress:    localAddr,
			RemoteAddress:   remoteAddr,
			Connection:      nil,
			PassivePort:     0,
			PassiveListener: nil,
			Security:        cmd.cs.Security,
			IsSetup:         false,
		}
		err := dataConnection.Setup()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error opening data connection: %v\n", err)
			return Replies.CreateReplyCantOpenDataConnection()
		}
		cmd.cs.DataConnection = dataConnection
		return Replies.CreateReplyCommandOkay()
	}

	if cmd.cs.EPSVAll {
		return Replies.CreateReplyBadCommandSequence()
	}
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

	protocols := map[int]string{
		1: "tcp4",
		2: "tcp6",
	}

	netProtocol, err := strconv.Atoi(fields[0])
	if err != nil {
		return Replies.CreateReplySyntaxErrorInParameters()
	}

	protocol, found := protocols[netProtocol]
	if !found {
		outValues := ""
		for k := range protocols {
			if outValues != "" {
				outValues += ", "
			}
			outValues += fmt.Sprintf("%v", k)
		}
		return Replies.CreateReplyUnsupportedExtendedPortProtocol(outValues)
	}
	cmd.cs.Type = Connection.TransferType(Connection.Active)
	return TCPExecutor(fields, protocol)
}
