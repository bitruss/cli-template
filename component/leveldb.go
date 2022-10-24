package component

import (
	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/basic/conf"
	"github.com/coreservice-io/cli-template/plugin/leveldb_plugin"
)

func InitLevelDB() error {

	toml_conf := conf.Get_config().Toml_config

	if toml_conf.Leveldb.Enable {
		level_db_conf := leveldb_plugin.Config{Db_folder: toml_conf.Leveldb.Path}
		basic.Logger.Infoln("Init leveldb plugin with config:", level_db_conf)
		if err := leveldb_plugin.Init(&level_db_conf); err == nil {
			basic.Logger.Infoln("### InitLevelDB success")
			return nil
		} else {
			return err
		}
	} else {
		return nil
	}
}
