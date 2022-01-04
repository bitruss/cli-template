package main

import (
	"github.com/universe-30/CliAppTemplate/cliCmd"
	"github.com/universe-30/CliAppTemplate/cmd/config"
	"github.com/universe-30/CliAppTemplate/cmd/defaultCmd"
	"github.com/universe-30/CliAppTemplate/cmd/log"
	"github.com/universe-30/CliAppTemplate/cmd/service"
)

func main() {
	cliCmd.InitLogger()
	cliCmd.ReadArgs()

	switch cliCmd.CmdToDo.CmdName {
	case cliCmd.CMD_NAME_LOG:
		log.StartLog()
	case cliCmd.CMD_NAME_SERVICE:
		service.RunServiceCmd()
	case cliCmd.CMD_NAME_CONFIG:
		config.ConfigSetting()
	default:
		cliCmd.Logger.Infoln("======== start default app ===")
		defaultCmd.StartDefault()
	}
}
