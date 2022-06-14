package default_

import (
	"time"

	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/cmd/default_/http"
	"github.com/coreservice-io/cli-template/cmd/default_/plugin"
	"github.com/coreservice-io/cli-template/plugin/auto_cert_plugin"
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

	for {
		//never quit
		time.Sleep(time.Duration(1) * time.Hour)
	}

}

func start_components() {

	//start the httpserver
	http.ServerStart()

}

func start_jobs() {
	//check all services already started
	if !http.ServerCheckStarted() {
		panic("http server not working")
	}

	// //start the auto_cert auto-updating job
	if basic.Get_config().Toml_config.Auto_cert.Enable {
		auto_cert_plugin.GetInstance().AutoUpdate(func(new_crt_str, new_key_str string) {
			//reload server
			sre := http.ServerReloadCert()
			if sre != nil {
				basic.Logger.Errorln("cert change reload echo server error:" + sre.Error())
			}
		})
	}

	basic.Logger.Infoln("start your jobs below")
}
