package main

import (
	"github.com/universe-30/CliAppTemplate/cli"
	"github.com/universe-30/CliAppTemplate/cmd/defaultCmd"
	"github.com/universe-30/CliAppTemplate/cmd/logs"
)

func main() {
	cli.ReadArgs()

	switch cli.CmdToDo.CmdName {
	case cli.CMD_NAME_LOG:
		logs.StartLog()
	case cli.CMD_NAME_SERVICE:
		//serviceCmd.RunServiceCmd()
	case cli.CMD_NAME_CONFIG:

	default:
		cli.Logger.Infoln("======== start default app ===")
		defaultCmd.StartDefault()
	}
}
