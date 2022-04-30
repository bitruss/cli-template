package default_

import (
	"github.com/coreservice-io/CliAppTemplate/basic"
	"github.com/coreservice-io/CliAppTemplate/cmd/default_/http"
	"github.com/coreservice-io/CliAppTemplate/cmd/default_/plugin"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func StartDefault(clictx *cli.Context) {

	//defer func() {
	//	//global.ReleaseResources()
	//}()
	color.Green(basic.Logo)
	//ini components and run example
	plugin.InitPlugin()

	//start threads jobs
	go start_jobs()

	start_components()
}

func start_components() {
	//start the httpserver
	http.StartDefaultHttpSever()
}

func start_jobs() {
	//check all services already started
	if !http.CheckDefaultHttpServerStarted() {
		panic("http server not working")
	}

	basic.Logger.Infoln("start your jobs below")
}
