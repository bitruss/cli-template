package plugin

import (
	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/basic/conf"
	"github.com/coreservice-io/cli-template/plugin/ecs_plugin"
)

func initElasticSearch() error {
	toml_conf := conf.Get_config().Toml_config

	if toml_conf.Elasticsearch.Enable {

		ecs_conf := ecs_plugin.Config{
			Address:  toml_conf.Elasticsearch.Host,
			UserName: toml_conf.Elasticsearch.Username,
			Password: toml_conf.Elasticsearch.Password}

		basic.Logger.Infoln("init elastic search plugin with config:", ecs_conf)
		return ecs_plugin.Init(&ecs_conf)
	}

	return nil
}
