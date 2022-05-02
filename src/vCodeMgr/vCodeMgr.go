package vCodeMgr

import (
	"context"
	"errors"
	"time"

	"github.com/coreservice-io/CliAppTemplate/basic"
	"github.com/coreservice-io/CliAppTemplate/plugin/redis_plugin"
	"github.com/coreservice-io/UUtils/rand_util"
	goredis "github.com/go-redis/redis/v8"
	"github.com/meson-network/mcdn/plugin/redisClient"
)

const coolDownPrefix = "vCodeCoolDown"
const codePrefix = "vCode"

func IsCoolDown(email string) bool {
	key := redis_plugin.GetInstance().GenKey(coolDownPrefix, email)
	_, err := redis_plugin.GetInstance().Get(context.Background(), key).Result()
	if err == goredis.Nil {
		return true
	}
	return false
}

func StartCoolDown(email string) {
	key := redisClient.GetInstance().GenKey(coolDownPrefix, email)
	redisClient.GetInstance().Set(context.Background(), key, 1, 25*time.Second)
}

//send vcode to user
func GenVCode(vCodeKey string) (string, error) {
	key := redisClient.GetInstance().GenKey(codePrefix, vCodeKey)
	code, _ := redisClient.GetInstance().Get(context.Background(), key).Result()
	if code == "" {
		code = rand_util.GenRandStr(4)
	}
	_, err := redisClient.GetInstance().Set(context.Background(), key, code, 4*time.Hour).Result()
	if err != nil {
		basic.Logger.Errorln("GenVCode set email vcode to redis error", "err", err)
		return "", errors.New("set email vcode to redis error")
	}

	basic.Logger.Debugln("vcode", "code", code, "vCodeKey", vCodeKey)

	return code, nil
}

func ValidateVCode(vCodeKey string, code string) bool {
	key := redisClient.GetInstance().GenKey(codePrefix, vCodeKey)
	value, err := redisClient.GetInstance().Get(context.Background(), key).Result()
	if err != nil && err != goredis.Nil {
		basic.Logger.Debugln("ValidateVCode from redis err", "err", err, "vCodeKey", vCodeKey)
		return false
	} else if err == goredis.Nil {
		return false
	}

	if value == code {
		redisClient.GetInstance().Del(context.Background(), key)
		return true
	}
	return false
}
