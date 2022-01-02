package components

import (
	"github.com/universe-30/UCache"
)

func InitCacheMgr() *UCache.Cache {
	return UCache.New()
}
