package Commands

import (
	"FTPserver/Connection"
	"FTPserver/Replies"
	"strconv"
	"strings"
)

type TYPE struct {
	cs *Connection.Status
}

func (cmd TYPE) Execute(args string) Replies.FTPReply {
	getForm := func(arg []string) string {
		if len(arg) > 1 || strings.ToLower(arg[1]) == "t" {
			return "Telnet format effectors"
		} else if len(arg) > 1 && strings.ToLower(arg[1]) == "c" {
			return "Carriage Control"
		}
		return "Non-print"
	}
	arguments := strings.Split(args, " ")
	if len(arguments) < 1 {
		return Replies.CreateReplySyntaxErrorInParameters()
	}
	switch arguments[0] {
	case "I":
		cmd.cs.TypeCode = "Image"
		return Replies.CreateReplyCommandOkay()
	case "A":
		cmd.cs.TypeCode = "ASCII"
		cmd.cs.FormCode = getForm(arguments)
		return Replies.CreateReplyCommandOkay()
	case "E":
		cmd.cs.TypeCode = "EBCDIC"
		cmd.cs.FormCode = getForm(arguments)
		return Replies.CreateReplyCommandOkay()
	case "L":
		cmd.cs.TypeCode = "Local byte Byte size"
		if len(arguments) < 2 {
			return Replies.CreateReplySyntaxErrorInParameters()
		}
		lbs, err := strconv.Atoi(arguments[1])
		if err != nil {
			return Replies.CreateReplySyntaxErrorInParameters()
		}
		cmd.cs.LocalByteSize = uint8(lbs)
		return Replies.CreateReplyCommandOkay()
	}
	return Replies.CreateReplySyntaxErrorInParameters()
}
