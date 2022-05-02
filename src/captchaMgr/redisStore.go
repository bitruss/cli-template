package captchaMgr

import (
	"context"
	"strings"
	"time"

	goredis "github.com/go-redis/redis/v8"
	"github.com/meson-network/mcdn/basic"
	"github.com/meson-network/mcdn/plugin/redisClient"
)

const keyPrefix = "captcha"

type RedisStore struct {
}

var store = &RedisStore{}

//set a capt
func (r *RedisStore) Set(id string, value string) error {
	key := redisClient.GetInstance().GenKey(keyPrefix, id)
	err := redisClient.GetInstance().Set(context.Background(), key, value, time.Minute*5).Err()
	if err != nil {
		basic.Logger.Errorln("captcha RedisStore Set error", "err", err, "id", id, "value", value)
		return err
	}
	return nil
}

//get a capt
func (r *RedisStore) get(id string, clear bool) string {
	key := redisClient.GetInstance().GenKey(keyPrefix, id)
	val, err := redisClient.GetInstance().Get(context.Background(), key).Result()
	if err != nil {
		if err != goredis.Nil {
			basic.Logger.Errorln("captcha RedisStore Get error", "err", err, "id", id)
		}
		return ""
	}
	if clear {
		err := redisClient.GetInstance().Del(context.Background(), key).Err()
		if err != nil {
			basic.Logger.Errorln("captcha RedisStore Del error", "err", err, "id", id)
		}
	}
	return val
}

//verify a capt
func (r *RedisStore) Verify(id, answer string, clear bool) bool {
	v := r.get(id, clear)
	v = strings.ToLower(v)
	answer = strings.ToLower(answer)
	return v == answer
}
