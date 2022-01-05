package service

import (
	"fmt"
	"log"
	"os"

	"github.com/universe-30/CliAppTemplate/boot"
)

func RunServiceCmd() {
	//check command
	subCmds := boot.CmdToDo.CliContext.Command.Names()
	if len(subCmds) == 0 {
		boot.Logger.Fatalln("no sub command")
		return
	}

	action := subCmds[0]
	log.Println(action)

	var status string
	var e error
	switch action {
	case "install":
		status, e = CompDeamon.Install()
		boot.Logger.Debugln("cmd install")
	case "remove":
		CompDeamon.Stop()
		status, e = CompDeamon.Remove()
		boot.Logger.Debugln("cmd remove")
	case "start":
		status, e = CompDeamon.Start()
		boot.Logger.Debugln("cmd start")
	case "stop":
		status, e = CompDeamon.Stop()
		boot.Logger.Debugln("cmd stop")
	case "restart":
		CompDeamon.Stop()
		status, e = CompDeamon.Start()
		boot.Logger.Debugln("cmd restart")
	case "status":
		status, e = CompDeamon.Status()
		boot.Logger.Debugln("cmd status")
	default:
		boot.Logger.Debugln("no sub command")
		return
	}

	if e != nil {
		fmt.Println(status, "\nError: ", e)
		os.Exit(1)
	}
	fmt.Println(status)
}
