package config

import (
	"github.com/coreservice-io/cli-template/basic/conf"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func Cli_get_flags() []cli.Flag {

	allflags := []cli.Flag{}
	allflags = append(allflags, &cli.StringFlag{Name: "log_level", Required: false})
	allflags = append(allflags, &cli.StringFlag{Name: "http.enable", Required: false})
	allflags = append(allflags, &cli.StringFlag{Name: "https.enable", Required: false})
	return allflags
}

func Cli_set_config(clictx *cli.Context) {
	config := conf.Get_config()

	if clictx.IsSet("log_level") {
		config.Toml_config.Log_level = clictx.String("log_level")
	}

	if clictx.IsSet("http.enable") {
		config.Toml_config.Http.Enable = clictx.Bool("http.enable")
	}

	if clictx.IsSet("https.enable") {
		config.Toml_config.Https.Enable = clictx.Bool("https.enable")
	}

	err := config.Save_config()
	if err != nil {
		color.Red("save config error:", err)
	}
}
