package service

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/daqnext/daemon"
	"github.com/universe-30/CliAppTemplate/cliCmd"
)

const (
	// name of the service
	name        = "template"
	description = "app template"
)

type Service struct {
	daemon.Daemon
}

func RunServiceCmd() {
	//check command
	subCmds := cliCmd.CmdToDo.CliContext.Command.Names()
	if len(subCmds) == 0 {
		cliCmd.Logger.Fatalln("no sub command")
		return
	}

	kind := daemon.SystemDaemon
	if runtime.GOOS == "darwin" {
		kind = daemon.UserAgent
	}
	srv, err := daemon.New(name, description, kind)
	if err != nil {
		cliCmd.Logger.Fatalln("run daemon error:", err)
	}
	service := &Service{srv}

	action := subCmds[0]
	log.Println(action)

	var status string
	var e error
	switch action {
	case "install":
		status, e = service.Install()
		cliCmd.Logger.Debugln("cmd install")
	case "remove":
		service.Stop()
		status, e = service.Remove()
		cliCmd.Logger.Debugln("cmd remove")
	case "start":
		status, e = service.Start()
		cliCmd.Logger.Debugln("cmd start")
	case "stop":
		status, e = service.Stop()
		cliCmd.Logger.Debugln("cmd stop")
	case "restart":
		service.Stop()
		status, e = service.Start()
		cliCmd.Logger.Debugln("cmd restart")
	case "status":
		status, e = service.Status()
		cliCmd.Logger.Debugln("cmd status")
	default:
		cliCmd.Logger.Debugln("no sub command")
		return
	}

	if e != nil {
		fmt.Println(status, "\nError: ", e)
		os.Exit(1)
	}
	fmt.Println(status)
}
