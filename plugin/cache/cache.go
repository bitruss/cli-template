package cache

import (
	"github.com/universe-30/UCache"
)

//var cache *UCache.Cache
//var once sync.Once
//
//func Init(){
//	//only run once
//	once.Do(func() {
//		cache = UCache.New()
//	})
//}
//
//func GetSingleInstance() *UCache.Cache {
//	Init()
//	return cache
//}

var cache *UCache.Cache

func GetSingleInstance() *UCache.Cache {
	return cache
}

func Init() {
	cache = UCache.New()
}
