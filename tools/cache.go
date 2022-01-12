package tools

import (
	"context"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/universe-30/UCache"
)

// check weather we need do refresh
// the probobility becomes lager when left seconds close to 0
// this goal of this function is to avoid big traffic glitch
func CheckTtlRefresh(secleft int64) bool {
	if secleft > 0 && secleft < 8 {
		if rand.Intn(int(secleft)*50) == 1 {
			return true
		}
	}
	return false
}

//LCR == Local Cache Sync with Remote Redis

// check local cache value is still available and sychronized with remote redis value
func LCR_Check(ctx context.Context, Redis *redis.ClusterClient, localCache *UCache.Cache, keystr string) (interface{}, int64, bool) {
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

// set both value to both local & remote redis
func LCR_Set(ctx context.Context, Redis *redis.ClusterClient, localCache *UCache.Cache, keystr string, value interface{}, ttlSecond int64) {
	localCache.Set(keystr, value, ttlSecond)
	randSyncStr := keystr + ":randsync"
	strsrc := localCache.SetRand(randSyncStr, ttlSecond+10)
	Redis.Set(ctx, randSyncStr, strsrc, time.Duration(ttlSecond+30)*time.Second)
}

//delete cache from both local and remote redis
func LCR_Del(ctx context.Context, Redis *redis.ClusterClient, localCache *UCache.Cache, keystr string) {
	randSyncStr := keystr + ":randsync"
	localCache.Delete(keystr)
	localCache.Delete(randSyncStr)
	Redis.Del(ctx, keystr)
	Redis.Del(ctx, randSyncStr)
}
