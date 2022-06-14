package plugin

import (
	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/basic/conf"
	"github.com/coreservice-io/cli-template/plugin/sqldb_plugin"
)

func initDB() error {
	toml_conf := conf.Get_config().Toml_config

	if toml_conf.Db.Enable {
		db_conf := sqldb_plugin.Config{
			Host:     toml_conf.Db.Host,
			Port:     toml_conf.Db.Port,
			DbName:   toml_conf.Db.Name,
			UserName: toml_conf.Db.Username,
			Password: toml_conf.Db.Password,
		}
		basic.Logger.Infoln("init db plugin with config:", db_conf)
		return sqldb_plugin.Init(&db_conf, basic.Logger)
	}

	return nil
}
