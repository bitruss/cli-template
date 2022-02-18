package service

import (
	"fmt"
	"os"

	"github.com/universe-30/CliAppTemplate/basic"
	"github.com/universe-30/CliAppTemplate/plugin/daemon"
	"github.com/urfave/cli/v2"
)

func RunServiceCmd(clictx *cli.Context) {
	//check command
	subCmds := clictx.Command.Names()
	if len(subCmds) == 0 {
		basic.Logger.Fatalln("no sub command")
	}

	action := subCmds[0]
	err := daemon.Init("")
	if err != nil {
		basic.Logger.Fatalln("init daemon service error:", err)
	}

	var status string
	var e error
	switch action {
	case "install":
		status, e = daemon.GetDefaultInstance().Install()
		basic.Logger.Debugln("cmd install")
	case "remove":
		daemon.GetDefaultInstance().Stop()
		status, e = daemon.GetDefaultInstance().Remove()
		basic.Logger.Debugln("cmd remove")
	case "start":
		status, e = daemon.GetDefaultInstance().Start()
		basic.Logger.Debugln("cmd start")
	case "stop":
		status, e = daemon.GetDefaultInstance().Stop()
		basic.Logger.Debugln("cmd stop")
	case "restart":
		daemon.GetDefaultInstance().Stop()
		status, e = daemon.GetDefaultInstance().Start()
		basic.Logger.Debugln("cmd restart")
	case "status":
		status, e = daemon.GetDefaultInstance().Status()
		basic.Logger.Debugln("cmd status")
	default:
		basic.Logger.Debugln("no sub command")
		return
	}

	if e != nil {
		fmt.Println(status, "\nError: ", e)
		os.Exit(1)
	}
	fmt.Println(status)
}
