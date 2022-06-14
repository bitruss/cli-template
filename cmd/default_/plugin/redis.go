package plugin

import (
	"github.com/coreservice-io/cli-template/basic/conf"
	"github.com/coreservice-io/cli-template/plugin/redis_plugin"
)

func initRedis() error {
	toml_conf := conf.Get_config().Toml_config

	if toml_conf.Redis.Enable {
		return redis_plugin.Init(&redis_plugin.Config{
			Address:   toml_conf.Redis.Host,
			UserName:  toml_conf.Redis.Username,
			Password:  toml_conf.Redis.Password,
			Port:      toml_conf.Redis.Port,
			KeyPrefix: toml_conf.Redis.Prefix,
			UseTLS:    toml_conf.Redis.Use_tls,
		})
	}

	return nil
}
