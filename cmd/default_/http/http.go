package http

import (
	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/cmd/default_/http/api"
	"github.com/coreservice-io/cli-template/configuration"
	"github.com/coreservice-io/cli-template/plugin/echo_plugin"
	"github.com/coreservice-io/utils/path_util"
)

//httpServer example
func ServerStart() {
	httpServer := echo_plugin.GetInstance()
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
		basic.Logger.Fatalln("StartDefaultHttpSever err:", err)
	}

}

func ServerReloadCert() error {
	return echo_plugin.GetInstance().ReloadCert()
}

func ServerCheckStarted() bool {
	return echo_plugin.GetInstance().CheckStarted()
}
