package plugin

import (
	"errors"

	"github.com/coreservice-io/cli-template/configuration"
	"github.com/coreservice-io/cli-template/plugin/auto_cert_plugin"
)

func initAutoCert() error {

	auto_cert_crt_path, err := configuration.Config.GetString("auto_cert_crt_path", "")
	if err != nil {
		return errors.New("auto_cert_crt_path [string] in config err," + err.Error())
	}

	auto_cert_key_path, err := configuration.Config.GetString("auto_cert_key_path", "")
	if err != nil {
		return errors.New("auto_cert_key_path [string] in config err," + err.Error())
	}

	auto_cert_url, err := configuration.Config.GetString("auto_cert_url", "")
	if err != nil {
		return errors.New("auto_cert_url [string] in config err," + err.Error())
	}

	auto_cert_check_interval, err := configuration.Config.GetInt("auto_cert_check_interval", 3600)
	if err != nil || auto_cert_check_interval <= 5 {
		return errors.New("auto_cert_check_interval [int64] in config err or too small interval")
	}

	auto_cert_init_download, err := configuration.Config.GetBool("auto_cert_init_download", true)
	if err != nil {
		return errors.New("auto_cert_init_download [bool] in config err or too small interval")
	}

	return auto_cert_plugin.Init(auto_cert_plugin.Config{
		Download_url:        auto_cert_url,
		Local_crt_path:      auto_cert_crt_path,
		Local_key_path:      auto_cert_key_path,
		Check_interval_secs: auto_cert_check_interval,
	}, auto_cert_init_download)

}
