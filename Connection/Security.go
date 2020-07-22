package Connection

import (
	"crypto/tls"
	"net"
	"net/textproto"
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
}
