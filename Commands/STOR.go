package Commands

import (
	"FTPserver/Connection"
	"FTPserver/Replies"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type STOR struct {
	cs       *Connection.Status
	isUnique bool
}

func (cmd STOR) Execute(args string) Replies.FTPReply {
	if cmd.cs.DataConnected {
		defer cmd.cs.DataDisconnect()
	}
	var file *os.File
	var err error
	if cmd.isUnique {
		file, err = ioutil.TempFile(cmd.cs.CurrentPath, "")
		if file == nil || err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to Create unique file\n")
			return Replies.CreateReplyRequestedActionNotTaken()
		}
	} else {
		fileName := filepath.Join(cmd.cs.CurrentPath, args)
		file, err = os.Create(fileName)
		if file == nil || err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to Create file: %v\n", fileName)
			return Replies.CreateReplyRequestedActionNotTaken()
		}
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(file)

	if cmd.cs.DataConnected {
		cmd.cs.SendFTPReply(Replies.CreateReplyDataConnectionOpen())
	} else {
		// Data connection actually needs to be open already, or I don't know where to send data
		return Replies.CreateReplyRequestedFileActionNotAvailable()
	}
	_, err = io.Copy(file, cmd.cs.DataConnection)
	if err != nil {
		return Replies.CreateReplyConnectionClosedTransferAborted()
	}

	if cmd.isUnique {
		return Replies.CreateReplyFileActionCompleted(file.Name())
	}
	return Replies.CreateReplyClosingDataConnection()
}
