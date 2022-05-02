package ip

import (
	"context"
	"time"

	"github.com/coreservice-io/cli-template/plugin/redis_plugin"
	"github.com/go-redis/redis/v8"
)

const redis_ip_limit_prefix = "ip_limit"

func LimitAction(ip string, action string, secs int) {
	key := redis_plugin.GetInstance().GenKey(redis_ip_limit_prefix, ip, action)
	redis_plugin.GetInstance().Set(context.Background(), key, 1, time.Duration(secs)*time.Second)
}

func HasLimitedAction(ip string, action string) bool {
	key := redis_plugin.GetInstance().GenKey(redis_ip_limit_prefix, ip, action)
	_, err := redis_plugin.GetInstance().Get(context.Background(), key).Result()
	return err == redis.Nil
}