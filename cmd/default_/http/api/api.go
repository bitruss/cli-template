package api

import (
	"github.com/coreservice-io/cli-template/basic"

	"github.com/coreservice-io/UUtils/path_util"
	_ "github.com/coreservice-io/cli-template/cmd/default_/http/api_docs"
	"github.com/coreservice-io/cli-template/configuration"
	"github.com/coreservice-io/cli-template/plugin/echo_plugin"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/swaggo/swag/gen"
)

// for swagger
// @title           api example
// @version         1.0
// @description     api example
// @termsOfService  https://domain.com
// @contact.name    Support
// @contact.url     https://domain.com
// @contact.email   contact@domain.com

// @license.name    Apache 2.0
// @license.url     http://www.apache.org/licenses/LICENSE-2.0.html

// @schemes         http

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization

func DeclareApi(httpServer *echo_plugin.EchoServer) {
	//health
	config_health(httpServer)
	//user
	config_user(httpServer)
}

func ConfigApi(httpServer *echo_plugin.EchoServer) {
	httpServer.GET("/swagger/*", echoSwagger.WrapHandler)
}

func Gen_Api_Docs() {

	api_doc_gen_search_dir, _ := configuration.Config.GetString("api_doc_gen_search_dir", "")
	api_doc_gen_mainfile, _ := configuration.Config.GetString("api_doc_gen_mainfile", "")
	api_doc_gen_output_dir, _ := configuration.Config.GetString("api_doc_gen_output_dir", "")

	if api_doc_gen_search_dir == "" ||
		api_doc_gen_mainfile == "" ||
		api_doc_gen_output_dir == "" {
		basic.Logger.Errorln("api_doc_gen_search_dir|api_doc_gen_mainfile|api_doc_gen_output_dir config errors")
		return
	}

	api_f, api_f_err := path_util.SmartExistPath(api_doc_gen_search_dir)
	if api_f_err != nil {
		basic.Logger.Errorln("api_doc_gen_search_dir folder not exist")
		return
	}
	api_doc_f, api_doc_f_err := path_util.SmartExistPath(api_doc_gen_output_dir)
	if api_doc_f_err != nil {
		basic.Logger.Errorln("api_doc_gen_output_dir folder not exist")
		return
	}

	config := &gen.Config{
		SearchDir:       api_f,
		OutputDir:       api_doc_f,
		MainAPIFile:     api_doc_gen_mainfile,
		OutputTypes:     []string{"go", "json", "yaml"},
		ParseDependency: true,
	}

	err := gen.New().Build(config)
	if err != nil {
		basic.Logger.Errorln("Gen_Api_Docs", err)
	}

}