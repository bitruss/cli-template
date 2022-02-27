package redisClient

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-redis/redis/v8"
)

var instanceMap = map[string]*RedisClient{}

type RedisClient struct {
	KeyPrefix string
	*redis.ClusterClient
}

func GetInstance() *RedisClient {
	return instanceMap["default"]
}

func GetInstance_(name string) *RedisClient {
	return instanceMap[name]
}

type Config struct {
	Address   string
	UserName  string
	Password  string
	Port      int
	KeyPrefix string
}

func Init(redisConfig Config) error {
	return Init_("default", redisConfig)
}

// Init a new instance.
//  If only need one instance, use empty name "". Use GetDefaultInstance() to get.
//  If you need several instance, run Init() with different <name>. Use GetInstance(<name>) to get.
func Init_(name string, redisConfig Config) error {
	if name == "" {
		name = "default"
	}

	prefix := strings.TrimSuffix(redisConfig.KeyPrefix, ":")

	if prefix == "" {
		return errors.New("redis key prefix error")
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

	instanceMap[name] = &RedisClient{
		prefix + ":",
		r,
	}
	return nil
}

func (rc *RedisClient) GenKey(keys ...string) string {
	if len(keys) == 0 {
		return rc.KeyPrefix + "emptyKey"
	}
	return rc.KeyPrefix + strings.Join(keys, ":")
}
