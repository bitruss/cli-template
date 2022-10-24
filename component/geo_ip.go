package component

import (
	"errors"

	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/basic/conf"
	"github.com/coreservice-io/cli-template/plugin/geo_ip_plugin"
)

func InitGeoIp() error {
	toml_conf := conf.Get_config().Toml_config

	if toml_conf.Geo_ip.Enable {
		dbFilePath_abs, dbFilePath_abs_exist, _ := basic.PathExist(toml_conf.Geo_ip.Db_path)
		if !dbFilePath_abs_exist {
			return errors.New("geo_ip db file path error," + toml_conf.Geo_ip.Db_path)
		}

		basic.Logger.Infoln("Init geo_ip plugin with localDbFile:", dbFilePath_abs)

		if err := geo_ip_plugin.Init(dbFilePath_abs); err == nil {
			basic.Logger.Infoln("### Init geo_ip success")
			return nil
		} else {
			return err
		}
	} else {
		return nil
	}
}
