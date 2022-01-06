package default_

import (
	"time"

	"github.com/fatih/color"
	"github.com/universe-30/CliAppTemplate/basic"
	"github.com/urfave/cli/v2"
)

func StartDefault(clictx *cli.Context) {

	color.Green(basic.Logo)
	basic.Logger.Infoln("hello world , this default app")

	//just for fun
	for i := 0; i < 10; i++ {
		basic.Logger.Infoln("running")
		time.Sleep(1 * time.Second)
	}
}
