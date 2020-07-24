package Commands

import "FTPserver/Replies"

type NOOP struct {
}

func (nop NOOP) Execute(_ string) Replies.FTPReply {
	return Replies.CreateReplyCommandOkay()
}

func (cmd NOOP) Name() string {
	return "NOOP"
}
