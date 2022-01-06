package config

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/universe-30/CliAppTemplate/basic"
	"github.com/urfave/cli/v2"
)

func ConfigSetting(clictx *cli.Context) {

	configModify := false

	for _, v := range IntConfParams {
		if clictx.IsSet(v) {
			newValue := clictx.Int(v)
			if newValue != 0 {
				basic.Config.Set(v, newValue)
				configModify = true
			}
		}
	}

	for _, v := range StringConfParams {
		if clictx.IsSet(v) {
			newValue := clictx.String(v)
			basic.Config.Set(v, newValue)
			configModify = true
		}
	}

	if configModify {
		err := basic.Config.WriteConfig()
		if err != nil {
			color.Red("config save error:", err)
			return
		}
		fmt.Println("config modified new config")
		fmt.Println(basic.Config.GetConfigAsString())
	}

}
