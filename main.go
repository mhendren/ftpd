package main

import (
	"FTPserver/Configuration"
	"FTPserver/Connection"
	"FTPserver/Server"
	"FTPserver/Validation/Implemntation"
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
		AllowNoLogin:        false,
		AllowAccount:        false,
		RequiresUser:        true,
		RequiresPassword:    true,
		RequiresAccount:     false,
		RootPath:            filepath.Join(usr.HomeDir, "FTPRoot"),
		BasePort:            21,
		Version:             fmt.Sprintf("%v %v", Server.Name, Server.VersionNumber),
		SupportedStructures: []Connection.Structure{Connection.Structure(Connection.File)},
		SupportedModes:      []Connection.TransferMode{Connection.TransferMode(Connection.Stream)},
		AccountValidator:    nil,
		AnonymousUsers:      []string{"anonymous", "ftp"},
		Umask:               os.FileMode(0022),
		IdleTimeout:         60,
		AuthCertFile:        "sample-cert.pem",
		AuthKeyFile:         "sample-key.pem",
		CCCStatus:           Connection.CCCStatus(Connection.Forbid),
	}
	ftpConfig.PasswordValidator = Implemntation.AnonymousOnlyValidator{
		AllowAnonymous: true,
		AnonymousUsers: ftpConfig.AnonymousUsers,
	}
	server := Server.SetupConnection(ftpConfig)
	Server.Loop(ftpConfig, server)
}
