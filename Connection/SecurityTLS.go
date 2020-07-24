package Connection

import "FTPserver/Replies"

type SecurityTLS struct {
}

func (sm SecurityTLS) SupportedProtections() []DataChannelProtectionLevel {
	return []DataChannelProtectionLevel{
		DataChannelProtectionLevel(Clear),
		DataChannelProtectionLevel(Private),
	}
}

func (sm SecurityTLS) Name() string {
	return "TLS"
}

func (sm SecurityTLS) ADAT(_ []byte) Replies.FTPReply {
	// There is no exchange of security data here. Just ignore whatever they sent.
	return Replies.CreateReplySecurityDataExchangeComplete()
}

func (sm SecurityTLS) MIC(_ []byte) Replies.FTPReply {
	return Replies.CreateReplyCommandProtectionNotSupported()
}

func (sm SecurityTLS) CONF(_ []byte) Replies.FTPReply {
	return Replies.CreateReplyCommandProtectionNotSupported()
}

func (sm SecurityTLS) ENC(_ []byte) Replies.FTPReply {
	return Replies.CreateReplyCommandProtectionNotSupported()
}

func (sm SecurityTLS) DiscloseADAT() bool {
	return false
}

func (sm SecurityTLS) DiscloseMIC() bool {
	return false
}

func (sm SecurityTLS) DiscloseCONF() bool {
	return false
}
func (sm SecurityTLS) DiscloseENC() bool {
	return false
}
