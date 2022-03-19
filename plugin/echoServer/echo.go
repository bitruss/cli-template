package echoServer

import (
	"fmt"
	"strconv"

	"github.com/coreservice-io/EchoMiddleware"
	"github.com/coreservice-io/EchoMiddleware/tool"
	"github.com/coreservice-io/ULog"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type EchoServer struct {
	*echo.Echo
	Http_port              int
	Http_static_abs_folder string
	Logger                 ULog.Logger
}

var instanceMap = map[string]*EchoServer{}

func GetInstance() *EchoServer {
	return instanceMap["default"]
}

func GetInstance_(name string) *EchoServer {
	return instanceMap[name]
}

/*
http_port
http_static_rel_folder
*/
type Config struct {
	Port         int
	StaticFolder string
}

func Init(serverConfig Config, OnPanicHanlder func(panic_err interface{}), logger ULog.Logger) error {
	return Init_("default", serverConfig, OnPanicHanlder, logger)
}

// Init a new instance.
//  If only need one instance, use empty name "". Use GetDefaultInstance() to get.
//  If you need several instance, run Init() with different <name>. Use GetInstance(<name>) to get.
func Init_(name string, serverConfig Config, OnPanicHanlder func(panic_err interface{}), logger ULog.Logger) error {
	if name == "" {
		name = "default"
	}

	_, exist := instanceMap[name]
	if exist {
		return fmt.Errorf("echo server instance <%s> has already initialized", name)
	}

	if serverConfig.Port == 0 {
		serverConfig.Port = 8080
	}

	if serverConfig.StaticFolder != "" {
		logger.Infoln("http server with static folder :", serverConfig.StaticFolder)
	}

	echoServer := &EchoServer{
		echo.New(),
		serverConfig.Port,
		serverConfig.StaticFolder,
		logger,
	}

	//cros
	echoServer.Use(middleware.CORS())
	//logger
	echoServer.Use(EchoMiddleware.LoggerWithConfig(EchoMiddleware.LoggerConfig{
		Logger:            logger,
		RecordFailRequest: false,
	}))
	//recover and panicHandler
	echoServer.Use(EchoMiddleware.RecoverWithConfig(EchoMiddleware.RecoverConfig{
		OnPanic: OnPanicHanlder,
	}))
	echoServer.JSONSerializer = tool.NewJsoniter()

	instanceMap[name] = echoServer
	return nil
}

func (s *EchoServer) Start() error {
	s.Logger.Infoln("http server started on port :" + strconv.Itoa(s.Http_port))
	return s.Echo.Start(":" + strconv.Itoa(s.Http_port))
}

func (s *EchoServer) StaticWeb() {
	//static
	s.Use(middleware.Static(s.Http_static_abs_folder))
}

func (s *EchoServer) Close() {
	s.Echo.Close()
}
