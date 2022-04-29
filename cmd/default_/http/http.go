package http

import (
	"github.com/coreservice-io/CliAppTemplate/basic"
	"github.com/coreservice-io/CliAppTemplate/cmd/default_/http/api"
	"github.com/coreservice-io/CliAppTemplate/configuration"
	"github.com/coreservice-io/CliAppTemplate/plugin/echoServer"
	"github.com/coreservice-io/UUtils/path_util"
)

//httpServer example
func StartHttpSever() {
	httpServer := echoServer.GetInstance()
	api.ConfigApi(httpServer)
	api.DeclareApi(httpServer)

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
		basic.Logger.Fatalln(err)
	}
}
