package plugin

// import (
// 	"errors"

// 	"github.com/coreservice-io/cli-template/configuration"
// 	"github.com/coreservice-io/cli-template/plugin/leveldb_plugin"
// 	"github.com/coreservice-io/utils/path_util"
// )

// func initLevelDB() error {

// 	dbf, dbf_err := configuration.Config.GetString("leveldb.path", "")
// 	if dbf_err != nil || dbf == "" {
// 		return errors.New("leveldb_folder not configured correctly")
// 	}

// 	leveldb_abs_folder_path, leveldb_abs_folder_path_exist, _ := path_util.SmartPathExist(dbf)
// 	if !leveldb_abs_folder_path_exist {
// 		return errors.New(dbf + " :leveldb.path not exist , please reset your leveldb.path :" + dbf)
// 	}

// 	return leveldb_plugin.Init(&leveldb_plugin.Config{
// 		Db_folder: leveldb_abs_folder_path,
// 	})
// }
