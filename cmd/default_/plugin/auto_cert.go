package plugin

import (
	"errors"

	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/basic/conf"
	"github.com/coreservice-io/cli-template/plugin/auto_cert_plugin"
)

func initAutoCert() error {
	toml_conf := conf.Get_config().Toml_config

	if toml_conf.Auto_cert.Enable {

		auto_cert_crt_path_abs, auto_cert_crt_path_abs_exist, _ := basic.PathExist(toml_conf.Auto_cert.Crt_path)
		if !auto_cert_crt_path_abs_exist {
			return errors.New("auto_cert.crt_path error:" +
				toml_conf.Auto_cert.Crt_path + ", please check crt file exist on your disk")
		}

		auto_cert_key_path_abs, auto_cert_key_path_abs_exist, _ := basic.PathExist(toml_conf.Auto_cert.Key_path)
		if !auto_cert_key_path_abs_exist {
			return errors.New("auto_cert.key_path error:" + toml_conf.Auto_cert.Key_path + ",please check key file exist on your disk")
		}

		auto_cert_conf := auto_cert_plugin.Config{
			Download_url:        toml_conf.Auto_cert.Url,
			Local_crt_path:      auto_cert_crt_path_abs,
			Local_key_path:      auto_cert_key_path_abs,
			Check_interval_secs: toml_conf.Auto_cert.Check_interval,
		}

		basic.Logger.Infoln("init auto_cert plugin with config:", auto_cert_conf)

		return auto_cert_plugin.Init(&auto_cert_conf, toml_conf.Auto_cert.Init_download)
	}

	return nil
}
