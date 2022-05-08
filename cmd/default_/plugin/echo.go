package plugin

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/configuration"
	"github.com/coreservice-io/cli-template/plugin/echo_plugin"
	tool_errors "github.com/coreservice-io/cli-template/tools/errors"
	"github.com/coreservice-io/utils/path_util"
)

func init_http_echo_server() error {

	http_on, _ := configuration.Config.GetBool("http_on", false)
	if http_on {
		http_port, err := configuration.Config.GetInt("http_port", 443)
		if err != nil {
			return errors.New("http_port [int] in config error," + err.Error())
		}

		http_html_dir, http_html_dir_err := configuration.Config.GetString("http_html_dir", "")
		if http_html_dir_err != nil {
			return errors.New("http_html_dir config error," + http_html_dir_err.Error())
		}

		html_file := ""
		if http_html_dir != "" {
			http_html_dir_abs, http_html_dir_abs_err := path_util.SmartExistPath(http_html_dir)
			if http_html_dir_abs_err != nil {
				return errors.New("http_html_dir  error," + http_html_dir_abs_err.Error())
			}
			http_html_dir = http_html_dir_abs
			html_file = filepath.Join(http_html_dir_abs, "index.html")
			_, err := os.Stat(html_file)
			if err != nil {
				return errors.New("index.html file not found inside " + http_html_dir_abs + " folder :")
			}
		}

		return echo_plugin.Init_("http", echo_plugin.Config{Port: http_port, Tls: false, Crt_path: "", Key_path: "",
			Html_dir: http_html_dir, Html_index_path: html_file,
		}, tool_errors.PanicHandler, basic.Logger)
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

		crt_path, cert_path_err := path_util.SmartExistPath(crt)
		if cert_path_err != nil {
			return errors.New("https crt file path error," + cert_path_err.Error())
		}

		key_path, key_path_err := path_util.SmartExistPath(key)
		if cert_path_err != nil {
			return errors.New("https key file path error," + key_path_err.Error())
		}

		https_html_dir, https_html_dir_err := configuration.Config.GetString("https_html_dir", "")
		if https_html_dir_err != nil {
			return errors.New("https_html_dir config error," + https_html_dir_err.Error())
		}

		html_file := ""
		if https_html_dir != "" {
			https_html_dir_abs, https_html_dir_abs_err := path_util.SmartExistPath(https_html_dir)
			if https_html_dir_abs_err != nil {
				return errors.New("https_html_dir  error," + https_html_dir_abs_err.Error())
			}
			https_html_dir = https_html_dir_abs
			html_file = filepath.Join(https_html_dir_abs, "index.html")
			_, err := os.Stat(html_file)
			if err != nil {
				return errors.New("index.html file not found inside " + https_html_dir_abs + " folder :")
			}
		}

		return echo_plugin.Init_("https", echo_plugin.Config{Port: https_port, Tls: true, Crt_path: crt_path, Key_path: key_path,
			Html_dir: https_html_dir, Html_index_path: html_file,
		}, tool_errors.PanicHandler, basic.Logger)
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
