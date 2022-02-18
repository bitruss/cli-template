package redisClient

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var instanceMap = map[string]*redis.ClusterClient{}

func GetDefaultInstance() *redis.ClusterClient {
	return instanceMap["default"]
}

func GetInstance(name string) *redis.ClusterClient {
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
		return fmt.Errorf("redis instance <%s> has already initialized", name)
	}

	if redisConfig.Address == "" {
		redisConfig.Address = "127.0.0.1"
	}
	if redisConfig.Port == 0 {
		redisConfig.Port = 6379
	}

	r := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{redisConfig.Address + ":" + strconv.Itoa(redisConfig.Port)},
		Username: redisConfig.UserName,
		Password: redisConfig.Password,
	})

	_, err := r.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	instanceMap[name] = r
	return nil
}
