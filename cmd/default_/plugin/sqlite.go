package plugin

import (
	"errors"

	"github.com/coreservice-io/UUtils/path_util"
	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/configuration"
	"github.com/coreservice-io/cli-template/plugin/sqlite_plugin"
)

func initSqlite() error {

	sf, sf_err := configuration.Config.GetString("sqlite_path", "")
	if sf_err != nil || sf == "" {
		return errors.New("sqlite_path not configured correctly")
	}

	sqlite_path, sqlite_path_err := path_util.SmartExistPath(sf)
	if sqlite_path_err != nil {
		return errors.New(sf + " :sqlite_path not exist , please reset your sqlite_path ")
	}

	return sqlite_plugin.Init(sqlite_plugin.Config{
		Sqlite_path: sqlite_path,
	}, basic.Logger)
}
