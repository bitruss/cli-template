package limiter

import (
	"time"

	"github.com/coreservice-io/cli-template/plugin/reference_plugin"
)

type limitInfo struct {
	CountLeft       int
	LastSetUnixTime int64
}

// check key allowed to pass
func Allow(key string, duration_second int64, Count int) bool {

	if duration_second <= 0 {
		return true
	}

	lKey := "rateLimit:" + key
	value, _ := reference_plugin.GetInstance().Get(lKey)

	var limit_info *limitInfo
	nowTime := time.Now().UTC().Unix()

	allow := false

	if value == nil {

		allow = true

		limit_info = &limitInfo{
			CountLeft:       Count,
			LastSetUnixTime: nowTime,
		}

		reference_plugin.GetInstance().Set(lKey, limit_info, duration_second*5+60)

	} else {
		limit_info = value.(*limitInfo)
		//if time past , add count
		if nowTime-limit_info.LastSetUnixTime >= duration_second {
			limit_info.CountLeft = Count
			limit_info.LastSetUnixTime = nowTime
			allow = true
		} else {
			limit_info.CountLeft--
			if limit_info.CountLeft >= 0 {
				allow = true
			} else {
				allow = false
			}
		}
	}

	return allow
}
