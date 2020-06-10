package Commands

import "FTPserver/Replies"

type Undefined struct{}

func (ud Undefined) Execute(_ string) Replies.FTPReply {
	return Replies.CreateReplyCommandNotImplementedSuperfluous()
}
