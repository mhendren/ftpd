package Commands

import (
	"FTPserver/Configuration"
	"FTPserver/Connection"
	"FTPserver/Replies"
	"reflect"
	"sort"
)

type HELP struct {
	cs     *Connection.Status
	config Configuration.FTPConfig
}

func (cmd HELP) Execute(_ string) Replies.FTPReply {
	commandSet := GetCommandMap(cmd.cs, cmd.config)
	keys := make([]string, 0, len(commandSet))
	for k := range commandSet {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	outStr := "\n    Site recognized commands\n    ('-' are not implemented)\n"
	outCount := 0
	for _, k := range keys {
		vType := reflect.TypeOf(commandSet[k]).String()
		if vType == "Commands.Undefined" {
			continue
		}
		if outCount == 0 {
			outStr += "    "
		}
		command := k
		if vType == "Commands.NotImplemented" {
			command += "-"
		}
		outStr += command
		outCount++
		if outCount < 9 {
			for i := 8; i >= len(command); i-- {
				outStr += " "
			}
		} else {
			outStr += "\n"
			outCount = 0
		}
	}
	outStr += "\nEnd of HELP"
	return Replies.CreateReplyHelpMessage(outStr)
}
