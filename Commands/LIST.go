package Commands

import (
	"FTPserver/Connection"
	"FTPserver/Replies"
	"fmt"
	"os"
)

type LIST struct {
	cs *Connection.Status
}

type LONGDIR struct {
	FileInfo os.FileInfo
}

func (ld LONGDIR) String() string {
	permGet := func(mode os.FileMode) string {
		var out string
		if mode.IsDir() {
			out += "d"
		} else if mode&os.ModeSymlink != 0 {
			out += "l"
		} else {
			out += "-"
		}
		out += mode.Perm().String()[1:]
		return out
	}

	mode := ld.FileInfo.Mode()
	nlink := 1
	size := ld.FileInfo.Size()
	mtime := ld.FileInfo.ModTime()
	name := ld.FileInfo.Name()

	return fmt.Sprintf("%v %v owner group %v %v %v", permGet(mode), nlink, size,
		mtime.Format("Jan 02 2006"), name)
}

func GetLongDirList(path string) ([]LONGDIR, error) {
	f, err := os.OpenFile(path, 0, 0)
	if err != nil {
		return nil, err
	}
	fileInfos, err := f.Readdir(0)
	_ = f.Close()
	dirData := make([]LONGDIR, len(fileInfos))
	for i := 0; i < len(fileInfos); i++ {
		dirData[i] = LONGDIR{FileInfo: fileInfos[i]}
	}
	return dirData, err
}

func (cmd LIST) Execute(_ string) Replies.FTPReply {
	// by default assume LIST of current directory
	fileInfo, err := GetLongDirList(cmd.cs.CurrentPath)
	if err != nil {
		cmd.cs.DataDisconnect()
		Replies.CreateReplyClosingDataConnection()
	}
	_, _ = fmt.Fprint(os.Stderr, "Data List\n")
	if cmd.cs.DataConnected {
		cmd.cs.SendFTPReply(Replies.CreateReplyDataConnectionOpen())
	} else {
		// Data connection actually needs to be open already, or I don't know where to send data
		return Replies.CreateReplyRequestedFileActionNotAvailable()
	}
	for _, info := range fileInfo {
		_, _ = fmt.Fprintf(os.Stderr, "file: %v\n", info)
		_, err := fmt.Fprintf(cmd.cs.DataConnection, "%v\r\n", info)
		if err != nil {
			cmd.cs.DataDisconnect()
			Replies.CreateReplyConnectionClosedTransferAborted()
		}
	}
	cmd.cs.DataDisconnect()
	return Replies.CreateReplyClosingDataConnection()
}
