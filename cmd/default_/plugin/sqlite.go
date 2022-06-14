package plugin

// import (
// 	"errors"

// 	"github.com/coreservice-io/cli-template/basic"
// 	"github.com/coreservice-io/cli-template/configuration"
// 	"github.com/coreservice-io/cli-template/plugin/sqlite_plugin"
// 	"github.com/coreservice-io/utils/path_util"
// )

// func initSqlite() error {

// 	sf, sf_err := configuration.Config.GetString("sqlite.path", "")
// 	if sf_err != nil || sf == "" {
// 		return errors.New("sqlite.path not configured correctly")
// 	}

// 	sqlite_abs_path, sqlite_abs_path_exist, _ := path_util.SmartPathExist(sf)
// 	if !sqlite_abs_path_exist {
// 		return errors.New(sf + " :sqlite.path not exist , please reset your sqlite.path :" + sf)
// 	}

// 	return sqlite_plugin.Init(&sqlite_plugin.Config{
// 		Sqlite_abs_path: sqlite_abs_path,
// 	}, basic.Logger)
// }
