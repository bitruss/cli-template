package echo_plugin

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/echo_middleware"
	"github.com/coreservice-io/echo_middleware/tool"
	"github.com/coreservice-io/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var match_echo = map[string]*MatchEcho{}

type MatchEcho struct {
	*echo.Echo
	Name  string
	Match func(string, string) bool //func(host string, req_uri string)
}

func GetMatchEcho(name string) *MatchEcho {
	return match_echo[name]
}

func CheckMatchedEcho(host string, req_uri string) *MatchEcho {
	for _, v := range match_echo {
		if v.Match(host, req_uri) {
			return v
		}
	}
	return nil
}

func InitHostEcho(name string, match func(string, string) bool) (*MatchEcho, error) {
	_, exist := match_echo[name]
	if exist {
		return nil, fmt.Errorf("MatchEcho instance <%s> has already been initialized", name)
	}
	match_echo[name] = &MatchEcho{
		echo.New(),
		name,
		match,
	}
	return match_echo[name], nil
}

type EchoServer struct {
	*echo.Echo
	Logger    log.Logger
	Http_port int
	Tls       bool
	Crt_path  string
	Key_path  string
	Cert      *tls.Certificate
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
*/
type Config struct {
	Port     int
	Tls      bool
	Crt_path string
	Key_path string
}

func Init(serverConfig Config, OnPanicHanlder func(panic_err interface{}), logger log.Logger) error {
	return Init_("default", serverConfig, OnPanicHanlder, logger)
}

// Init a new instance.
//  If only need one instance, use empty name "". Use GetDefaultInstance() to get.
//  If you need several instance, run Init() with different <name>. Use GetInstance(<name>) to get.
func Init_(name string, serverConfig Config, OnPanicHanlder func(panic_err interface{}), logger log.Logger) error {
	if name == "" {
		name = "default"
	}

	_, exist := instanceMap[name]
	if exist {
		return fmt.Errorf("echo server instance <%s> has already been initialized", name)
	}

	if serverConfig.Port == 0 {
		serverConfig.Port = 8080
	}

	echoServer := &EchoServer{
		echo.New(),
		logger,
		serverConfig.Port,
		serverConfig.Tls,
		serverConfig.Crt_path,
		serverConfig.Key_path,
		nil,
	}

	//cros
	echoServer.Use(middleware.CORS())

	//logger
	echoServer.Use(echo_middleware.LoggerWithConfig(echo_middleware.LoggerConfig{
		Logger:            logger,
		RecordFailRequest: false,
	}))
	//recover and panicHandler
	echoServer.Use(echo_middleware.RecoverWithConfig(echo_middleware.RecoverConfig{
		OnPanic: OnPanicHanlder,
	}))

	echoServer.JSONSerializer = tool.NewJsoniter()

	instanceMap[name] = echoServer
	return nil
}

func (s *EchoServer) Start() error {
	if s.Tls {
		cert, err := tls.LoadX509KeyPair(s.Crt_path, s.Key_path)
		if err != nil {
			return err
		}
		s.Cert = &cert
		tlsconf := new(tls.Config)
		tlsconf.GetCertificate = func(*tls.ClientHelloInfo) (*tls.Certificate, error) {
			return s.Cert, nil
		}

		server := http.Server{
			Addr:      ":" + strconv.Itoa(s.Http_port),
			TLSConfig: tlsconf,
		}

		s.Logger.Infoln("https server started on port :" + strconv.Itoa(s.Http_port))
		return s.StartServer(&server)

	} else {
		s.Logger.Infoln("http server started on port :" + strconv.Itoa(s.Http_port))
		return s.Echo.Start(":" + strconv.Itoa(s.Http_port))
	}
}

func (s *EchoServer) ReloadCert() error {
	if s.Tls {
		cert, err := tls.LoadX509KeyPair(s.Crt_path, s.Key_path)
		if err != nil {
			return err
		}

		basic.Logger.Debugln("GetCertificate reloading happend")
		s.Cert = &cert
	}
	return nil
}

func (s *EchoServer) Close() {
	s.Echo.Close()
}

//check the server is indeed up
func (s *EchoServer) CheckStarted() bool {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return false
		case <-ticker.C:
			addr_tls := s.Echo.TLSListenerAddr()
			if addr_tls != nil && strings.Contains(addr_tls.String(), ":") {
				return true
			}
			addr := s.Echo.ListenerAddr()
			if addr != nil && strings.Contains(addr.String(), ":") {
				return true
			}
		}
	}
}
