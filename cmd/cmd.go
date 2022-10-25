package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/coreservice-io/cli-template/config"
	"github.com/coreservice-io/cli-template/db_cmd"
	"github.com/coreservice-io/cli-template/default_cmd"
	"github.com/coreservice-io/cli-template/default_cmd/http/api"
	"github.com/coreservice-io/cli-template/log_cmd"
)

const CMD_NAME_DEFAULT = "default"
const CMD_NAME_GEN_API = "gen_api"
const CMD_NAME_LOG = "log"
const CMD_NAME_DB = "db"
const CMD_NAME_CONFIG = "config"

// //////config to do cmd ///////////
func ConfigCmd() *cli.App {

	real_args := config.ConfigBasic()

	var defaultAction = func(clictx *cli.Context) error {
		default_cmd.StartDefault(clictx)
		return nil
	}

	if len(real_args) > 1 {
		defaultAction = nil
	}

	return &cli.App{
		Action: defaultAction, //only run if no sub command

		//run if sub command not correct
		CommandNotFound: func(context *cli.Context, s string) {
			fmt.Println("command not find, use -h or --help show help")
		},

		Commands: []*cli.Command{
			{
				Name:  CMD_NAME_GEN_API,
				Usage: "api command",
				Action: func(clictx *cli.Context) error {
					api.GenApiDocs()
					return nil
				},
			},
			{
				Name:  CMD_NAME_LOG,
				Usage: "print all logs",
				Flags: log_cmd.GetFlags(),
				Action: func(clictx *cli.Context) error {
					log_cmd.StartLog(clictx)
					return nil
				},
			},
			{
				Name:  CMD_NAME_DB,
				Usage: "db command",
				Subcommands: []*cli.Command{
					{
						Name:  "init",
						Usage: "initialize db data",
						Action: func(clictx *cli.Context) error {
							fmt.Println("======== start of db data initialization ========")
							db_cmd.StartDBComponent(config.Get_config().Toml_config)
							db_cmd.Initialize()
							fmt.Println("======== end  of  db data initialization ========")
							return nil
						},
					},
					{
						Name:  "reconfig",
						Usage: "reconfig db data",
						Action: func(clictx *cli.Context) error {
							fmt.Println("======== start of db data reconfiguration ========")
							db_cmd.StartDBComponent(config.Get_config().Toml_config)
							db_cmd.Reconfig()
							fmt.Println("======== end  of  db data reconfiguration ========")
							return nil
						},
					},
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
							configs, _ := config.Get_config().Read_merge_config()
							fmt.Println(configs)
							fmt.Println("======== end  of  config ========")
							return nil
						},
					},
					//set config
					{
						Name:  "set",
						Usage: "set config",
						Flags: append(Cli_get_flags(), &cli.StringFlag{Name: "config", Required: false}),
						Action: func(clictx *cli.Context) error {
							return Cli_set_config(clictx)
						},
					},
				},
			},
		},
	}
}
