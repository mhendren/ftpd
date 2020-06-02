package Commands

import (
	"FTPserver/Connection"
	"FTPserver/Replies"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type STOR struct {
	cs *Connection.Status
}

func (cmd STOR) Execute(args string) Replies.FTPReply {
	fileName := filepath.Join(cmd.cs.CurrentPath, args)
	file, err := os.Create(fileName)
	if file == nil || err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to Create file: %v\n", fileName)
		return Replies.CreateReplyRequestedActionNotTaken()
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
	cmd.cs.DataDisconnect()
	if err != nil {
		return Replies.CreateReplyConnectionClosedTransferAborted()
	}
	return Replies.CreateReplyClosingDataConnection()
}
