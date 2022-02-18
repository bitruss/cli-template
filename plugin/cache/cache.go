package cache

import (
	"fmt"

	"github.com/universe-30/UCache"
)

var instanceMap = map[string]*UCache.Cache{}

func GetDefaultInstance() *UCache.Cache {
	return instanceMap["default"]
}

func GetInstance(name string) *UCache.Cache {
	return instanceMap[name]
}

// Init a new instance.
//  If only need one instance, use empty name "". Use GetDefaultInstance() to get.
//  If you need several instance, run Init() with different <name>. Use GetInstance(<name>) to get.
func Init(name string) error {
	if name == "" {
		name = "default"
	}

	_, exist := instanceMap[name]
	if exist {
		return fmt.Errorf("cache instance <%s> has already initialized", name)
	}
	instanceMap[name] = UCache.New()
	return nil
}
