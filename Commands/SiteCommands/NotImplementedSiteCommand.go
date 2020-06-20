package SiteCommands

import "FTPserver/Replies"

type NotImplementedSiteCommand struct {
}

func (_ NotImplementedSiteCommand) Execute(_ string) Replies.FTPReply {
	return Replies.CreateReplyCommandNotImplemented()
}

func (_ NotImplementedSiteCommand) Help(_ string) string {
	return "The site command has not yet been implemented"
}
