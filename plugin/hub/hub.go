package hub

import (
	"fmt"

	uhub "github.com/coreservice-io/UHub"
)

var instanceMap = map[string]*uhub.Hub{}

func GetInstance() *uhub.Hub {
	return instanceMap["default"]
}

func GetInstance_(name string) *uhub.Hub {
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
	instanceMap[name] = &uhub.Hub{}
	return nil
}
