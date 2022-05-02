package smart_cache

import (
	"context"
	"math/rand"
	"reflect"
	"time"

	"github.com/coreservice-io/UReference"
	"github.com/coreservice-io/cli-template/tools/json"
	"github.com/go-redis/redis/v8"
)

type temp_nil_error string

func (e temp_nil_error) Error() string { return string(e) }

const TempNil = temp_nil_error("temp_nil")
const temp_nil = "temp_nil"
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
	return Ref_Set_RTTL(localRef, keystr, value, local_reference_secs)
}

func Ref_Set_RTTL(localRef *UReference.Reference, keystr string, value interface{}, ref_ttl_second int64) error {
	return localRef.Set(keystr, value, ref_ttl_second)
}

// //first try from localRef if not found then try from remote redis
func Redis_Get(ctx context.Context, Redis *redis.ClusterClient, isJSON bool, keystr string, result interface{}) error {

	scmd := Redis.Get(ctx, keystr) //trigger remote redis get
	r_bytes, err := scmd.Bytes()
	if err != nil {
		return err
	}

	if string(r_bytes) == temp_nil {
		return TempNil
	}

	if isJSON {
		return json.Unmarshal(r_bytes, result)
	} else {
		return scmd.Scan(result)
	}
}

func RR_Set(ctx context.Context, Redis *redis.ClusterClient, localRef *UReference.Reference, isJSON bool, keystr string, value interface{}, redis_ttl_second int64) error {
	return RR_Set_RTTL(ctx, Redis, localRef, isJSON, keystr, value, redis_ttl_second, local_reference_secs)
}

// reference set && redis set
// set both value to both local reference & remote redis
func RR_Set_RTTL(ctx context.Context, Redis *redis.ClusterClient, localRef *UReference.Reference, isJSON bool, keystr string, value interface{}, redis_ttl_second int64, ref_ttl_second int64) error {
	if value == nil {
		//rare case normally should not happen
		//set 10 seconds to fast refresh
		return Redis.Set(ctx, keystr, temp_nil, time.Duration(10)*time.Second).Err()
	}
	if isJSON {
		err := localRef.Set(keystr, value, ref_ttl_second)
		if err != nil {
			return err
		}
		v_json, err := json.Marshal(value)
		if err != nil {
			return err
		}
		return Redis.Set(ctx, keystr, v_json, time.Duration(redis_ttl_second)*time.Second).Err()
	} else {
		err := localRef.Set(keystr, value, ref_ttl_second)
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

func RR_Del(ctx context.Context, Redis *redis.ClusterClient, localRef *UReference.Reference, keystr string) {
	localRef.Delete(keystr)
	Redis.Del(ctx, keystr)
}
