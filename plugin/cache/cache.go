package cache

import (
	"fmt"

	"github.com/coreservice-io/UCache"
)

var instanceMap = map[string]*UCache.Cache{}

func GetInstance() *UCache.Cache {
	return instanceMap["default"]
}

func GetInstance_(name string) *UCache.Cache {
	return instanceMap[name]
}

func Init() error {
	return Init_("default")
}

// Init a new instance.
//  If only need one instance, use empty name "". Use GetDefaultInstance() to get.
//  If you need several instance, run Init() with different <name>. Use GetInstance(<name>) to get.
func Init_(name string) error {
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
