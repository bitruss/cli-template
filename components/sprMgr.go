package components

import (
	"errors"

	"github.com/universe-30/CliAppTemplate/cliCmd"
	"github.com/universe-30/RedisSpr"
)

func InitSprMgr() (*RedisSpr.SprJobMgr, error) {

	//////// ini spr job //////////////////////
	redis_addr, _redis_addr_err := cliCmd.Config.GetString("redis_addr", "127.0.0.1")
	if _redis_addr_err != nil {
		return nil, errors.New("redis_addr [string] in config.json err," + _redis_addr_err.Error())
	}

	redis_username, redis_username_err := cliCmd.Config.GetString("redis_username", "")
	if redis_username_err != nil {
		return nil, errors.New("redis_username [string] in config.json err," + redis_username_err.Error())
	}

	redis_password, redis_password_err := cliCmd.Config.GetString("redis_password", "")
	if redis_password_err != nil {
		return nil, errors.New("redis_password [string] in config.json err," + redis_password_err.Error())
	}

	redis_port, redis_port_err := cliCmd.Config.GetInt("redis_port", 6379)
	if redis_port_err != nil {
		return nil, errors.New("redis_port [int] in config.json err," + redis_port_err.Error())
	}

	SprMgr, SPR_go_err := RedisSpr.New(RedisSpr.RedisConfig{
		Addr:     redis_addr,
		Port:     redis_port,
		Password: redis_password,
		UserName: redis_username,
	})

	if SPR_go_err != nil {
		return nil, SPR_go_err
	}

	SprMgr.SetULogger(cliCmd.Logger)

	return SprMgr, nil

}
