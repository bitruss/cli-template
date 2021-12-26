package cli

import (
	"log"
	"os"

	"github.com/universe-30/UCliAppTemplate/cli/logger"
)

func ReadArgs() {
	//print any initialzation panic
	//defer func() {
	//	if err := recover(); err != nil {
	//
	//		color.ColorPrintln(color_util.Red, "panic errors:", err.(error).Error())
	//	}
	//}()

	//ini logger
	logger.IniLocalLogger()

	//config app to run
	errRun := configCliCmd().Run(os.Args)
	if errRun != nil {
		log.Fatal(errRun)
		panic(errRun.Error())
	}
}
