package service

import (
	"fmt"
	"log"
	"os"

	"github.com/universe-30/CliAppTemplate/basic"
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
	log.Println(action)

	var status string
	var e error
	switch action {
	case "install":
		status, e = CompDeamon.Install()
		basic.Logger.Debugln("cmd install")
	case "remove":
		CompDeamon.Stop()
		status, e = CompDeamon.Remove()
		basic.Logger.Debugln("cmd remove")
	case "start":
		status, e = CompDeamon.Start()
		basic.Logger.Debugln("cmd start")
	case "stop":
		status, e = CompDeamon.Stop()
		basic.Logger.Debugln("cmd stop")
	case "restart":
		CompDeamon.Stop()
		status, e = CompDeamon.Start()
		basic.Logger.Debugln("cmd restart")
	case "status":
		status, e = CompDeamon.Status()
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
