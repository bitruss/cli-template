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

	if c.IsSet("http_port") {
		newValue := c.Int("http_port")
		if newValue != 0 {
			boot.Config.Set("http_port", newValue)
			configModify = true
		}
	}

	if c.IsSet("db_host") {
		newValue := c.String("db_host")
		boot.Config.Set("db_host", newValue)
		configModify = true
	}

	if c.IsSet("db_port") {
		newValue := c.Int("db_port")
		if newValue != 0 {
			boot.Config.Set("db_port", newValue)
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
