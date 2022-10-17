package main

import (
	"os"

	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/cmd"
)

func main() {
	basic.InitLogger()

	//config app to run
	errRun := cmd.ConfigCmd().Run(os.Args)
	if errRun != nil {
		basic.Logger.Fatalln(errRun)
	}
}
