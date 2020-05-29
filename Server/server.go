package Server

import (
	"FTPserver/Configuration"
	"fmt"
	"net"
	"os"
	"os/user"
	"path/filepath"
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
	usr, err := user.Current()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return
	}
	ftpConfig = Configuration.FTPConfig{
		AllowAnonymous: true,
		AllowNoLogin:   true,
		RootPath:       filepath.Join(usr.HomeDir, "FTPRoot"),
	}
	for {
		connectionSocket, err := listenSocket.Accept()
		if err != nil {
			// log connection error
		}
		go connectionHandle(connectionSocket, ftpConfig)
	}
}
