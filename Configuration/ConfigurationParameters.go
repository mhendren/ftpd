package Configuration

import "FTPserver/Connection"

type FTPConfig struct {
	AllowAnonymous      bool
	AllowNoLogin        bool
	RootPath            string
	BasePort            uint16
	SystemType          string
	Version             string
	SupportedStructures []Connection.Structure
	SupportedModes      []Connection.TransferMode
}
