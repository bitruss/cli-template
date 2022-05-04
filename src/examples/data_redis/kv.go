package data_redis

import (
	"context"

	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/plugin/redis_plugin"
	"github.com/coreservice-io/cli-template/plugin/reference_plugin"
	"github.com/coreservice-io/cli-template/tools/smart_cache"
	"github.com/go-redis/redis/v8"
)

//example for GormDB and tools cache
type PeerInfo struct {
	Tag      string `json:"tag"`
	Location string `json:"location"`
	IP       string `json:"ip"`
}

func SetPeer(peerInfo *PeerInfo, tag string) error {
	key := redis_plugin.GetInstance().GenKey("peerInfo", tag)
	return smart_cache.RR_Set(context.Background(), redis_plugin.GetInstance().ClusterClient, reference_plugin.GetInstance(), true, key, peerInfo, 600)
}

func DeletePeer(tag string) {
	key := redis_plugin.GetInstance().GenKey("peerInfo", tag)
	smart_cache.RR_Del(context.Background(), redis_plugin.GetInstance().ClusterClient, reference_plugin.GetInstance(), key)
}

func GetPeer(tag string) (*PeerInfo, error) {
	key := redis_plugin.GetInstance().GenKey("peerInfo", tag)

	// from local reference ,local cache ttl is very small , fresh value
	// try to get from reference
	r_result := smart_cache.Ref_Get(reference_plugin.GetInstance(), key)
	if r_result != nil {
		basic.Logger.Debugln("GetPeer hit from reference")
		return r_result.(*PeerInfo), nil
	}
	// try to get from redis
	redis_result := &PeerInfo{}
	err := smart_cache.Redis_Get(context.Background(), redis_plugin.GetInstance().ClusterClient, true, key, redis_result)
	if err == nil {
		basic.Logger.Debugln("GetPeer hit from redis")
		smart_cache.Ref_Set(reference_plugin.GetInstance(), key, redis_result)
		return redis_result, nil
	} else if err == redis.Nil || err == smart_cache.TempNil {
		return nil, nil
	} else {
		return nil, err
	}
}
