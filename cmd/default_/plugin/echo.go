package plugin

import (
	"errors"

	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/plugin/echo_plugin"
	tool_errors "github.com/coreservice-io/cli-template/tools/errors"
	"github.com/coreservice-io/utils/path_util"
)

func init_http_echo_server() error {

	toml_conf := basic.Get_config().Toml_config

	if toml_conf.Http.Enable {
		return echo_plugin.Init_("http", echo_plugin.Config{Port: toml_conf.Http.Port, Tls: false, Crt_path: "", Key_path: ""},
			tool_errors.PanicHandler, basic.Logger)
	}

	return nil
}

func init_https_echo_server() error {

	toml_conf := basic.Get_config().Toml_config
	if toml_conf.Https.Enable {

		crt_abs_path, crt_path_exist, _ := path_util.SmartPathExist(toml_conf.Https.Crt_path)
		if !crt_path_exist {
			return errors.New("https crt file path error:" + toml_conf.Https.Crt_path)
		}

		key_abs_path, key_path_exist, _ := path_util.SmartPathExist(toml_conf.Https.Key_path)
		if !key_path_exist {
			return errors.New("https key file path error:" + toml_conf.Https.Key_path)
		}

		return echo_plugin.Init_("https", echo_plugin.Config{Port: toml_conf.Https.Port, Tls: true, Crt_path: crt_abs_path, Key_path: key_abs_path},
			tool_errors.PanicHandler, basic.Logger)
	}

	return nil
}

func initEchoServer() error {
	http_err := init_http_echo_server()
	if http_err != nil {
		return http_err
	}
	https_err := init_https_echo_server()
	if https_err != nil {
		return https_err
	}

	return nil
}
