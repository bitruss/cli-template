package plugin

import (
	"errors"

	"github.com/coreservice-io/cli-template/configuration"
	"github.com/coreservice-io/cli-template/plugin/redis_plugin"
)

func initRedis() error {
	redis_addr, err := configuration.Config.GetString("redis.host", "127.0.0.1")
	if err != nil {
		return errors.New("redis.host [string] in config err," + err.Error())
	}

	redis_username, err := configuration.Config.GetString("redis.username", "")
	if err != nil {
		return errors.New("redis.username [string] in config err," + err.Error())
	}

	redis_password, err := configuration.Config.GetString("redis.password", "")
	if err != nil {
		return errors.New("redis.password [string] in config err," + err.Error())
	}

	redis_port, err := configuration.Config.GetInt("redis.port", 6379)
	if err != nil {
		return errors.New("redis.port [int] in config err," + err.Error())
	}

	redis_prefix, err := configuration.Config.GetString("redis.prefix", "")
	if err != nil {
		return errors.New("redis.prefix [string] in config err," + err.Error())
	}

	redis_useTls, err := configuration.Config.GetBool("redis.use_tls", false)
	if err != nil {
		return errors.New("redis.use_tls [bool] in config err," + err.Error())
	}

	return redis_plugin.Init(&redis_plugin.Config{
		Address:   redis_addr,
		UserName:  redis_username,
		Password:  redis_password,
		Port:      redis_port,
		KeyPrefix: redis_prefix,
		UseTLS:    redis_useTls,
	})
}
