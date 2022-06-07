package plugin

import (
	"errors"

	"github.com/coreservice-io/cli-template/configuration"
	"github.com/coreservice-io/cli-template/plugin/mail_plugin"
)

func initSmtpMail() error {
	host, err := configuration.Config.GetString("smtp_host", "smtp.gmail.com")
	if err != nil {
		return errors.New("smtp_host [string] in config err," + err.Error())
	}

	port, err := configuration.Config.GetInt("smtp_port", 578)
	if err != nil {
		return errors.New("smtp_port [int] in config err," + err.Error())
	}

	username, err := configuration.Config.GetString("smtp_username", "username")
	if err != nil {
		return errors.New("smtp_username [string] in config err," + err.Error())
	}

	password, err := configuration.Config.GetString("smtp_password", "password")
	if err != nil {
		return errors.New("smtp_password [string] in config err," + err.Error())
	}

	return mail_plugin.Init(&mail_plugin.Config{
		Host:     host,
		Port:     port,
		Password: password,
		UserName: username,
	})

}
