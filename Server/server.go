package Server

import (
	"FTPserver/Configuration"
	"fmt"
	"net"
)

var ftpConfig Configuration.FTPConfig

func SetupConnection() net.Listener {
	listenSocket, err := net.Listen("tcp", ":21")
	if err != nil {
		fmt.Println(fmt.Errorf("error: %s", err))
		return nil
	}
	return listenSocket
}

func Loop(listenSocket net.Listener) {
	ftpConfig = Configuration.FTPConfig{
		AllowAnonymous: true,
		AllowNoLogin:   true,
	}
	for {
		connectionSocket, err := listenSocket.Accept()
		if err != nil {
			// log connection error
		}
		go connectionHandle(connectionSocket, ftpConfig)
	}
}
