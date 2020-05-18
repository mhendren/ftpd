package Commands

import "FTPserver/Replies"

type NOOP struct {
}

func (nop NOOP) Execute(args string) Replies.FTPReply {
	return Replies.CreateReplyCommandOkay()
}
