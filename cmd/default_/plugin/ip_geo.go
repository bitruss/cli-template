package plugin

import (
	"errors"

	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/plugin/ip_geo_plugin"
	"github.com/coreservice-io/cli-template/plugin/reference_plugin"
	tool_errors "github.com/coreservice-io/cli-template/tools/errors"
	"github.com/coreservice-io/ip_geo"
	"github.com/coreservice-io/utils/path_util"
)

func initIpGeo() error {

	toml_conf := basic.Get_config().Toml_config

	if toml_conf.Ip_geo.Enable {
		dbFilePath_abs, dbFilePath_abs_exist, _ := path_util.SmartPathExist(toml_conf.Ip_geo.Ip2l.Db_path)
		if !dbFilePath_abs_exist {
			return errors.New("ip2Location db file path error," + toml_conf.Ip_geo.Ip2l.Db_path)
		}
		ip_geo_redis_config := ip_geo.RedisConfig{
			Addr:     toml_conf.Redis.Host,
			UserName: toml_conf.Redis.Username,
			Password: toml_conf.Redis.Password,
			Port:     toml_conf.Redis.Port,
			Prefix:   toml_conf.Redis.Prefix,
			UseTLS:   toml_conf.Redis.Use_tls,
		}

		reference_plugin.Init_("ip_geo")
		return ip_geo_plugin.Init(toml_conf.Ip_geo.Ipstack_key, dbFilePath_abs, toml_conf.Ip_geo.Ip2l.Upgrade_url, int64(toml_conf.Ip_geo.Ip2l.Upgrade_interval),
			reference_plugin.GetInstance_("ip_geo"), ip_geo_redis_config, basic.Logger, tool_errors.PanicHandler)
	}
	return nil
}
