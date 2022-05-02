package captchaMgr

import (
	"context"
	"time"

	"github.com/coreservice-io/CliAppTemplate/plugin/redis_plugin"
	"github.com/coreservice-io/UCaptcha/math"
	"github.com/coreservice-io/UUtils/rand_util"
	goredis "github.com/go-redis/redis/v8"
)

const coolDownPrefix = "captcha_cool_down"

func IsCaptchaCoolDown(ip string) bool {
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
	err = store.Set(id, keystr)
	if err != nil {
		return "", "", err
	}

	return id, img_encode64_str, nil
}

func VerifyCaptcha(id, captchaCode string) bool {
	if id == "" || captchaCode == "" {
		return false
	}
	if store.Verify(id, captchaCode, true) {
		return true
	} else {
		return false
	}
}
