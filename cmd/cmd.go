package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/basic/conf"
	"github.com/coreservice-io/cli-template/cmd/config"
	"github.com/coreservice-io/cli-template/cmd/default_"
	"github.com/coreservice-io/cli-template/cmd/default_/http/api"
	"github.com/coreservice-io/cli-template/cmd/log"
	ilog "github.com/coreservice-io/log"
	"github.com/urfave/cli/v2"
)

const CMD_NAME_DEFAULT = "default"
const CMD_NAME_GEN_API = "gen_api"
const CMD_NAME_LOG = "log"
const CMD_NAME_CONFIG = "config"

////////config to do cmd ///////////
func ConfigCmd() *cli.App {

	//////////init config/////////////
	toml_conf_path := "configs/default.toml"

	real_args := []string{}
	for _, arg := range os.Args {
		arg_lower := strings.ToLower(arg)
		if strings.HasPrefix(arg_lower, "-conf=") || strings.HasPrefix(arg_lower, "--conf=") {

			toml_target := strings.Trim(arg_lower, "-conf=")
			toml_target = strings.Trim(toml_target, "--conf=")
			toml_conf_path = "configs/" + toml_target + ".toml"
			fmt.Println("toml_conf_path", toml_conf_path)
			continue
		}
		real_args = append(real_args, arg)
	}

	os.Args = real_args

	conf_err := conf.Init_config(toml_conf_path)
	if conf_err != nil {
		basic.Logger.Fatalln("config err", conf_err)
	}

	configuration := conf.Get_config()

	/////set loglevel//////
	basic.Logger.SetLevel(ilog.ParseLogLevel(configuration.Toml_config.Log_level))
	////////////////////////////////

	return &cli.App{
		Action: func(clictx *cli.Context) error {
			default_.StartDefault(clictx)
			return nil
		},

		Commands: []*cli.Command{
			{
				Name:  CMD_NAME_LOG,
				Usage: "print all logs",
				Flags: log.GetFlags(),
				Action: func(clictx *cli.Context) error {
					log.StartLog(clictx)
					return nil
				},
			},
			{
				Name:  CMD_NAME_GEN_API,
				Usage: "api command",
				Action: func(clictx *cli.Context) error {
					api.Gen_Api_Docs()
					return nil
				},
			},
			{
				Name:  CMD_NAME_CONFIG,
				Usage: "config command",
				Subcommands: []*cli.Command{
					//show config
					{
						Name:  "show",
						Usage: "show configs",
						Action: func(clictx *cli.Context) error {
							fmt.Println("======== start of config ========")
							configs, _ := conf.Get_config().Read_config_file()
							fmt.Println(configs)
							fmt.Println("======== end  of  config ========")
							return nil
						},
					},
					//set config
					{
						Name:  "set",
						Usage: "set config",
						Flags: append(config.Cli_get_flags(), &cli.StringFlag{Name: "config", Required: false}),
						Action: func(clictx *cli.Context) error {
							config.Cli_set_config(clictx)
							return nil
						},
					},
				},
			},
		},
	}
}
