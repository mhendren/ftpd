package Commands

import (
	"FTPserver/Configuration"
	"FTPserver/Connection"
)

type ExtendedCommand interface {
	IsExtendedCommand(status *Connection.Status, config Configuration.FTPConfig) bool
	AcceptedArguments(status *Connection.Status, config Configuration.FTPConfig) []string
}
