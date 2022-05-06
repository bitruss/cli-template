package plugin

import (
	"errors"

	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/configuration"
	"github.com/coreservice-io/cli-template/plugin/echo_plugin"
	tool_errors "github.com/coreservice-io/cli-template/tools/errors"
	"github.com/coreservice-io/utils/path_util"
)

func initEchoServer() error {
	http_port, err := configuration.Config.GetInt("http_port", 80)
	if err != nil {
		return errors.New("http_port [int] in config error," + err.Error())
	}

	tls, _ := configuration.Config.GetBool("http_tls", false)
	if tls {

		crt, err := configuration.Config.GetString("http_tls_crt", "")
		if err != nil {
			return errors.New("http_tls_crt [string] in config.json err," + err.Error())
		}

		key, err := configuration.Config.GetString("http_tls_key", "")
		if err != nil {
			return errors.New("http_tls_key [string] in config.json err," + err.Error())
		}

		crt_path, cert_path_err := path_util.SmartExistPath(crt)
		if cert_path_err != nil {
			return errors.New("http crt file path error," + cert_path_err.Error())
		}

		key_path, key_path_err := path_util.SmartExistPath(key)
		if cert_path_err != nil {
			return errors.New("http key file path error," + key_path_err.Error())
		}

		return echo_plugin.Init(echo_plugin.Config{Port: http_port, Tls: true, Crt_path: crt_path, Key_path: key_path},
			tool_errors.PanicHandler, basic.Logger)

	} else {
		return echo_plugin.Init(echo_plugin.Config{Port: http_port, Tls: false, Crt_path: "", Key_path: ""},
			tool_errors.PanicHandler, basic.Logger)
	}

}
