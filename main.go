package main

import (
	"os"

	"github.com/coreservice-io/CliAppTemplate/basic"
	"github.com/coreservice-io/CliAppTemplate/cmd"
)

func main() {
	basic.InitLogger()

	//config app to run
	errRun := cmd.ConfigCmd().Run(os.Args)
	if errRun != nil {
		basic.Logger.Panicln(errRun)
	}
}
