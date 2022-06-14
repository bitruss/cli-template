package plugin

import (
	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/plugin/ecs_plugin"
)

func initElasticSearch() error {
	toml_conf := basic.Get_config().Toml_config

	if toml_conf.Elasticsearch.Enable {
		return ecs_plugin.Init(&ecs_plugin.Config{
			Address:  toml_conf.Elasticsearch.Host,
			UserName: toml_conf.Elasticsearch.Username,
			Password: toml_conf.Elasticsearch.Password})
	}

	return nil
}
