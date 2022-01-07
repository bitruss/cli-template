package components

import (
	"errors"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/universe-30/CliAppTemplate/basic"
	"github.com/universe-30/EchoMiddleware"
	"github.com/universe-30/EchoMiddleware/tool"
	"github.com/universe-30/UUtils/path_util"
)

type EchoServer struct {
	*echo.Echo
	Http_port              int
	Http_static_abs_folder string
}

/*
http_port
http_static_rel_folder
*/
func NewEchoServer() (*EchoServer, error) {
	http_port, err := basic.Config.GetInt("http_port", 8080)
	if err != nil {
		return nil, errors.New("http_port [int] in config error," + err.Error())
	}

	http_static_rel_folder, err := basic.Config.GetString("http_static_rel_folder", "")
	if err != nil {
		return nil, errors.New("http_static_rel_folder [string] in config error," + err.Error())
	}

	esP := &EchoServer{
		echo.New(),
		http_port,
		"",
	}
	if http_static_rel_folder != "" {
		esP.Http_static_abs_folder = path_util.GetAbsPath(http_static_rel_folder)
	}

	if basic.Logger != nil {
		esP.Use(EchoMiddleware.LoggerWithConfig(EchoMiddleware.LoggerConfig{
			Logger:            basic.Logger,
			RecordFailRequest: true,
		}))
	}

	return esP, nil
}

//use jsoniter
func (s *EchoServer) UseJsoniter() {
	s.JSONSerializer = tool.NewJsoniter()
}

//set panic handler
func (s *EchoServer) SetPanicHandler(panicHandler func(panic_err interface{})) {
	s.Use(EchoMiddleware.RecoverWithConfig(EchoMiddleware.RecoverConfig{
		OnPanic: panicHandler,
	}))
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
