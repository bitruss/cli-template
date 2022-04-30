package hub_plugin

import (
	"fmt"

	"github.com/coreservice-io/UHub"
)

var instanceMap = map[string]*UHub.Hub{}

func GetInstance() *UHub.Hub {
	return instanceMap["default"]
}

func GetInstance_(name string) *UHub.Hub {
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
		return fmt.Errorf("hub instance <%s> has already initialized", name)
	}
	instanceMap[name] = &UHub.Hub{}
	return nil
}
