package plugin

import (
	"errors"

	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/configuration"
	"github.com/coreservice-io/cli-template/plugin/spr_plugin"
	"github.com/coreservice-io/redis_spr"
)

func initSpr() error {
	redis_addr, err := configuration.Config.GetString("redis_addr", "127.0.0.1")
	if err != nil {
		return errors.New("redis_addr [string] in config.json err," + err.Error())
	}

	redis_username, err := configuration.Config.GetString("redis_username", "")
	if err != nil {
		return errors.New("redis_username [string] in config.json err," + err.Error())
	}

	redis_password, err := configuration.Config.GetString("redis_password", "")
	if err != nil {
		return errors.New("redis_password [string] in config.json err," + err.Error())
	}

	redis_port, err := configuration.Config.GetInt("redis_port", 6379)
	if err != nil {
		return errors.New("redis_port [int] in config.json err," + err.Error())
	}

	redis_prefix, err := configuration.Config.GetString("redis_prefix", "")
	if err != nil {
		return errors.New("redis_prefix [string] in config err," + err.Error())
	}

	redis_useTls, err := configuration.Config.GetBool("redis_useTls", false)
	if err != nil {
		return errors.New("redis_useTls [bool] in config err," + err.Error())
	}

	return spr_plugin.Init(&redis_spr.RedisConfig{
		Addr:     redis_addr,
		UserName: redis_username,
		Password: redis_password,
		Port:     redis_port,
		Prefix:   redis_prefix,
		UseTLS:   redis_useTls,
	}, basic.Logger)
}
