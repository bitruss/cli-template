package cache

import (
	"github.com/universe-30/UCache"
)

var cache *UCache.Cache

func GetSingleInstance() *UCache.Cache {
	return cache
}

func Init() {
	if cache != nil {
		return
	}
	cache = UCache.New()
}
