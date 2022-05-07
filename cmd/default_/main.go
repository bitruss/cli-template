package default_

import (
	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/cmd/default_/http"
	"github.com/coreservice-io/cli-template/cmd/default_/plugin"
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
	http.ServerStart()
}

func start_jobs() {
	//check all services already started
	if !http.ServerCheckStarted() {
		panic("http server not working")
	}

	// //start the auto_cert auto-updating job
	// auto_cert_plugin.GetInstance().AutoUpdate(func(new_crt_str, new_key_str string) {
	// 	//reload server
	// 	sre := http.ServerReloadCert()
	// 	if sre != nil {
	// 		basic.Logger.Errorln("cert change reload echo server error:" + sre.Error())
	// 	}
	// })

	basic.Logger.Infoln("start your jobs below")
}
