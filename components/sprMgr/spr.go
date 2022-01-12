package sprMgr

import (
	"errors"
	"sync"

	"github.com/universe-30/CliAppTemplate/basic"
	"github.com/universe-30/CliAppTemplate/configuration"
	"github.com/universe-30/RedisSpr"
)

var spr *RedisSpr.SprJobMgr
var once sync.Once

func Init() {
	//only run once
	once.Do(func() {
		var err error = nil
		spr, err = newRedisSpr()
		if err != nil {
			basic.Logger.Fatalln(err)
		}
	})
}

func GetSingleInstance() *RedisSpr.SprJobMgr {
	Init()
	return spr
}

func newRedisSpr() (*RedisSpr.SprJobMgr, error) {

	//////// ini spr job //////////////////////
	redis_addr, _redis_addr_err := configuration.Config.GetString("redis_addr", "127.0.0.1")
	if _redis_addr_err != nil {
		return nil, errors.New("redis_addr [string] in config.json err," + _redis_addr_err.Error())
	}

	redis_username, redis_username_err := configuration.Config.GetString("redis_username", "")
	if redis_username_err != nil {
		return nil, errors.New("redis_username [string] in config.json err," + redis_username_err.Error())
	}

	redis_password, redis_password_err := configuration.Config.GetString("redis_password", "")
	if redis_password_err != nil {
		return nil, errors.New("redis_password [string] in config.json err," + redis_password_err.Error())
	}

	redis_port, redis_port_err := configuration.Config.GetInt("redis_port", 6379)
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

	SprMgr.SetULogger(basic.Logger)

	return SprMgr, nil

}
