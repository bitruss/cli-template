package main

import (
	"github.com/universe-30/CliAppTemplate/cliCmd"
	"github.com/universe-30/CliAppTemplate/cmd/defaultCmd"
	"github.com/universe-30/CliAppTemplate/cmd/logs"
	"github.com/universe-30/CliAppTemplate/cmd/service"
)

func main() {
	cliCmd.InitLogger()
	cliCmd.ReadArgs()

	switch cliCmd.CmdToDo.CmdName {
	case cliCmd.CMD_NAME_LOG:
		logs.StartLog()
	case cliCmd.CMD_NAME_SERVICE:
		service.RunServiceCmd()
	case cliCmd.CMD_NAME_CONFIG:

	default:
		cliCmd.Logger.Infoln("======== start default app ===")
		defaultCmd.StartDefault()
	}
}
