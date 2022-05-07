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

	//start the auto_cert auto-updating job
	// auto_cert_plugin.GetInstance().AutoUpdate(func(new_crt_str, new_key_str string) {
	// 	basic.Logger.Infoln("new_crt_str", new_crt_str)
	// 	basic.Logger.Infoln("new_key_str", new_key_str)

	// 	//reload server
	// 	sre := http.ServerReloadCert()
	// 	if sre != nil {
	// 		basic.Logger.Errorln("cert change reload echo server error:" + sre.Error())
	// 	}

	// 	//restart nginx
	// 	go func() {
	// 		cmd := exec.Command("/bin/bash", "-c", "sudo nginx -s reload")
	// 		err := cmd.Run()
	// 		if err != nil {
	// 			basic.Logger.Errorln("cert change reload nginx error:", err)
	// 		}
	// 	}()
	// })

	basic.Logger.Infoln("start your jobs below")
}
