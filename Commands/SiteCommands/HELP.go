package SiteCommands

import (
	"FTPserver/Configuration"
	"FTPserver/Connection"
	"FTPserver/Replies"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

type HELP struct {
	cs     *Connection.Status
	config Configuration.FTPConfig
}

func (siteCommand HELP) Execute(argString string) Replies.FTPReply {
	commandSet := GetCommandMap(siteCommand.cs, siteCommand.config)
	args := strings.Split(argString, " ")
	for _, val := range args {
		strings.TrimSpace(val)
	}
	updatedArgString := strings.Join(args[1:], " ")

	var outStr string
	if updatedArgString == "" {
		keys := make([]string, 0, len(commandSet))
		for k := range commandSet {
			keys = append(keys, k)
		}

		sort.Strings(keys)

		outStr = "\n    Recognized SITE commands\n    ('-' are not implemented)\n"
		outCount := 0
		for _, k := range keys {
			vType := reflect.TypeOf(commandSet[k]).String()
			if vType == "SiteCommands.Undefined" {
				continue
			}
			if outCount == 0 {
				outStr += "    "
			}
			command := k
			if vType == "SiteCommands.NotImplementedSiteCommand" {
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
	} else {
		updatedArgs := strings.Split(updatedArgString, " ")
		siteCmd, ok := commandSet[updatedArgs[0]]
		if !ok {
			outStr = fmt.Sprintf("Unknown SITE command \"%v\"", updatedArgs[0])
		} else {
			outStr = siteCmd.Help(updatedArgString)
		}
	}

	return Replies.CreateReplyHelpMessage(outStr)
}

func (siteCommand HELP) Help(_ string) string {
	return "Syntax: SITE HELP [<sp> site-command]"
}
