package config

import (
	"fmt"

	"github.com/coreservice-io/cli-template/configuration"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func ConfigSetting(clictx *cli.Context) {

	for k, v := range allKeysMap {
		if clictx.IsSet(k) {
			switch v {
			case "string":
				newValue := clictx.String(k)
				configuration.Config.Viper.Set(k, newValue)
			case "bool":
				newValue := clictx.Bool(k)
				configuration.Config.Viper.Set(k, newValue)
			case "float64":
				newValue := clictx.Float64(k)
				configuration.Config.Viper.Set(k, newValue)
			case "int":
				newValue := clictx.Int(k)
				configuration.Config.Viper.Set(k, newValue)
			default:
				newValue := clictx.Value(k)
				configuration.Config.Viper.Set(k, newValue)
			}
		}
	}

	err := configuration.Config.WriteConfig()
	if err != nil {
		color.Red("config save error:", err)
		return
	}
	fmt.Println("config modified")
}
