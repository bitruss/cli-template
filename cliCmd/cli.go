package cliCmd

import (
	"os"
)

func ReadArgs() {
	//config app to run
	errRun := configCliCmd().Run(os.Args)
	if errRun != nil {
		Logger.Panicln(errRun)
	}
}
