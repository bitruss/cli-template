package plugin

import (
	"errors"

	"github.com/coreservice-io/CliAppTemplate/basic"
	"github.com/coreservice-io/CliAppTemplate/configuration"
	"github.com/coreservice-io/CliAppTemplate/plugin/ecsUploader"
)

func initEcsUploader() error {
	elasticSearchAddr, err := configuration.Config.GetString("elasticsearch_addr", "")
	if err != nil {
		return errors.New("elasticsearch_addr [string] in config error," + err.Error())
	}

	elasticSearchUserName, err := configuration.Config.GetString("elasticsearch_username", "")
	if err != nil {
		return errors.New("elasticsearch_username_err [string] in config error," + err.Error())
	}

	elasticSearchPassword, err := configuration.Config.GetString("elasticsearch_password", "")
	if err != nil {
		return errors.New("elasticsearch_password [string] in config error," + err.Error())
	}

	return ecsUploader.Init(ecsUploader.Config{
		Address:  elasticSearchAddr,
		UserName: elasticSearchUserName,
		Password: elasticSearchPassword}, basic.Logger)

}
