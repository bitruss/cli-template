package default_

import (
	"time"

	"github.com/fatih/color"
	"github.com/universe-30/CliAppTemplate/boot"
)

func StartDefault() {

	color.Green(boot.Logo)
	boot.Logger.Infoln("hello world , this default app")

	//just for fun
	for i := 0; i < 10; i++ {
		boot.Logger.Infoln("running")
		time.Sleep(1 * time.Second)
	}
}
