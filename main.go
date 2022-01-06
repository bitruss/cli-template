package main

import (
	"os"

	"github.com/universe-30/CliAppTemplate/basic"
	"github.com/universe-30/CliAppTemplate/cmd"
)

func main() {

	//config app to run
	errRun := cmd.ConfigCmd().Run(os.Args)
	if errRun != nil {
		basic.Logger.Panicln(errRun)
	}

	//basic.InitLogger()
	// basic.ReadArgs()

	// switch basic.CmdToDo.CmdName {
	// case basic.CMD_NAME_LOG:
	// 	log.StartLog()
	// case basic.CMD_NAME_SERVICE:
	// 	service.RunServiceCmd()
	// case basic.CMD_NAME_CONFIG:
	// 	config.ConfigSetting()
	// default:
	// 	basic.Logger.Infoln("======== start default app ========")
	// 	default_.StartDefault()
	// }
}
