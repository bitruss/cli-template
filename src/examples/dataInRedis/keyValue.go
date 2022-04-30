package dataInRedis

import (
	"context"

	"github.com/coreservice-io/CliAppTemplate/basic"
	"github.com/coreservice-io/CliAppTemplate/plugin/redis_plugin"
	"github.com/coreservice-io/CliAppTemplate/plugin/reference_plugin"
	"github.com/coreservice-io/CliAppTemplate/tools/smartCache"
	"github.com/go-redis/redis/v8"
)

//example for GormDB and tools cache
type PeerInfo struct {
	Tag      string
	Location string
	IP       string
}

func SetPeer(peerInfo *PeerInfo, tag string) error {
	//peerInfo in param data
	//&PeerInfo{
	//	Tag: "abcd",
	//	Location:"USA",
	//	IP:"127.0.0.1",
	//}

	key := redis_plugin.GetInstance().GenKey("peerInfo", tag)
	return smartCache.RR_Set(context.Background(), redis_plugin.GetInstance().ClusterClient, reference_plugin.GetInstance(), true, key, peerInfo, 600)
}

func DeletePeer(tag string) {
	key := redis_plugin.GetInstance().GenKey("peerInfo", tag)
	smartCache.RR_Del(context.Background(), redis_plugin.GetInstance().ClusterClient, reference_plugin.GetInstance(), key)
}

func GetPeer(tag string, forceUpdate bool) (*PeerInfo, error) {
	key := redis_plugin.GetInstance().GenKey("peerInfo", tag)
	if !forceUpdate {
		// try to get from reference
		result := smartCache.Ref_Get(reference_plugin.GetInstance(), key)
		if result != nil {
			basic.Logger.Debugln("GetPeer hit from reference")
			return result.(*PeerInfo), nil
		}
	}

	// try to get from redis
	redis_result := &PeerInfo{}
	err := smartCache.Redis_Get(context.Background(), redis_plugin.GetInstance().ClusterClient, true, key, redis_result)
	if err == nil {
		basic.Logger.Debugln("GetPeer hit from redis")
		smartCache.Ref_Set(reference_plugin.GetInstance(), key, redis_result)
		return redis_result, nil
	} else if err == redis.Nil || err == smartCache.TempNil {
		return nil, nil
	} else {
		return nil, err
	}
}
