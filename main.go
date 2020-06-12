package main

import (
	"FTPserver/Configuration"
	"FTPserver/Connection"
	"FTPserver/Server"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return
	}

	ftpConfig := Configuration.FTPConfig{
		AllowAnonymous:      true,
		AllowNoLogin:        true,
		RootPath:            filepath.Join(usr.HomeDir, "FTPRoot"),
		BasePort:            21,
		Version:             fmt.Sprintf("%v %v", Server.Name, Server.VersionNumber),
		SupportedStructures: []Connection.Structure{Connection.Structure(Connection.File)},
		SupportedModes:      []Connection.TransferMode{Connection.TransferMode(Connection.Stream)},
	}

	server := Server.SetupConnection(ftpConfig)
	Server.Loop(ftpConfig, server)
}
