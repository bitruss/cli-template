package api

import (
	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/basic/conf"
	"github.com/labstack/echo/v4"

	_ "github.com/coreservice-io/cli-template/cmd/default_/http/api_docs"
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

// @schemes         https

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization

func DeclareApi(httpServer *echo.Echo) {
	//health
	configHealth(httpServer)
	//user
	configUser(httpServer)
}

func ConfigApi(httpServer *echo.Echo) {
	httpServer.GET("/*", echoSwagger.WrapHandler)
}

func GenApiDocs() {

	toml_config := conf.Get_config().Toml_config

	api_f, api_f_exist, _ := basic.PathExist(toml_config.Api.Doc_gen_search_dir)
	if !api_f_exist {
		basic.Logger.Errorln("api_doc_gen_search_dir folder not exist:" + toml_config.Api.Doc_gen_search_dir)
		return
	}
	api_doc_f, api_doc_f_exist, _ := basic.PathExist(toml_config.Api.Doc_gen_output_dir)
	if !api_doc_f_exist {
		basic.Logger.Errorln("api_doc_gen_output_dir folder not exist:" + toml_config.Api.Doc_gen_output_dir)
		return
	}

	if toml_config.Api.Doc_gen_mainfile == "" {
		basic.Logger.Errorln("doc_gen_mainfile can not be empty")
		return
	}

	config := &gen.Config{
		SearchDir:       api_f,
		OutputDir:       api_doc_f,
		MainAPIFile:     toml_config.Api.Doc_gen_mainfile,
		OutputTypes:     []string{"go", "json", "yaml"},
		ParseDependency: true,
	}

	err := gen.New().Build(config)
	if err != nil {
		basic.Logger.Errorln("Gen_Api_Docs", err)
	}

}
