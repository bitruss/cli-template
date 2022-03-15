package default_

import (
	"os"
	"path/filepath"

	"github.com/coreservice-io/CliAppTemplate/basic"
	"github.com/coreservice-io/CliAppTemplate/cmd/default_/api"
	"github.com/coreservice-io/CliAppTemplate/configuration"
	"github.com/coreservice-io/CliAppTemplate/plugin/echoServer"
	"github.com/coreservice-io/UUtils/path_util"
	"github.com/fatih/color"
	"github.com/labstack/echo/v4/middleware"
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
	config_http_server_static(httpServer)
	httpServer.Start()
}

func config_http_server_static(httpServer *echoServer.EchoServer) {

	http_s_f, _ := configuration.Config.GetString("http_static_rel_folder", "")
	if http_s_f != "" {
		exist, _ := path_util.AbsPathExist(http_s_f)
		if !exist {
			//if user run from root as working directory
			currDir, err := os.Getwd()
			if err == nil {
				w_f := filepath.Join(currDir, http_s_f)
				exist, _ := path_util.AbsPathExist(w_f)
				if exist {
					basic.Logger.Infoln("http server static folder:", w_f)
					httpServer.Use(middleware.Static(w_f))
				}
			}
		} else {
			basic.Logger.Infoln("http server static folder:", http_s_f)
			httpServer.Use(middleware.Static(http_s_f))
		}
	}

}
