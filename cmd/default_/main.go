package default_

import (
	"github.com/coreservice-io/CliAppTemplate/basic"
	"github.com/coreservice-io/CliAppTemplate/cmd/default_/api"
	"github.com/coreservice-io/CliAppTemplate/configuration"
	"github.com/coreservice-io/CliAppTemplate/plugin/echoServer"
	"github.com/coreservice-io/UUtils/path_util"
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
	api.ConfigApi(httpServer)

	//static
	conf_http_static_dir, sd_err := configuration.Config.GetString("http_static_dir", "")
	if sd_err == nil && conf_http_static_dir != "" {
		h_s_d, err := path_util.SmartExistPath(conf_http_static_dir)
		if err == nil {
			httpServer.Static("/", h_s_d)
			basic.Logger.Infoln("http static folder:", h_s_d)
		}
	}

	err := httpServer.Start()
	if err != nil {
		basic.Logger.Errorln(err)
	}
}
