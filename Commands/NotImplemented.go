package Commands

import "FTPserver/Replies"

type NotImplemented struct{}

func (ni NotImplemented) Execute(_ string) Replies.FTPReply {
	return Replies.CreateReplyCommandNotImplemented()
}

func (ni NotImplemented) Name() string {
	return "NotImplemented"
}
