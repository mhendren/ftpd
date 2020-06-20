package SiteCommands

import "FTPserver/Replies"

type Undefined struct {
}

func (ud Undefined) Execute(_ string) Replies.FTPReply {
	return Replies.CreateReplyCommandNotImplementedSuperfluous()
}

func (ud Undefined) Help(_ string) string {
	return "Command explicitly not defined"
}
