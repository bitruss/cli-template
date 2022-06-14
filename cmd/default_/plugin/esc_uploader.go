package plugin

import (
	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/basic/conf"
	"github.com/coreservice-io/cli-template/plugin/ecs_uploader_plugin"
)

func initEcsUploader() error {
	toml_conf := conf.Get_config().Toml_config

	if toml_conf.Elasticsearch.Enable {

		ecs_uploader_conf := ecs_uploader_plugin.Config{
			Address:  toml_conf.Elasticsearch.Host,
			UserName: toml_conf.Elasticsearch.Username,
			Password: toml_conf.Elasticsearch.Password}

		basic.Logger.Infoln("init ecs uploader plugin with config:", ecs_uploader_conf)
		return ecs_uploader_plugin.Init(&ecs_uploader_conf, basic.Logger)
	}

	return nil
}
