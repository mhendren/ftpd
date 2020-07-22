package Server

import (
	"FTPserver/Configuration"
	"fmt"
	"net"
)

func SetupConnection(ftpConfig Configuration.FTPConfig) net.Listener {
	usePort := ftpConfig.BasePort
	listenSocket, err := net.Listen("tcp", fmt.Sprintf(":%v", usePort))
	if err != nil {
		fmt.Println(fmt.Errorf("error: %s", err))
		return nil
	}
	return listenSocket
}

func Loop(ftpConfig Configuration.FTPConfig, listenSocket net.Listener) {
	for {
		connectionSocket, err := listenSocket.Accept()
		if err != nil {
			// log connection error
		}
		go connectionHandle(connectionSocket, ftpConfig)
	}
}
