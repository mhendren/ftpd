package OptionsCommands

import (
	"FTPserver/Configuration"
	"FTPserver/Connection"
)

type OptionsCommand interface {
	OptionsSet(status *Connection.Status, command string, args string) error
}

func GetCommandMap(cs *Connection.Status, config Configuration.FTPConfig) map[string]OptionsCommand {
	return map[string]OptionsCommand{}
}

func GetCommandFunction(command string, cs *Connection.Status, config Configuration.FTPConfig) (OptionsCommand, bool) {
	commandMap := GetCommandMap(cs, config)
	cmd, ok := commandMap[command]
	return cmd, ok
}
