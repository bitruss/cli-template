package component

import (
	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/basic/conf"
	"github.com/coreservice-io/cli-template/plugin/redis_plugin"
)

func InitRedis() error {

	toml_conf := conf.Get_config().Toml_config

	if toml_conf.Redis.Enable {
		redis_conf := redis_plugin.Config{
			Address:   toml_conf.Redis.Host,
			UserName:  toml_conf.Redis.Username,
			Password:  toml_conf.Redis.Password,
			Port:      toml_conf.Redis.Port,
			KeyPrefix: toml_conf.Redis.Prefix,
			UseTLS:    toml_conf.Redis.Use_tls,
		}

		basic.Logger.Infoln("Init redis plugin with config:", redis_conf)
		if err := redis_plugin.Init(&redis_conf); err == nil {
			basic.Logger.Infoln("### InitRedis success")
			return nil
		} else {
			return err
		}
	} else {
		return nil
	}
}
