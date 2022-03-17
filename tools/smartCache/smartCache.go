package smartCache

import (
	"context"
	"math/rand"
	"reflect"
	"time"

	"github.com/coreservice-io/CliAppTemplate/plugin/redisClient"
	"github.com/coreservice-io/CliAppTemplate/tools/json"
	"github.com/coreservice-io/UReference"
)

const temp_null = "temp_null"
const local_reference_secs = 5 //don't change this number as 5 is the proper number

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

func Ref_Get(localRef *UReference.Reference, keystr string) (result interface{}) {
	localvalue, ttl, localexist := localRef.Get(keystr)
	if !CheckTtlRefresh(ttl) && localexist {
		return localvalue
	}
	return nil
}

func Ref_Set(localRef *UReference.Reference, keystr string, value interface{}) error {
	return localRef.Set(keystr, value, local_reference_secs)
}

//first try from localRef if not found then try from remote redis
func Redis_Get(ctx context.Context, Redis *redisClient.RedisClient, isJSON bool, keystr string, result interface{}) error {

	scmd := Redis.Get(ctx, keystr) //trigger remote redis get
	r_bytes, err := scmd.Bytes()
	if err != nil {
		return err
	}

	if string(r_bytes) == temp_null {
		return nil
	}

	if isJSON {
		return json.Unmarshal(r_bytes, result)
	} else {
		return scmd.Scan(result)
	}
}

// reference set && redis set
// set both value to both local reference & remote redis
func RR_Set(ctx context.Context, Redis *redisClient.RedisClient, localRef *UReference.Reference, isJSON bool, keystr string, value interface{}, redis_ttl_second int64) error {
	if value == nil {
		return Redis.Set(ctx, keystr, temp_null, time.Duration(redis_ttl_second)*time.Second).Err()
	}
	if isJSON {
		err := localRef.Set(keystr, value, local_reference_secs)
		if err != nil {
			return err
		}
		v_json, err := json.Marshal(value)
		if err != nil {
			return err
		}
		return Redis.Set(ctx, keystr, v_json, time.Duration(redis_ttl_second)*time.Second).Err()
	} else {
		err := localRef.Set(keystr, value, local_reference_secs)
		if err != nil {
			return err
		}
		tp := reflect.TypeOf(value).Kind()
		if tp == reflect.Ptr {
			return Redis.Set(ctx, keystr, reflect.ValueOf(value).Elem().Interface(), time.Duration(redis_ttl_second)*time.Second).Err()
		} else {
			return Redis.Set(ctx, keystr, value, time.Duration(redis_ttl_second)*time.Second).Err()
		}
	}
}

func RR_Del(ctx context.Context, Redis *redisClient.RedisClient, localRef *UReference.Reference, keystr string) {
	localRef.Delete(keystr)
	Redis.Del(ctx, keystr)
}
