package lrcache

import (
	"context"
	"math/rand"
	"time"

	"github.com/coreservice-io/CliAppTemplate/basic"
	"github.com/coreservice-io/CliAppTemplate/plugin/redisClient"
	"github.com/coreservice-io/CliAppTemplate/tools/json"
	"github.com/coreservice-io/UCache"
	"github.com/go-redis/redis/v8"
)

const LOCAL_CACHE_TIME = 5 //don't change this number as 5 is the proper number
const TEMP_NULL = "temp_null"

// check weather we need do refresh
// the probobility becomes lager when left seconds close to 0
// this goal of this function is to avoid big traffic glitch
func CheckTtlRefresh(secleft int64) bool {
	if secleft > 0 && secleft <= 3 {
		if rand.Intn(int(secleft)*10) == 1 {
			return true
		}
	}
	return false
}

//first try from localcache if not found then try from remote redis cache
func LRC_Get(ctx context.Context, Redis *redisClient.RedisClient, localCache *UCache.Cache, isJSON bool, keystr string, result interface{}) {

	localvalue, ttl, localexist := localCache.Get(keystr)
	if !CheckTtlRefresh(ttl) && localexist {
		result = localvalue
		return
	}

	//try from remote redis
	r_bytes, err := Redis.Get(ctx, keystr).Bytes()
	if err != nil {
		if err != redis.Nil {
			basic.Logger.Errorln(err)
		}
		result = nil
		return
	} else {
		if isJSON {
			err := json.Unmarshal(r_bytes, result)
			if err == nil {
				localCache.Set(keystr, result, LOCAL_CACHE_TIME)
				return
			} else {
				basic.Logger.Errorln(err)
				result = nil
				return
			}
		} else {
			localCache.Set(keystr, r_bytes, LOCAL_CACHE_TIME)
			result = r_bytes
			return
		}
	}
}

// set both value to both local & remote redis
func LRC_Set(ctx context.Context, Redis *redisClient.RedisClient, localCache *UCache.Cache, isJSON bool, keystr string, value interface{}, redis_ttl_second int64) error {
	if value != nil && isJSON {
		v_json, err := json.Marshal(value)
		if err != nil {
			return err
		}
		localCache.Set(keystr, value, LOCAL_CACHE_TIME)
		Redis.Set(ctx, keystr, v_json, time.Duration(redis_ttl_second)*time.Second)
		return nil
	} else {
		localCache.Set(keystr, value, LOCAL_CACHE_TIME)
		Redis.Set(ctx, keystr, value, time.Duration(redis_ttl_second)*time.Second)
		return nil
	}

}

func LRC_Del(ctx context.Context, Redis *redisClient.RedisClient, localCache *UCache.Cache, keystr string) {
	localCache.Delete(keystr)
	Redis.Del(ctx, keystr)
}
