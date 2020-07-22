package Commands

import (
	"FTPserver/Connection"
	"FTPserver/Replies"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type STOR struct {
	cs       *Connection.Status
	isUnique bool
	isAppend bool
}

func (cmd STOR) Execute(args string) Replies.FTPReply {
	defer cmd.cs.DataDisconnect()
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
		if cmd.isAppend {
			file, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666&cmd.cs.Umask)
		} else {
			file, err = os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666&cmd.cs.Umask)
		}
		if file == nil || err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to Create file: %v\n", fileName)
			return Replies.CreateReplyRequestedActionNotTaken()
		}
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
	_, err = cmd.cs.DataConnection.CopyIn(file)
	if err != nil {
		return Replies.CreateReplyConnectionClosedTransferAborted()
	}

	if cmd.isUnique {
		return Replies.CreateReplyFileActionCompleted(file.Name())
	}
	return Replies.CreateReplyClosingDataConnection()
}
