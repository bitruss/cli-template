package captchaMgr

import (
	"context"
	"strings"
	"time"

	"github.com/coreservice-io/CliAppTemplate/basic"
	"github.com/coreservice-io/CliAppTemplate/plugin/redis_plugin"
	"github.com/coreservice-io/UCaptcha/math"
	"github.com/coreservice-io/UUtils/rand_util"
	goredis "github.com/go-redis/redis/v8"
)

const coolDownPrefix = "captcha_cool_down"
const keyPrefix = "captcha"

func IsCoolDown(ip string) bool {
	key := redis_plugin.GetInstance().GenKey(coolDownPrefix, ip)
	_, err := redis_plugin.GetInstance().Get(context.Background(), key).Result()
	if err == goredis.Nil {
		return true
	}
	return false
}

func StartCoolDown(ip string) {
	key := redis_plugin.GetInstance().GenKey(coolDownPrefix, ip)
	redis_plugin.GetInstance().Set(context.Background(), key, 1, 2*time.Second)
}

func GenCaptcha() (id, b64s string, err error) {
	keystr, img_encode64_str := math.Gen_svg_base64_prefix(400, 100, "#606060")

	//gen id
	id = rand_util.GenRandStr(24)
	err = set(id, keystr)
	if err != nil {
		return "", "", err
	}

	return id, img_encode64_str, nil
}

func VerifyCaptcha(id, captchaCode string) bool {
	if id == "" || captchaCode == "" {
		return false
	}
	if verify(id, captchaCode, true) {
		return true
	} else {
		return false
	}
}

func set(id string, value string) error {
	key := redis_plugin.GetInstance().GenKey(keyPrefix, id)
	err := redis_plugin.GetInstance().Set(context.Background(), key, value, time.Minute*5).Err()
	if err != nil {
		basic.Logger.Errorln("captcha RedisStore Set error", "err", err, "id", id, "value", value)
		return err
	}
	return nil
}

//get a capt
func get(id string, clear bool) string {
	key := redis_plugin.GetInstance().GenKey(keyPrefix, id)
	val, err := redis_plugin.GetInstance().Get(context.Background(), key).Result()
	if err != nil {
		if err != goredis.Nil {
			basic.Logger.Errorln("captcha RedisStore Get error", "err", err, "id", id)
		}
		return ""
	}
	if clear {
		err := redis_plugin.GetInstance().Del(context.Background(), key).Err()
		if err != nil {
			basic.Logger.Errorln("captcha RedisStore Del error", "err", err, "id", id)
		}
	}
	return val
}

//verify a capt
func verify(id, answer string, clear bool) bool {
	v := get(id, clear)
	v = strings.ToLower(v)
	answer = strings.ToLower(answer)
	return v == answer
}
