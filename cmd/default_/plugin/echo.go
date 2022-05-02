package plugin

import (
	"errors"

	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/configuration"
	"github.com/coreservice-io/cli-template/plugin/echo_plugin"
	tool_errors "github.com/coreservice-io/cli-template/tools/errors"
)

func initEchoServer() error {
	http_port, err := configuration.Config.GetInt("http_port", 8080)
	if err != nil {
		return errors.New("http_port [int] in config error," + err.Error())
	}

	return echo_plugin.Init(echo_plugin.Config{Port: http_port}, tool_errors.PanicHandler, basic.Logger)
}
