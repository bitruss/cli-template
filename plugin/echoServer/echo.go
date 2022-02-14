package echoServer

import (
	"errors"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/universe-30/CliAppTemplate/basic"
	"github.com/universe-30/CliAppTemplate/configuration"
	"github.com/universe-30/CliAppTemplate/tools"
	"github.com/universe-30/EchoMiddleware"
	"github.com/universe-30/EchoMiddleware/tool"
	"github.com/universe-30/UUtils/path_util"
)

type EchoServer struct {
	*echo.Echo
	Http_port              int
	Http_static_abs_folder string
}

var echoServer *EchoServer

func GetSingleInstance() *EchoServer {
	return echoServer
}

/*
http_port
http_static_rel_folder
*/
func Init() error {
	if echoServer != nil {
		return nil
	}
	http_port, err := configuration.Config.GetInt("http_port", 8080)
	if err != nil {
		return errors.New("http_port [int] in config error," + err.Error())
	}

	http_static_rel_folder, err := configuration.Config.GetString("http_static_rel_folder", "")
	if err != nil {
		return errors.New("http_static_rel_folder [string] in config error," + err.Error())
	}

	echoServer = &EchoServer{
		echo.New(),
		http_port,
		path_util.GetAbsPath(http_static_rel_folder),
	}

	//cros
	echoServer.Use(middleware.CORS())
	//logger
	echoServer.Use(EchoMiddleware.LoggerWithConfig(EchoMiddleware.LoggerConfig{
		Logger:            basic.Logger,
		RecordFailRequest: false,
	}))
	//recover and panicHandler
	echoServer.Use(EchoMiddleware.RecoverWithConfig(EchoMiddleware.RecoverConfig{
		OnPanic: tools.PanicHandler,
	}))
	echoServer.JSONSerializer = tool.NewJsoniter()

	return nil
}

func (s *EchoServer) Start() error {
	basic.Logger.Infoln("http server started on port :" + strconv.Itoa(s.Http_port))
	if s.Http_static_abs_folder != "" {
		basic.Logger.Infoln("http server with static folder:" + s.Http_static_abs_folder)
		s.Echo.Static("/", s.Http_static_abs_folder)
	}

	return s.Echo.Start(":" + strconv.Itoa(s.Http_port))
}

func (s *EchoServer) Close() {
	s.Echo.Close()
}
