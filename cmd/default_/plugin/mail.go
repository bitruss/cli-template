package plugin

// import (
// 	"errors"

// 	"github.com/coreservice-io/cli-template/configuration"
// 	"github.com/coreservice-io/cli-template/plugin/mail_plugin"
// )

// func initSmtpMail() error {
// 	host, err := configuration.Config.GetString("smtp.host", "smtp.gmail.com")
// 	if err != nil {
// 		return errors.New("smtp.host [string] in config err," + err.Error())
// 	}

// 	port, err := configuration.Config.GetInt("smtp.port", 578)
// 	if err != nil {
// 		return errors.New("smtp.port [int] in config err," + err.Error())
// 	}

// 	username, err := configuration.Config.GetString("smtp.username", "username")
// 	if err != nil {
// 		return errors.New("smtp.username [string] in config err," + err.Error())
// 	}

// 	password, err := configuration.Config.GetString("smtp.password", "password")
// 	if err != nil {
// 		return errors.New("smtp.password [string] in config err," + err.Error())
// 	}

// 	return mail_plugin.Init(&mail_plugin.Config{
// 		Host:     host,
// 		Port:     port,
// 		Password: password,
// 		UserName: username,
// 	})

// }
