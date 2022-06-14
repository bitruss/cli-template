package plugin

import (
	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/plugin/sqldb_plugin"
)

func initDB() error {
	toml_conf := basic.Get_config().Toml_config

	if toml_conf.Db.Enable {
		return sqldb_plugin.Init(sqldb_plugin.Config{
			Host:     toml_conf.Db.Host,
			Port:     toml_conf.Db.Port,
			DbName:   toml_conf.Db.Name,
			UserName: toml_conf.Db.Username,
			Password: toml_conf.Db.Password,
		}, basic.Logger)
	}

	return nil
}
