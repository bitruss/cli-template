package plugin

// import (
// 	"errors"

// 	"github.com/coreservice-io/cli-template/configuration"
// 	"github.com/coreservice-io/cli-template/plugin/leveldb_plugin"
// 	"github.com/coreservice-io/utils/path_util"
// )

// func initLevelDB() error {

// 	dbf, dbf_err := configuration.Config.GetString("leveldb_folder", "")
// 	if dbf_err != nil || dbf == "" {
// 		return errors.New("leveldb_folder not configured correctly")
// 	}

// 	leveldb_abs_folder_path, leveldb_abs_folder_path_exist, _ := path_util.SmartPathExist(dbf)
// 	if !leveldb_abs_folder_path_exist {
// 		return errors.New(dbf + " :leveldb_folder not exist , please reset your leveldb_folder :" + dbf)
// 	}

// 	return leveldb_plugin.Init(&leveldb_plugin.Config{
// 		Db_folder: leveldb_abs_folder_path,
// 	})
// }
