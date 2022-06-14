package plugin

import (
	"errors"

	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/plugin/sqlite_plugin"
	"github.com/coreservice-io/utils/path_util"
)

func initSqlite() error {
	toml_conf := basic.Get_config().Toml_config

	if toml_conf.Sqlite.Enable {
		sqlite_abs_path, sqlite_abs_path_exist, _ := path_util.SmartPathExist(toml_conf.Sqlite.Path)
		if !sqlite_abs_path_exist {
			return errors.New(toml_conf.Sqlite.Path + " :sqlite.path not exist , please reset your sqlite.path :" + toml_conf.Sqlite.Path)
		}

		return sqlite_plugin.Init(&sqlite_plugin.Config{
			Sqlite_abs_path: sqlite_abs_path,
		}, basic.Logger)
	}

	return nil
}
