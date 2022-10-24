package component

import (
	"errors"

	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/basic/conf"
	"github.com/coreservice-io/cli-template/plugin/sqlite_plugin"
)

func InitSqlite() error {
	toml_conf := conf.Get_config().Toml_config

	if toml_conf.Sqlite.Enable {
		sqlite_abs_path, sqlite_abs_path_exist, _ := basic.PathExist(toml_conf.Sqlite.Path)
		if !sqlite_abs_path_exist {
			return errors.New(toml_conf.Sqlite.Path + " :sqlite.path not exist , please reset your sqlite.path :" + toml_conf.Sqlite.Path)
		}

		sqlite_conf := sqlite_plugin.Config{
			Sqlite_abs_path: sqlite_abs_path,
		}

		basic.Logger.Infoln("Init sqlite plugin with config:", sqlite_conf)
		if err := sqlite_plugin.Init(&sqlite_conf, basic.Logger); err == nil {
			basic.Logger.Infoln("### InitSqlite success")
			return nil
		} else {
			return err
		}
	} else {
		return nil
	}
}
