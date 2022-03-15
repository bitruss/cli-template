package default_

import (
	"github.com/coreservice-io/CliAppTemplate/basic"
	"github.com/coreservice-io/CliAppTemplate/cmd/default_/api"
	"github.com/coreservice-io/CliAppTemplate/configuration"
	"github.com/coreservice-io/CliAppTemplate/plugin/echoServer"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func StartDefault(clictx *cli.Context) {

	color.Green(basic.Logo)
	//basic.Logger.Infoln("hello world , this cli app")

	//ini components and run example
	initComponent()

	//defer func() {
	//	//global.ReleaseResources()
	//}()
	start_http_sever()
}

//httpServer example
func start_http_sever() {
	httpServer := echoServer.GetInstance()
	api.DeclareApi(httpServer)
	http_api, _ := configuration.Config.GetBool("http_api", false)
	if http_api {
		api.ConfigApi(httpServer)
	} else {
		httpServer.StaticWeb()
	}
	httpServer.Start()
}
