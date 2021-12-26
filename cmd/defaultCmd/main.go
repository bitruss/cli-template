package defaultCmd

import (
	"time"

	"github.com/fatih/color"
	"github.com/universe-30/UCliAppTemplate/cli"
	"github.com/universe-30/UCliAppTemplate/cli/logger"
)

func StartDefault() {
	//global.Init()
	//defer func() {
	//	logger.LocalLogger.Infoln("StartDefault closed , start to ReleaseResource()")
	//	global.ReleaseResource()
	//}()

	//controllers.DeployApi()

	//print logo

	color.Green(cli.Logo)
	logger.LocalLogger.Infoln("hello world , this default app")
	//somepack.HowToGetGlobalParam()
	///start the server
	//err := global.EchoServer.Start()
	//if err != nil {
	//	logger.LocalLogger.Fatalln(err)
	//}

	for i := 0; i < 10; i++ {
		logger.LocalLogger.Infoln("running")
		time.Sleep(1 * time.Second)
	}

}
