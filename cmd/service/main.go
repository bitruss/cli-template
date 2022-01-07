package service

import (
	"fmt"
	"os"

	"github.com/universe-30/CliAppTemplate/basic"
	"github.com/universe-30/CliAppTemplate/components"
	"github.com/urfave/cli/v2"
)

func RunServiceCmd(clictx *cli.Context) {
	//check command
	subCmds := clictx.Command.Names()
	if len(subCmds) == 0 {
		basic.Logger.Fatalln("no sub command")
		return
	}

	action := subCmds[0]
	compDeamon := components.NewDaemonService()

	var status string
	var e error
	switch action {
	case "install":
		status, e = compDeamon.Install()
		basic.Logger.Debugln("cmd install")
	case "remove":
		compDeamon.Stop()
		status, e = compDeamon.Remove()
		basic.Logger.Debugln("cmd remove")
	case "start":
		status, e = compDeamon.Start()
		basic.Logger.Debugln("cmd start")
	case "stop":
		status, e = compDeamon.Stop()
		basic.Logger.Debugln("cmd stop")
	case "restart":
		compDeamon.Stop()
		status, e = compDeamon.Start()
		basic.Logger.Debugln("cmd restart")
	case "status":
		status, e = compDeamon.Status()
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
