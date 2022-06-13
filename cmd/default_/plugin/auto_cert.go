package plugin

import (
	"errors"

	"github.com/coreservice-io/cli-template/configuration"
	"github.com/coreservice-io/cli-template/plugin/auto_cert_plugin"
	"github.com/coreservice-io/utils/path_util"
)

func initAutoCert() error {

	///////
	auto_cert_crt_path, err := configuration.Config.GetString("auto_cert.crt_path", "")
	if err != nil {
		return errors.New("auto_cert.crt_path [string] in config err," + err.Error())
	}
	if auto_cert_crt_path == "" {
		return errors.New("auto_cert.crt_path config error")
	}

	auto_cert_crt_path_abs, auto_cert_crt_path_abs_exist, _ := path_util.SmartPathExist(auto_cert_crt_path)
	if !auto_cert_crt_path_abs_exist {
		return errors.New("auto_cert.crt_path error:" +
			auto_cert_crt_path + ", please check crt file exist on your disk")
	}
	///////

	auto_cert_key_path, err := configuration.Config.GetString("auto_cert.key_path", "")
	if err != nil {
		return errors.New("auto_cert.key_path [string] in config err," + err.Error())
	}
	if auto_cert_key_path == "" {
		return errors.New("auto_cert.key_path config error")
	}
	auto_cert_key_path_abs, auto_cert_key_path_abs_exist, _ := path_util.SmartPathExist(auto_cert_key_path)
	if !auto_cert_key_path_abs_exist {
		return errors.New("auto_cert.key_path error:" + auto_cert_key_path_abs + ",please check key file exist on your disk")
	}
	////////////
	auto_cert_url, err := configuration.Config.GetString("auto_cert.url", "")
	if err != nil {
		return errors.New("auto_cert.url [string] in config err," + err.Error())
	}
	if auto_cert_url == "" {
		return errors.New("auto_cert.url config error")
	}

	auto_cert_check_interval, err := configuration.Config.GetInt("auto_cert.check_interval", 3600)
	if err != nil || auto_cert_check_interval <= 5 {
		return errors.New("auto_cert.check_interval [int64] in config err or too small interval")
	}
	auto_cert_init_download, err := configuration.Config.GetBool("auto_cert.init_download", true)
	if err != nil {
		return errors.New("auto_cert.init_download [bool] in config err or too small interval")
	}

	return auto_cert_plugin.Init(&auto_cert_plugin.Config{
		Download_url:        auto_cert_url,
		Local_crt_path:      auto_cert_crt_path_abs,
		Local_key_path:      auto_cert_key_path_abs,
		Check_interval_secs: auto_cert_check_interval,
	}, auto_cert_init_download)

}
