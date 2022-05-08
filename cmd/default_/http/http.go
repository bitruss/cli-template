package http

import (
	"strconv"
	"strings"

	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/cmd/default_/http/api"
	"github.com/coreservice-io/cli-template/plugin/echo_plugin"
	"github.com/labstack/echo/v4"
)

//httpServer example
func ServerStart() {

	//config https server
	https_srv := echo_plugin.GetInstance_("https")
	if https_srv != nil {
		api.ConfigApi(https_srv)
		api.DeclareApi(https_srv)
		//html folder
		if https_srv.Html_index_path != "" {
			https_srv.Static("/", https_srv.Html_dir)
		}
		//error handler
		conf_error_handler(https_srv)

		go func() {
			err := https_srv.Start()
			if err != nil {
				basic.Logger.Fatalln("default https echo server start err:", err)
			}
		}()
	}

	//config http server
	//http just redirect to https
	http_srv := echo_plugin.GetInstance_("http")
	if http_srv != nil {

		if https_srv != nil {
			//redirect
			http_srv.Any("/*", func(ctx echo.Context) error {
				return ctx.Redirect(301, "https://"+ctx.Request().Host+":"+strconv.Itoa(https_srv.Http_port)+ctx.Request().URL.String())
			})
		}

		go func() {
			err := http_srv.Start()
			if err != nil {
				basic.Logger.Fatalln("default http echo server start err:", err)
			}
		}()
	}

}

func conf_error_handler(server *echo_plugin.EchoServer) {
	server.HTTPErrorHandler = func(err error, ctx echo.Context) {
		if strings.HasPrefix(ctx.Request().URL.String(), "/api") ||
			strings.HasPrefix(ctx.Request().URL.String(), "/swagger") {

		} else {
			if server.Html_index_path != "" {
				ctx.File(server.Html_index_path)
				return
			}
		}

		ctx.HTML(500, err.Error())
	}
}

func ServerReloadCert() error {
	return echo_plugin.GetInstance_("https").ReloadCert()
}

func ServerCheckStarted() bool {

	http_srv := echo_plugin.GetInstance_("http")
	if http_srv != nil {
		if !http_srv.CheckStarted() {
			return false
		}
	}

	https_srv := echo_plugin.GetInstance_("https")
	if https_srv != nil {
		if !https_srv.CheckStarted() {
			return false
		}
	}

	return true
}
