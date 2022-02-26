package echoServer

import (
	"fmt"
	"strconv"

	"github.com/coreservice-io/CliAppTemplate/basic"
	"github.com/coreservice-io/CliAppTemplate/tools"
	"github.com/coreservice-io/EchoMiddleware"
	"github.com/coreservice-io/EchoMiddleware/tool"
	"github.com/coreservice-io/UUtils/path_util"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type EchoServer struct {
	*echo.Echo
	Http_port              int
	Http_static_abs_folder string
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

func Init(serverConfig Config) error {
	return Init_("default", serverConfig)
}

// Init a new instance.
//  If only need one instance, use empty name "". Use GetDefaultInstance() to get.
//  If you need several instance, run Init() with different <name>. Use GetInstance(<name>) to get.
func Init_(name string, serverConfig Config) error {
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

	echoServer := &EchoServer{
		echo.New(),
		serverConfig.Port,
		path_util.GetAbsPath(serverConfig.StaticFolder),
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

	instanceMap[name] = echoServer
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
