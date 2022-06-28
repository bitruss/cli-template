package plugin

import (
	"errors"
	"strconv"

	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/basic/conf"
	"github.com/coreservice-io/cli-template/plugin/ip_local_plugin"
	common_errors "github.com/coreservice-io/cli-template/src/common/errors"
	"github.com/coreservice-io/utils/path_util"
)

func initIpLocal() error {

	toml_conf := conf.Get_config().Toml_config

	if toml_conf.IpLocal.Enable {
		dbFilePath_abs, dbFilePath_abs_exist, _ := path_util.SmartPathExist(toml_conf.IpLocal.Db_path)
		if !dbFilePath_abs_exist {
			return errors.New("ip2Location db file path error," + toml_conf.IpLocal.Db_path)
		}

		if toml_conf.IpLocal.Upgrade_url == "" {
			return errors.New("IpLocal.Upgrade_url not exist")
		}

		if toml_conf.IpLocal.Upgrade_interval < 24*3600 {
			return errors.New("IpLocal.Upgrade_interval can not less than 24*3600")
		}

		basic.Logger.Infoln("init ecs uploader plugin with ",
			"localDbFile:", dbFilePath_abs, "Upgrade_url:", toml_conf.IpLocal.Upgrade_url,
			"upgrade_interval:", strconv.Itoa(toml_conf.IpLocal.Upgrade_interval),
		)

		return ip_local_plugin.Init(dbFilePath_abs, toml_conf.IpLocal.Upgrade_url, int64(toml_conf.IpLocal.Upgrade_interval),
			basic.Logger, common_errors.PanicHandler)
	}
	return nil
}
