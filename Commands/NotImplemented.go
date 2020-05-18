package Commands

import "FTPserver/Replies"

type NotImplemented struct{}

func (ni NotImplemented) Execute(args string) Replies.FTPReply {
	return Replies.CreateReplyCommandNotImplemented()
}
