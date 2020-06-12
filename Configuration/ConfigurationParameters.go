package Configuration

import (
	"FTPserver/Connection"
	"FTPserver/Validation"
)

type FTPConfig struct {
	AllowAnonymous      bool
	AllowNoLogin        bool
	AllowAccount        bool
	RequiresUser        bool
	RequiresPassword    bool
	RequiresAccount     bool
	RootPath            string
	BasePort            uint16
	SystemType          string
	Version             string
	AnonymousUsers      []string
	SupportedStructures []Connection.Structure
	SupportedModes      []Connection.TransferMode
	AccountValidator    Validation.AccountValidator
	PasswordValidator   Validation.PasswordValidator
}
