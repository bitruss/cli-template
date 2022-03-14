package lrcache

import (
	"github.com/coreservice-io/UCache"
)

func Local_Get(localCache *UCache.Cache, keystr string) (result interface{}) {
	localvalue, ttl, localexist := localCache.Get(keystr)
	if !CheckTtlRefresh(ttl) && localexist {
		return localvalue
	}
	//need get from origin
	return nil
}
