package Commands

import (
	"FTPserver/Configuration"
	"FTPserver/Connection"
	"FTPserver/Replies"
	"sort"
	"strings"
)

type FEAT struct {
	cs     *Connection.Status
	config Configuration.FTPConfig
}

func (cmd FEAT) Execute(_ string) Replies.FTPReply {
	commandSet := GetCommandMap(cmd.cs, cmd.config)
	keys := make([]string, 0, len(commandSet))
	for k := range commandSet {
		var x interface{} = commandSet[k]
		command, ok := x.(ExtendedCommand)
		if ok && command.IsExtendedCommand(cmd.cs, cmd.config) {
			keys = append(keys, k+" "+strings.Join(command.AcceptedArguments(cmd.cs, cmd.config), ";"))
		}
	}
	sort.Strings(keys)
	outStr := "Extensions Supported\n"
	for _, k := range keys {
		outStr += " " + k + "\n"
	}
	outStr += "END"
	return Replies.CreateReplyHelpMessage(outStr)
}

func (cmd FEAT) Name() string {
	return "FEAT"
}

func (cmd FEAT) IsExtendedCommand(_ *Connection.Status, _ Configuration.FTPConfig) bool {
	return false
}

func (cmd FEAT) AcceptedArguments(_ *Connection.Status, _ Configuration.FTPConfig) []string {
	return nil
}
