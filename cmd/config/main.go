package config

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/universe-30/CliAppTemplate/boot"
)

func ConfigSetting() {
	//fmt.Println(cliCmd.Config.GetConfigAsString())
	c := boot.CmdToDo.CliContext
	configModify := false

	//example
	intConfParams := []string{"http_port", "db_host"}
	stringConfParams := []string{"db_host", "db_name"}

	for _, v := range intConfParams {
		if c.IsSet(v) {
			newValue := c.Int(v)
			if newValue != 0 {
				boot.Config.Set(v, newValue)
				configModify = true
			}
		}
	}

	for _, v := range stringConfParams {
		if c.IsSet(v) {
			newValue := c.String(v)
			boot.Config.Set(v, newValue)
			configModify = true
		}
	}

	if configModify {
		err := boot.Config.WriteConfig()
		if err != nil {
			color.Red("config save error:", err)
			return
		}
		fmt.Println("config modified new config")
		fmt.Println(boot.Config.GetConfigAsString())
	}

}
