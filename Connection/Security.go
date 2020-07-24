package Connection

import (
	"crypto/tls"
	"net/textproto"
)

type DataChannelProtectionLevel int

const (
	Clear int = iota
	Safe
	Confidential
	Private
)

type CCCStatus int

const (
	Forbid = iota
	Allow
)

type Security struct {
	SecureCommandChannel bool
	CommandConnection    *tls.Conn
	TextProto            *textproto.Conn
	Certificate          tls.Certificate
	Config               *tls.Config
	CertFile             string
	KeyFile              string
	ProtectedBSize       int
	ProtectionLevel      DataChannelProtectionLevel
	SecurityMechanism    SecurityMechanism
	CCCStatus            CCCStatus
}
