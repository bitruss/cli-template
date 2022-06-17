package local_rate_limiter

import (
	"time"

	"github.com/coreservice-io/cli-template/plugin/reference_plugin"
)

type limitInfo struct {
	CountLeft       int
	LastSetUnixTime int64
}

func CheckRateAllow(key string, durationSecond int64, limitCount int) bool {

	lKey := "rateLimit:" + key
	value, _ := reference_plugin.GetInstance().Get(lKey)

	nowTime := time.Now().UTC().Unix()

	var limit_info *limitInfo
	if value == nil {
		limit_info = &limitInfo{
			CountLeft:       limitCount,
			LastSetUnixTime: nowTime,
		}
	} else {
		limit_info = value.(*limitInfo)
		timePast := nowTime - limit_info.LastSetUnixTime

		//if time past , add count
		if timePast >= durationSecond {
			limit_info.CountLeft = limitCount
			limit_info.LastSetUnixTime = nowTime
		}
	}

	allow := false
	limit_info.CountLeft--
	if limit_info.CountLeft >= 0 {
		allow = true
	} else {
		limit_info.CountLeft = 0
		allow = false
	}

	reference_plugin.GetInstance().Set(lKey, limit_info, durationSecond*5)

	return allow

}
