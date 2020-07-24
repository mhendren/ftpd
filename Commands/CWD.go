package Commands

import (
	"FTPserver/Connection"
	"FTPserver/Replies"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type CWD struct {
	cs     *Connection.Status
	isCDUP bool
}

func (cmd CWD) Execute(args string) Replies.FTPReply {
	if cmd.isCDUP {
		args = ".."
	}
	argsArray := strings.Split(args, " ")
	if len(argsArray) != 1 {
		return Replies.CreateReplySyntaxErrorInParameters()
	}
	dir := argsArray[0]
	_, _ = fmt.Fprintf(os.Stderr, "CWD %v\n", dir)
	var newDir string
	if filepath.IsAbs(dir) {
		newDir = filepath.Clean(dir)
	} else {
		newDir = filepath.Clean(filepath.Join(cmd.cs.CurrentPath, dir))
	}
	_, _ = fmt.Fprintf(os.Stderr, "CWD: newDir \"%v\"\n", newDir)
	cmd.cs.SetPath(newDir)
	return Replies.CreateReplyCommandOkay()
}

func (cmd CWD) Name() string {
	return "CWD"
}
