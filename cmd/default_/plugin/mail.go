package plugin

import (
	"github.com/coreservice-io/cli-template/basic/conf"
	"github.com/coreservice-io/cli-template/plugin/mail_plugin"
)

func initSmtpMail() error {
	toml_conf := conf.Get_config().Toml_config

	if toml_conf.Smtp.Enable {
		return mail_plugin.Init(&mail_plugin.Config{
			Host:     toml_conf.Smtp.Host,
			Port:     toml_conf.Smtp.Port,
			Password: toml_conf.Smtp.Password,
			UserName: toml_conf.Smtp.Username,
		})
	}
	return nil
}
