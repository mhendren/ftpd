package Connection

import (
	"crypto/tls"
	"net"
	"net/textproto"
)

type DataChannelProtectionLevel int

const (
	Clear int = iota
	Safe
	Confidential
	Private
)

type Security struct {
	IsSecure          bool
	CommandConnection net.Conn
	TextProto         *textproto.Conn
	Certificate       tls.Certificate
	Config            *tls.Config
	CertFile          string
	KeyFile           string
	ProtectedBSize    int
	ProtectionLevel   DataChannelProtectionLevel
	SecurityMechanism SecurityMechanism
}
