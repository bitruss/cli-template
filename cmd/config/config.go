package config

import (
	"errors"
	"strings"

	"github.com/coreservice-io/cli-template/basic/conf"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func Cli_get_flags() []cli.Flag {
	allflags := []cli.Flag{}
	allflags = append(allflags, &cli.StringFlag{Name: "log.level", Required: false})
	allflags = append(allflags, &cli.StringFlag{Name: "http.enable", Required: false})
	allflags = append(allflags, &cli.StringFlag{Name: "https.enable", Required: false})
	return allflags
}

var log_level_map = map[string]struct{}{
	"TRAC":  {},
	"TRACE": {},
	"DEBU":  {},
	"DEBUG": {},
	"INFO":  {},
	"WARN":  {},
	"ERRO":  {},
	"ERROR": {},
	"FATA":  {},
	"FATAL": {},
	"PANI":  {},
	"PANIC": {},
}

func Cli_set_config(clictx *cli.Context) error {
	config := conf.Get_config()

	if clictx.IsSet("log.level") {
		log_level := strings.ToUpper(clictx.String("log.level"))
		_, exist := log_level_map[log_level]
		if !exist {
			return errors.New("log level error")
		}

		config.Custom_config_tree.Set("log.level", log_level)
	}

	if clictx.IsSet("http.enable") {
		http_enable := clictx.Bool("http.enable")
		config.Custom_config_tree.Set("http.enable", http_enable)
	}

	if clictx.IsSet("https.enable") {
		https_enable := clictx.Bool("https.enable")
		config.Custom_config_tree.Set("https.enable", https_enable)
	}

	err := config.Save_custom_config()
	if err != nil {
		color.Red("save custom config error:", err)
		return err
	} else {
		color.Green("new config set success")
	}
	return nil
}
