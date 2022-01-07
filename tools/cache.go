package tools

import (
	"context"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/universe-30/UCache"
)

//to do : smart cache tool
//for redis cluster , when return true , you need to refresh the key-value
//func CheckTtlRefresh(Redis *redis.ClusterClient, ctx context.Context, keystr string) bool {
//	if rand.Intn(10) == 1 {
//		Duration, error := Redis.TTL(ctx, keystr).Result()
//		if error != nil {
//			return true
//		}
//		secleft := Duration.Seconds()
//		if secleft > 0 && secleft < 10 {
//			if rand.Intn(int(secleft)*5) == 1 {
//				return true
//			} else {
//				return false
//			}
//		} else {
//			return false
//		}
//	} else {
//		return false
//	}
//}

func CheckTtlRefresh(secleft int64) bool {
	if secleft > 0 && secleft < 8 {
		if rand.Intn(int(secleft)*50) == 1 {
			return true
		}
	}
	return false
}

func SmartCheck_LocalCache_Redis(ctx context.Context, Redis *redis.ClusterClient, localCache *UCache.Cache, keystr string) (interface{}, int64, bool) {
	localvalue, ttl, localexist := localCache.Get(keystr)
	if !CheckTtlRefresh(ttl) && localexist {
		randSyncStr := keystr + ":randsync"
		rresult, err := Redis.Get(ctx, randSyncStr).Result()
		if err == nil && rresult == localCache.GetRand(randSyncStr) {
			return localvalue, ttl, true
		}
	}
	return nil, 0, false
}

func SmartSet_LocalCache_Redis(ctx context.Context, Redis *redis.ClusterClient, localCache *UCache.Cache, keystr string, value interface{}, ttlSecond int64) {
	localCache.Set(keystr, value, ttlSecond)
	randSyncStr := keystr + ":randsync"
	strsrc := localCache.SetRand(randSyncStr, ttlSecond+10)
	Redis.Set(ctx, randSyncStr, strsrc, time.Duration(ttlSecond+30)*time.Second)
}

func SmartDel_LocalCache_Redis(ctx context.Context, Redis *redis.ClusterClient, localCache *UCache.Cache, keystr string) {
	randSyncStr := keystr + ":randsync"
	localCache.Delete(keystr)
	localCache.Delete(randSyncStr)
	Redis.Del(ctx, keystr)
	Redis.Del(ctx, randSyncStr)
}
