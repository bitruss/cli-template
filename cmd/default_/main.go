package default_

import (
	"github.com/coreservice-io/CliAppTemplate/basic"
	"github.com/coreservice-io/CliAppTemplate/cmd/default_/plugin"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func StartDefault(clictx *cli.Context) {

	//defer func() {
	//	//global.ReleaseResources()
	//}()
	color.Green(basic.Logo)
	//ini components and run example
	plugin.InitPlugin()

	start_service()
}
