package main

import (
	"github.com/universe-30/CliAppTemplate/boot"
	"github.com/universe-30/CliAppTemplate/cmd/config"
	default_ "github.com/universe-30/CliAppTemplate/cmd/default_"
	"github.com/universe-30/CliAppTemplate/cmd/log"
	"github.com/universe-30/CliAppTemplate/cmd/service"
)

func main() {
	//boot.InitLogger()
	boot.ReadArgs()

	switch boot.CmdToDo.CmdName {
	case boot.CMD_NAME_LOG:
		log.StartLog()
	case boot.CMD_NAME_SERVICE:
		service.RunServiceCmd()
	case boot.CMD_NAME_CONFIG:
		config.ConfigSetting()
	default:
		boot.Logger.Infoln("======== start default app ========")
		default_.StartDefault()
	}
}
