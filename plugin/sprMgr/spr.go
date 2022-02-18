package sprMgr

import (
	"fmt"

	"github.com/universe-30/CliAppTemplate/basic"
	"github.com/universe-30/RedisSpr"
)

var instanceMap = map[string]*RedisSpr.SprJobMgr{}

func GetDefaultInstance() *RedisSpr.SprJobMgr {
	return instanceMap["default"]
}

func GetInstance(name string) *RedisSpr.SprJobMgr {
	return instanceMap[name]
}

type Config struct {
	Address  string
	UserName string
	Password string
	Port     int
}

// Init a new instance.
//  If only need one instance, use empty name "". Use GetDefaultInstance() to get.
//  If you need several instance, run Init() with different <name>. Use GetInstance(<name>) to get.
func Init(name string, redisConfig Config) error {
	if name == "" {
		name = "default"
	}

	_, exist := instanceMap[name]
	if exist {
		return fmt.Errorf("spr instance <%s> has already initialized", name)
	}

	if redisConfig.Address == "" {
		redisConfig.Address = "127.0.0.1"
	}
	if redisConfig.Port == 0 {
		redisConfig.Port = 6379
	}
	//////// ini spr job //////////////////////

	spr, err := RedisSpr.New(RedisSpr.RedisConfig{
		Addr:     redisConfig.Address,
		Port:     redisConfig.Port,
		Password: redisConfig.Password,
		UserName: redisConfig.UserName,
	})

	if err != nil {
		return err
	}

	spr.SetULogger(basic.Logger)

	instanceMap[name] = spr

	return nil
}
