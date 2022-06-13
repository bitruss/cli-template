package plugin

import (
	"errors"

	"github.com/coreservice-io/cli-template/configuration"
	"github.com/coreservice-io/cli-template/plugin/ecs_plugin"
)

func initElasticSearch() error {
	elasticSearchAddr, err := configuration.Config.GetString("elasticsearch.host", "")
	if err != nil {
		return errors.New("elasticsearch.host [string] in config error," + err.Error())
	}

	elasticSearchUserName, err := configuration.Config.GetString("elasticsearch.username", "")
	if err != nil {
		return errors.New("elasticsearch.username_err [string] in config error," + err.Error())
	}

	elasticSearchPassword, err := configuration.Config.GetString("elasticsearch.password", "")
	if err != nil {
		return errors.New("elasticsearch.password [string] in config error," + err.Error())
	}

	return ecs_plugin.Init(&ecs_plugin.Config{
		Address:  elasticSearchAddr,
		UserName: elasticSearchUserName,
		Password: elasticSearchPassword})

}
