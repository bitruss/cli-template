package plugin

import (
	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/plugin/leveldb_plugin"
)

func initLevelDB() error {

	toml_conf := basic.Get_config().Toml_config

	if toml_conf.Leveldb.Enable {
		return leveldb_plugin.Init(&leveldb_plugin.Config{
			Db_folder: toml_conf.Leveldb.Path,
		})
	}
	return nil
}
