package Commands

import (
	"FTPserver/Connection"
	"FTPserver/Replies"
	"io"
	"os"
	"path/filepath"
)

type RETR struct {
	cs *Connection.Status
}

func (cmd RETR) Execute(args string) Replies.FTPReply {
	if cmd.cs.DataConnected {
		defer cmd.cs.DataDisconnect()
	}

	if args == "" {
		return Replies.CreateReplySyntaxErrorInParameters()
	}
	file, err := os.Open(filepath.Join(cmd.cs.CurrentPath, args))
	if err != nil {
		return Replies.CreateReplyRequestedActionNotTaken()
	}

	defer func(f *os.File) {
		_ = f.Close()
	}(file)

	if cmd.cs.DataConnected {
		cmd.cs.SendFTPReply(Replies.CreateReplyDataConnectionOpen())
	} else {
		return Replies.CreateReplyRequestedFileActionNotAvailable()
	}

	_, err = io.Copy(cmd.cs.DataConnection, file)
	cmd.cs.DataDisconnect()
	if err != nil {
		return Replies.CreateReplyConnectionClosedTransferAborted()
	}
	return Replies.CreateReplyClosingDataConnection()
}
