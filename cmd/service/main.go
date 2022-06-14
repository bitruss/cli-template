package service

import (
	"os"
	"path"
	"path/filepath"

	"github.com/coreservice-io/cli-template/basic"
	"github.com/kardianos/service"
	"github.com/urfave/cli/v2"
)

func RunServiceCmd(clictx *cli.Context, daemon_name string, action string, s service.Service) {

	if daemon_name == "" {
		basic.Logger.Fatalln("daemon_name in config should not be vacant")
		return
	}

	exe_path, exe_path_err := os.Executable()
	if exe_path_err != nil {
		basic.Logger.Errorln(exe_path_err)
		return

	}

	exeDir := filepath.Dir(exe_path)

	if _, dir_err := os.Stat(path.Join(exeDir, "assets")); dir_err != nil {
		basic.Logger.Errorln("error -> please check:")
		basic.Logger.Errorln("1.dont directly `go run` for service, always `go build` first")
		basic.Logger.Errorln("2.the assets folder exist parellel to the excutable file ")
		return
	}

	basic.Logger.Infoln("exefile:" + exe_path + " to be service target")

	switch action {
	case "install":
		err := s.Install()
		if err != nil {
			basic.Logger.Errorln("install service error:", err)
		} else {
			basic.Logger.Infoln("service installed")
		}
	case "remove":
		err := s.Uninstall()
		if err != nil {
			basic.Logger.Errorln("remove service error:", err)
		} else {
			basic.Logger.Infoln("service removed")
		}
	case "start":
		err := s.Start()
		if err != nil {
			basic.Logger.Errorln("start service error:", err)
		} else {
			basic.Logger.Infoln("service started")
		}
	case "stop":
		err := s.Stop()
		if err != nil {
			basic.Logger.Errorln("stop service error:", err)
		} else {
			basic.Logger.Infoln("service stopped")
		}
	case "restart":
		err := s.Restart()
		if err != nil {
			basic.Logger.Errorln("restart service error:", err)
		} else {
			basic.Logger.Infoln("service restarted")
		}
	case "status":
		status, err := s.Status()
		if err != nil {
			basic.Logger.Errorln(err)
		}
		switch status {
		case service.StatusRunning:
			basic.Logger.Infoln("service status:", "RUNNING")
		case service.StatusStopped:
			basic.Logger.Infoln("service status:", "STOPPED")
		default:
			basic.Logger.Infoln("service status:", "UNKNOWN")
		}
	default:
		basic.Logger.Debugln("no sub command")
		return
	}
}
