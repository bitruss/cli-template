package plugin

import (
	"errors"

	"github.com/coreservice-io/CliAppTemplate/basic"
	"github.com/coreservice-io/CliAppTemplate/configuration"
	"github.com/coreservice-io/CliAppTemplate/plugin/sqldb_plugin"
)

func initDB() error {
	db_host, err := configuration.Config.GetString("db_host", "127.0.0.1")
	if err != nil {
		return errors.New("db_host [string] in config err," + err.Error())
	}

	db_port, err := configuration.Config.GetInt("db_port", 3306)
	if err != nil {
		return errors.New("db_port [int] in config err," + err.Error())
	}

	db_name, err := configuration.Config.GetString("db_name", "dbname")
	if err != nil {
		return errors.New("db_name [string] in config err," + err.Error())
	}

	db_username, err := configuration.Config.GetString("db_username", "username")
	if err != nil {
		return errors.New("db_username [string] in config err," + err.Error())
	}

	db_password, err := configuration.Config.GetString("db_password", "password")
	if err != nil {
		return errors.New("db_password [string] in config err," + err.Error())
	}

	return sqldb_plugin.Init(sqldb_plugin.Config{
		Host:     db_host,
		Port:     db_port,
		DbName:   db_name,
		UserName: db_username,
		Password: db_password,
	}, basic.Logger)
}
