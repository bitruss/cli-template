package plugin

import (
	"errors"

	"github.com/coreservice-io/CliAppTemplate/configuration"
	"github.com/coreservice-io/CliAppTemplate/plugin/redis_plugin"
)

func initRedis() error {
	redis_addr, err := configuration.Config.GetString("redis_addr", "127.0.0.1")
	if err != nil {
		return errors.New("redis_addr [string] in config err," + err.Error())
	}

	redis_username, err := configuration.Config.GetString("redis_username", "")
	if err != nil {
		return errors.New("redis_username [string] in config err," + err.Error())
	}

	redis_password, err := configuration.Config.GetString("redis_password", "")
	if err != nil {
		return errors.New("redis_password [string] in config err," + err.Error())
	}

	redis_port, err := configuration.Config.GetInt("redis_port", 6379)
	if err != nil {
		return errors.New("redis_port [int] in config err," + err.Error())
	}

	redis_prefix, err := configuration.Config.GetString("redis_prefix", "")
	if err != nil {
		return errors.New("redis_prefix [string] in config err," + err.Error())
	}

	redis_useTls, err := configuration.Config.GetBool("redis_useTls", false)
	if err != nil {
		return errors.New("redis_useTls [bool] in config err," + err.Error())
	}

	return redis_plugin.Init(redis_plugin.Config{
		Address:   redis_addr,
		UserName:  redis_username,
		Password:  redis_password,
		Port:      redis_port,
		KeyPrefix: redis_prefix,
		UseTLS:    redis_useTls,
	})
}
