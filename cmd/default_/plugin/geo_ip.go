package plugin

import (
	"errors"

	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/basic/conf"
	"github.com/coreservice-io/cli-template/plugin/geo_ip_plugin"
	"github.com/coreservice-io/utils/path_util"
)

func initGeoIp() error {
	toml_conf := conf.Get_config().Toml_config

	if toml_conf.GeoIp.Enable {
		dbFilePath_abs, dbFilePath_abs_exist, _ := path_util.SmartPathExist(toml_conf.GeoIp.Db_path)
		if !dbFilePath_abs_exist {
			return errors.New("geo_ip db file path error," + toml_conf.GeoIp.Db_path)
		}

		basic.Logger.Infoln("init geo_ip plugin with ",
			"localDbFile:", dbFilePath_abs,
		)

		return geo_ip_plugin.Init(dbFilePath_abs)
	}
	return nil
}
