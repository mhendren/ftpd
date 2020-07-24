package Connection

import "FTPserver/Replies"

type SecurityMechanism interface {
	SupportedProtections() []DataChannelProtectionLevel
	Name() string
	ADAT(args []byte) Replies.FTPReply
	MIC(args []byte) Replies.FTPReply
	ENC(args []byte) Replies.FTPReply
	CONF(args []byte) Replies.FTPReply
	DiscloseADAT() bool
	DiscloseMIC() bool
	DiscloseCONF() bool
	DiscloseENC() bool
}
