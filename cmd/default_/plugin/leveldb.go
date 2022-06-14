package plugin

import (
	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/basic/conf"
	"github.com/coreservice-io/cli-template/plugin/leveldb_plugin"
)

func initLevelDB() error {

	toml_conf := conf.Get_config().Toml_config

	if toml_conf.Leveldb.Enable {
		level_db_conf := leveldb_plugin.Config{Db_folder: toml_conf.Leveldb.Path}
		basic.Logger.Infoln("init leveldb plugin with config:", level_db_conf)
		return leveldb_plugin.Init(&level_db_conf)
	}
	return nil
}
