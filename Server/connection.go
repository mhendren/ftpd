package Server

import (
	"FTPserver/Commands"
	"FTPserver/Configuration"
	"FTPserver/Connection"
	"FTPserver/Replies"
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"net"
	"net/textproto"
	"os"
	"strings"
)

const preambleText = "Welcome to mhendren FTP\nTesting that this will handle more than 2 lines\nlike this"
const openPreambleText = "Welcome to mhendren FTP\nThis is a free for all\nNo account needed to browse"

func connectionLog(connectionSocket net.Conn) {
	fmt.Printf("connection from: %v\n", connectionSocket.RemoteAddr())
}

var connectionStatus Connection.Status

func connectionPreamble(connectionSocket net.Conn, config Configuration.FTPConfig) {
	var err error
	if config.AllowNoLogin {
		_, err = fmt.Fprint(connectionSocket, Replies.CreateReplyUserLoggedIn(openPreambleText))
	} else {
		_, err = fmt.Fprint(connectionSocket, Replies.CreateReplyReadyForNewUser(preambleText))
	}
	if err != nil {
		fmt.Printf("error writing data: %v\n", err)
	}
}

func splitData(dataLine string) (string, string) {
	split := strings.SplitAfterN(dataLine, " ", 2)
	var command string
	var args string
	if len(split) > 0 {
		command = split[0]
		if len(split) > 1 {
			args = split[1]
		}
	}
	return strings.TrimSpace(command), strings.TrimSpace(args)
}

func commandLoop(connectionSocket net.Conn, config Configuration.FTPConfig) {
	connectionStatus.TextProto = textproto.NewConn(connectionSocket)

	for connectionStatus.IsConnected() {
		line, err := connectionStatus.TextProto.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
			break
		}
		fmt.Printf("read %v bytes: %v\n", len(line), line)
		command, args := splitData(line)
		if !connectionStatus.CanUseCommand(command) {
			_, _ = fmt.Fprint(connectionSocket, Replies.CreateReplyBadCommandSequence())
			continue
		}
		fmt.Printf("command %v -> args %v\n", command, args)
		commandFunction, ok := Commands.GetCommandFunction(command, &connectionStatus, config)
		if !ok {
			_, err = fmt.Fprint(connectionSocket, Replies.CreateReplySyntaxError())
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
			}
			continue
		}
		connectionStatus.SendFTPReply(commandFunction.Execute(args))
	}
}

func connectionHandle(connectionSocket net.Conn, config Configuration.FTPConfig) {
	defer connectionSocket.Close()
	cert, err := tls.LoadX509KeyPair(config.AuthCertFile, config.AuthKeyFile)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Certificate error: %v", err)
	}
	security := Connection.Security{
		IsSecure:    false,
		Certificate: cert,
		CertFile:    config.AuthCertFile,
		KeyFile:     config.AuthKeyFile,
		Config: &tls.Config{
			Certificates:             []tls.Certificate{cert},
			MinVersion:               tls.VersionTLS12,
			PreferServerCipherSuites: true,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			},
			Rand: rand.Reader,
		},
	}
	connectionStatus = Connection.Status{
		Connected:          false,
		StandardConnection: connectionSocket,
		Remote:             connectionSocket.RemoteAddr().String(),
		Authenticated:      false,
		Anonymous:          false,
		User:               "",
		CurrentPath:        config.RootPath,
		TypeCode:           "ASCII",
		FormCode:           "Non-print",
		Type:               Connection.TransferType(Connection.Active),
		Structure:          Connection.Structure(Connection.File),
		Mode:               Connection.TransferMode(Connection.Stream),
		Umask:              config.Umask,
		IdleTimeout:        config.IdleTimeout,
		PreferredEProtocol: 0,
		EPSVAll:            false,
		Security:           security,
	}
	connectionStatus.CommandConnection = connectionStatus.StandardConnection
	connectionStatus.Connect()
	connectionLog(connectionSocket)
	connectionPreamble(connectionSocket, config)
	commandLoop(connectionSocket, config)
}
