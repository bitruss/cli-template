package plugin

import (
	"errors"

	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/basic/conf"
	"github.com/coreservice-io/cli-template/plugin/ip_remote_plugin"
	"github.com/coreservice-io/ip_geo/remote"
)

func initIpRemote() error {

	toml_conf := conf.Get_config().Toml_config

	if toml_conf.IpRemote.Enable {
		ip_geo_redis_config := remote.RedisConfig{
			Addr:     toml_conf.IpRemote.Redis.Host,
			UserName: toml_conf.IpRemote.Redis.Username,
			Password: toml_conf.IpRemote.Redis.Password,
			Port:     toml_conf.IpRemote.Redis.Port,
			Prefix:   toml_conf.IpRemote.Redis.Prefix,
			UseTLS:   toml_conf.IpRemote.Redis.Use_tls,
		}

		if toml_conf.IpRemote.Key == "" {
			return errors.New("ip stack key empty")
		}

		basic.Logger.Infoln("init ecs uploader plugin with ipstack_key:", toml_conf.IpRemote.Key,
			"ip_geo_redis_config:", ip_geo_redis_config)

		return ip_remote_plugin.Init(toml_conf.IpRemote.Key, ip_geo_redis_config, basic.Logger)
	}
	return nil
}
