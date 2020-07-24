package Commands

import (
	"FTPserver/Connection"
	"FTPserver/Replies"
	"os"
	"path/filepath"
)

type RETR struct {
	cs *Connection.Status
}

func (cmd RETR) Execute(args string) Replies.FTPReply {
	defer cmd.cs.DataDisconnect()

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

	err = cmd.cs.DataConnection.FinishSetup()
	if err != nil {
		return Replies.CreateReplyClosingDataConnection()
	}
	cmd.cs.DataConnect()
	cmd.cs.SendFTPReply(Replies.CreateReplyFileStatusOkay())

	_, err = cmd.cs.DataConnection.CopyOut(file)
	cmd.cs.DataDisconnect()
	if err != nil {
		return Replies.CreateReplyConnectionClosedTransferAborted()
	}
	return Replies.CreateReplyClosingDataConnection()
}

func (cmd RETR) Name() string {
	return "RETR"
}
