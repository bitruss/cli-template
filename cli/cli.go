package cli

import (
	"log"
	"os"
)

func Init() {
	InitLogger()
}

func ReadArgs() {
	//print any initialzation panic
	//defer func() {
	//	if err := recover(); err != nil {
	//
	//		color.ColorPrintln(color_util.Red, "panic errors:", err.(error).Error())
	//	}
	//}()

	//ini logger
	InitLogger()

	//config app to run
	errRun := configCliCmd().Run(os.Args)
	if errRun != nil {
		log.Fatal(errRun)
		panic(errRun.Error())
	}
}
