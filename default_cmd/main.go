package default_cmd

import (
	"time"

	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/config"
	"github.com/coreservice-io/cli-template/default_cmd/http"
	"github.com/coreservice-io/cli-template/plugin/auto_cert_plugin"
	"github.com/coreservice-io/cli-template/plugin/geo_ip_plugin"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func StartDefault(clictx *cli.Context) {

	//defer func() {
	//	//global.ReleaseResources()
	//}()
	color.Green(basic.Logo)
	//ini components and run example
	InitComponent()

	//start threads jobs
	go start_jobs()

	start_components()

	basic.Logger.Infoln(geo_ip_plugin.GetInstance().GetInfo("129.146.243.246"))
	basic.Logger.Infoln(geo_ip_plugin.GetInstance().GetInfo("192.168.189.125"))
	basic.Logger.Infoln(geo_ip_plugin.GetInstance().GetInfo("2600:4040:a912:a200:a438:9968:96d9:c3e4"))
	basic.Logger.Infoln(geo_ip_plugin.GetInstance().GetInfo("2600:387:1:809::3a"))

	go func() {
		for {
			basic.Logger.Infoln("running")
			time.Sleep(3 * time.Second)
		}
	}()

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
	if config.Get_config().Toml_config.Auto_cert.Enable {
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
