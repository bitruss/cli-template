package echoServer

import (
	"errors"
	"strconv"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/universe-30/CliAppTemplate/basic"
	"github.com/universe-30/CliAppTemplate/configuration"
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
var once sync.Once

func Init() {
	//only run once
	once.Do(func() {
		var err error = nil
		echoServer, err = newEchoServer()
		if err != nil {
			basic.Logger.Fatalln(err)
		}
	})
}

func GetSingleInstance() *EchoServer {
	Init()
	return echoServer
}

/*
http_port
http_static_rel_folder
*/
func newEchoServer() (*EchoServer, error) {
	http_port, err := configuration.Config.GetInt("http_port", 8080)
	if err != nil {
		return nil, errors.New("http_port [int] in config error," + err.Error())
	}

	http_static_rel_folder, err := configuration.Config.GetString("http_static_rel_folder", "")
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

	esP.Use(EchoMiddleware.LoggerWithConfig(EchoMiddleware.LoggerConfig{
		Logger:            basic.Logger,
		RecordFailRequest: true,
	}))

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
