package Commands

import (
	"FTPserver/Configuration"
	"FTPserver/Connection"
	"FTPserver/Replies"
	"fmt"
	"os"
	"path/filepath"
)

type STAT struct {
	cs     *Connection.Status
	config Configuration.FTPConfig
}

func ExecuteSystem(cs *Connection.Status, config Configuration.FTPConfig) Replies.FTPReply {
	message := "FTP server status:\n"
	message += fmt.Sprintf("Version: %v\n", config.Version)
	message += fmt.Sprintf("Connected to: %v\n", cs.CommandConnection.RemoteAddr())
	if cs.Anonymous {
		message += fmt.Sprintf("Logged in anonymous\n")
	} else {
		message += fmt.Sprintf("Logged in as \"%v\"\n", cs.User)
	}
	message += fmt.Sprintf("TYPE: %v, FORM: %v\n", cs.Mode, cs.FormCode)
	message += fmt.Sprintf("Current Path: %v\n", cs.CurrentPath)
	message += fmt.Sprintf("Data Connected: %v\n", cs.DataConnected)

	message += "End of FTP server status"
	return Replies.CreateReplySystemStatus(message)
}

func ExecuteFile(cs *Connection.Status, fileName string) Replies.FTPReply {
	fileInfo, err := os.Stat(filepath.Join(cs.CurrentPath, fileName))
	if err != nil {
		return Replies.CreateReplyRequestedActionNotTakenErrorFilename()
	}
	fileData := LONGDIR{
		FileInfo:    fileInfo,
		abbreviated: false,
	}
	message := fmt.Sprintf("Status of %v\n%v\nEnd of status", fileName, fileData)

	reply := Replies.CreateReplySystemStatus(message)
	reply.Interim = true
	return reply
}

func (cmd STAT) Execute(args string) Replies.FTPReply {
	if args != "" {
		return ExecuteFile(cmd.cs, args)
	}
	return ExecuteSystem(cmd.cs, cmd.config)
}
