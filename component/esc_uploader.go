package component

import (
	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/basic/conf"
	"github.com/coreservice-io/cli-template/plugin/ecs_uploader_plugin"
)

func InitEcsUploader() error {
	toml_conf := conf.Get_config().Toml_config

	if toml_conf.Elastic_search.Enable {

		ecs_uploader_conf := ecs_uploader_plugin.Config{
			Address:  toml_conf.Elastic_search.Host,
			UserName: toml_conf.Elastic_search.Username,
			Password: toml_conf.Elastic_search.Password}

		basic.Logger.Infoln("Init ecs uploader plugin with config:", ecs_uploader_conf)
		if err := ecs_uploader_plugin.Init(&ecs_uploader_conf, basic.Logger); err == nil {
			basic.Logger.Infoln("### InitEcsUploader success")
			return nil
		} else {
			return err
		}
	} else {
		return nil
	}
}
