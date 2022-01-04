package components

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/universe-30/CliAppTemplate/cliCmd"
	"github.com/universe-30/UUtils/path_util"
)

type EchoServer struct {
	Echo                   *echo.Echo
	Http_port              int
	Http_static_abs_folder string
}

/*
http_port
http_static_rel_folder
*/
func InitEchoServer() (*EchoServer, error) {
	http_port, err := cliCmd.Config.GetInt("http_port", 8080)
	if err != nil {
		return nil, errors.New("http_port [int] in config error," + err.Error())
	}

	http_static_rel_folder, err := cliCmd.Config.GetString("http_static_rel_folder", "")
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

	return esP, nil
}

func (s *EchoServer) Close() {
	s.Echo.Close()
}
