package plugin

import (
	"errors"

	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/configuration"
	"github.com/coreservice-io/cli-template/plugin/sqldb_plugin"
)

func initDB() error {
	db_host, err := configuration.Config.GetString("db.host", "127.0.0.1")
	if err != nil {
		return errors.New("db.host [string] in config err," + err.Error())
	}

	db_port, err := configuration.Config.GetInt("db.port", 3306)
	if err != nil {
		return errors.New("db.port [int] in config err," + err.Error())
	}

	db_name, err := configuration.Config.GetString("db.name", "dbname")
	if err != nil {
		return errors.New("db.name [string] in config err," + err.Error())
	}

	db_username, err := configuration.Config.GetString("db.username", "username")
	if err != nil {
		return errors.New("db.username [string] in config err," + err.Error())
	}

	db_password, err := configuration.Config.GetString("db.password", "password")
	if err != nil {
		return errors.New("db.password [string] in config err," + err.Error())
	}

	return sqldb_plugin.Init(sqldb_plugin.Config{
		Host:     db_host,
		Port:     db_port,
		DbName:   db_name,
		UserName: db_username,
		Password: db_password,
	}, basic.Logger)
}
