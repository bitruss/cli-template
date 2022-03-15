package example_run

import (
	"github.com/coreservice-io/CliAppTemplate/basic"
	"github.com/coreservice-io/CliAppTemplate/configuration"
)

//example get complex config
func ComplexConfig_run() {
	provide_folder, err := configuration.Config.GetProvideFolders()
	if err != nil {
		basic.Logger.Errorln(err)
	}
	for _, v := range provide_folder {
		basic.Logger.Debugln("path:", v.AbsPath, "size:", v.SizeGB)
	}
}
