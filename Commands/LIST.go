package Commands

import (
	"FTPserver/Connection"
	"FTPserver/Replies"
	"fmt"
	"os"
)

type LIST struct {
	cs          *Connection.Status
	abbreviated bool
}

type LONGDIR struct {
	FileInfo    os.FileInfo
	abbreviated bool
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

	if ld.abbreviated {
		return name
	}

	return fmt.Sprintf("%v %v owner group %v %v %v", permGet(mode), nlink, size,
		mtime.Format("Jan 02 2006"), name)
}

func GetLongDirList(path string, abbreviated bool) ([]LONGDIR, error) {
	f, err := os.OpenFile(path, 0, 0)
	if err != nil {
		return nil, err
	}
	fileInfos, err := f.Readdir(0)
	_ = f.Close()
	dirData := make([]LONGDIR, len(fileInfos))
	for i := 0; i < len(fileInfos); i++ {
		dirData[i] = LONGDIR{FileInfo: fileInfos[i], abbreviated: abbreviated}
	}
	return dirData, err
}

func (cmd LIST) Execute(_ string) Replies.FTPReply {
	if cmd.cs.DataConnected {
		defer cmd.cs.DataDisconnect()
	}
	// by default assume LIST of current directory
	fileInfo, err := GetLongDirList(cmd.cs.CurrentPath, cmd.abbreviated)
	if err != nil {
		cmd.cs.DataDisconnect()
		Replies.CreateReplyClosingDataConnection()
	}
	_, _ = fmt.Fprint(os.Stderr, "Data List\n")
	if cmd.cs.DataConnected {
		cmd.cs.SendFTPReply(Replies.CreateReplyDataConnectionOpen())
	} else {
		return Replies.CreateReplyRequestedFileActionNotAvailable()
	}
	for _, info := range fileInfo {
		_, _ = fmt.Fprintf(os.Stderr, "file: %v\n", info)
		_, err := fmt.Fprintf(cmd.cs.DataConnection, "%v\r\n", info)
		if err != nil {
			return Replies.CreateReplyConnectionClosedTransferAborted()
		}
	}
	return Replies.CreateReplyClosingDataConnection()
}
