package lrcache

import (
	"context"
	"math/rand"
	"reflect"
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
		reflect.ValueOf(result).Set(reflect.ValueOf(localvalue).)
		//*result = localvalue
		return
	}

	if isJSON {
		r_bytes, err := Redis.Get(ctx, keystr).Bytes()
		if err != nil && err == redis.Nil {
			result = nil
			return
		}

		if string(r_bytes) == TEMP_NULL {
			result = nil
			return
		}

		err = json.Unmarshal(r_bytes, result)
		if err == nil {
			localCache.Set(keystr, result, LOCAL_CACHE_TIME)
			return
		} else {
			basic.Logger.Errorln(err)
			result = nil
			return
		}
	} else {
		rCmd := Redis.Get(ctx, keystr)
		switch result.(type) {
		case *uint8:
			r, err := rCmd.Uint64()
			if err != nil {
				basic.Logger.Errorln(err)
				return
			}
			*result.(*uint8) = uint8(r)
		case *uint16:
			r, err := rCmd.Uint64()
			if err != nil {
				basic.Logger.Errorln(err)
				return
			}
			*result.(*uint16) = uint16(r)
		case *uint32:
			r, err := rCmd.Uint64()
			if err != nil {
				basic.Logger.Errorln(err)
				return
			}
			*result.(*uint32) = uint32(r)
		case *uint64:
			r, err := rCmd.Uint64()
			if err != nil {
				basic.Logger.Errorln(err)
				return
			}
			*result.(*uint64) = r
		case *uint:
			r, err := rCmd.Uint64()
			if err != nil {
				basic.Logger.Errorln(err)
				return
			}
			*result.(*uint) = uint(r)
		case *int8:
			r, err := rCmd.Int64()
			if err != nil {
				basic.Logger.Errorln(err)
				return
			}
			*result.(*int8) = int8(r)
		case *int16:
			r, err := rCmd.Int64()
			if err != nil {
				basic.Logger.Errorln(err)
				return
			}
			*result.(*int16) = int16(r)
		case *int32:
			r, err := rCmd.Int64()
			if err != nil {
				basic.Logger.Errorln(err)
				return
			}
			*result.(*int32) = int32(r)
		case *int64:
			r, err := rCmd.Int64()
			if err != nil {
				basic.Logger.Errorln(err)
				return
			}
			*result.(*int64) = r
		case *int:
			r, err := rCmd.Int()
			if err != nil {
				basic.Logger.Errorln(err)
				return
			}
			*result.(*int) = r
		case *float32:
			r, err := rCmd.Float32()
			if err != nil {
				basic.Logger.Errorln(err)
				return
			}
			*result.(*float32) = r
		case *float64:
			r, err := rCmd.Float64()
			if err != nil {
				basic.Logger.Errorln(err)
				return
			}
			*result.(*float64) = r
		case complex64, *complex64:

		case complex128, *complex128:

		case *bool:
			r, err := rCmd.Bool()
			if err != nil {
				basic.Logger.Errorln(err)
				return
			}
			*result.(*bool) = r
		case *string:
			r := rCmd.String()
			*result.(*string) = r
		default:
			basic.Logger.Errorln("value type error")
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
