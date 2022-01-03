package cliCmd

import (
	"os"
)

func ReadArgs() {
	//print any initialzation panic
	//defer func() {
	//	if err := recover(); err != nil {
	//
	//		color.ColorPrintln(color_util.Red, "panic errors:", err.(error).Error())
	//	}
	//}()

	//config app to run
	errRun := configCliCmd().Run(os.Args)
	if errRun != nil {
		Logger.Panic(errRun)
	}
}
