package plugin

import (
	"errors"

	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/configuration"
	"github.com/coreservice-io/cli-template/plugin/echo_plugin"
	tool_errors "github.com/coreservice-io/cli-template/tools/errors"
	"github.com/coreservice-io/utils/path_util"
)

func init_http_echo_server() error {

	http_on, _ := configuration.Config.GetBool("http_on", false)
	if http_on {
		http_port, err := configuration.Config.GetInt("http_port", 80)
		if err != nil {
			return errors.New("http_port [int] in config error," + err.Error())
		}

		return echo_plugin.Init_("http", echo_plugin.Config{Port: http_port, Tls: false, Crt_path: "", Key_path: ""},
			tool_errors.PanicHandler, basic.Logger)
	}
	return nil
}

func init_https_echo_server() error {

	https_on, _ := configuration.Config.GetBool("https_on", false)
	if https_on {
		https_port, err := configuration.Config.GetInt("https_port", 443)
		if err != nil {
			return errors.New("https_port [int] in config error," + err.Error())
		}

		crt, err := configuration.Config.GetString("https_crt_path", "")
		if err != nil || crt == "" {
			return errors.New("https_crt_path [string] in config.json err")
		}

		key, err := configuration.Config.GetString("https_key_path", "")
		if err != nil || key == "" {
			return errors.New("https_key_path [string] in config.json err")
		}

		crt_path, crt_path_exist, _ := path_util.SmartPathExist(crt)
		if !crt_path_exist {
			return errors.New("https crt file path error:" + crt)
		}

		key_path, key_path_exist, _ := path_util.SmartPathExist(key)
		if !key_path_exist {
			return errors.New("https key file path error:" + key)
		}

		return echo_plugin.Init_("https",
			echo_plugin.Config{Port: https_port, Tls: true, Crt_path: crt_path, Key_path: key_path},
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
