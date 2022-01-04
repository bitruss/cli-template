package defaultCmd

import (
	"time"

	"github.com/fatih/color"
	"github.com/universe-30/CliAppTemplate/cliCmd"
)

func StartDefault() {
	//defer func() {
	//	logger.LocalLogger.Infoln("StartDefault closed , start to ReleaseResource()")
	//	global.ReleaseResource()
	//}()

	//controllers.DeployApi()

	//print logo

	color.Green(cliCmd.Logo)
	cliCmd.Logger.Infoln("hello world , this default app")

	///start the server
	//err := global.EchoServer.Start()
	//if err != nil {
	//	logger.LocalLogger.Fatalln(err)
	//}

	for i := 0; i < 10; i++ {
		cliCmd.Logger.Infoln("running")
		time.Sleep(1 * time.Second)
	}

}
