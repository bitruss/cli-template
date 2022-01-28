package sprMgr

import (
	"errors"

	"github.com/universe-30/CliAppTemplate/basic"
	"github.com/universe-30/CliAppTemplate/configuration"
	"github.com/universe-30/RedisSpr"
)

var spr *RedisSpr.SprJobMgr

func GetSingleInstance() *RedisSpr.SprJobMgr {
	return spr
}

func Init() error {
	//////// ini spr job //////////////////////
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

	spr, err = RedisSpr.New(RedisSpr.RedisConfig{
		Addr:     redis_addr,
		Port:     redis_port,
		Password: redis_password,
		UserName: redis_username,
	})

	if err != nil {
		return err
	}

	spr.SetULogger(basic.Logger)

	return nil
}
