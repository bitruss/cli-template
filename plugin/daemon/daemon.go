package daemon

import (
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

var service *Service

func GetSingleInstance() *Service {
	return service
}

func Init() error {
	kind := daemon.SystemDaemon
	if runtime.GOOS == "darwin" {
		kind = daemon.UserAgent
	}
	srv, err := daemon.New(name, description, kind)
	if err != nil {
		return err
	}
	service = &Service{srv}
	return nil
}
