package daemon

import (
	"fmt"
	"runtime"

	"github.com/daqnext/daemon"
)

const (
	// name of the service
	name        = "template"
	description = "app template"
)

type Service struct {
	daemon.Daemon
}

var instanceMap = map[string]*Service{}

func GetDefaultInstance() *Service {
	return instanceMap["default"]
}

func GetInstance(name string) *Service {
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
		return fmt.Errorf("daemon instance <%s> has already initialized", name)
	}

	kind := daemon.SystemDaemon
	if runtime.GOOS == "darwin" {
		kind = daemon.UserAgent
	}
	srv, err := daemon.New(name, description, kind)
	if err != nil {
		return err
	}
	instanceMap[name] = &Service{srv}
	return nil
}
