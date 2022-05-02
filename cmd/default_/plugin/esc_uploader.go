package plugin

import (
	"errors"

	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/configuration"
	"github.com/coreservice-io/cli-template/plugin/ecs_uploader_plugin"
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

	return ecs_uploader_plugin.Init(ecs_uploader_plugin.Config{
		Address:  elasticSearchAddr,
		UserName: elasticSearchUserName,
		Password: elasticSearchPassword}, basic.Logger)

}
