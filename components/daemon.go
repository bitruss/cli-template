package components

import (
	"runtime"

	"github.com/daqnext/daemon"
	"github.com/universe-30/CliAppTemplate/basic"
)

const (
	// name of the service
	name        = "template"
	description = "app template"
)

type Service struct {
	daemon.Daemon
}

func NewDaemonService() *Service {
	kind := daemon.SystemDaemon
	if runtime.GOOS == "darwin" {
		kind = daemon.UserAgent
	}
	srv, err := daemon.New(name, description, kind)
	if err != nil {
		basic.Logger.Fatalln("run daemon error:", err)
	}
	return &Service{srv}
}
