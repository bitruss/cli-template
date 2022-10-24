package cmd

import (
	"fmt"
	"log"

	"os"
	"strings"

	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/basic/conf"
	"github.com/coreservice-io/cli-template/cmd/config"
	"github.com/coreservice-io/cli-template/cmd/default_"
	"github.com/coreservice-io/cli-template/cmd/default_/http/api"
	cmd_log "github.com/coreservice-io/cli-template/cmd/log"
	"github.com/coreservice-io/cli-template/db"
	ilog "github.com/coreservice-io/log"
	"github.com/urfave/cli/v2"
)

const CMD_NAME_DEFAULT = "default"
const CMD_NAME_GEN_API = "gen_api"
const CMD_NAME_LOG = "log"
const CMD_NAME_DB = "db"
const CMD_NAME_CONFIG = "config"

// //////config to do cmd ///////////
func ConfigCmd() *cli.App {

	//////////init config/////////////
	toml_target := "default"

	real_args := []string{}
	for _, arg := range os.Args {
		arg_lower := strings.ToLower(arg)
		if strings.HasPrefix(arg_lower, "-conf=") || strings.HasPrefix(arg_lower, "--conf=") {
			toml_target = strings.TrimPrefix(arg_lower, "--conf=")
			toml_target = strings.TrimPrefix(toml_target, "-conf=")
			continue
		}
		real_args = append(real_args, arg)
	}

	os.Args = real_args
	conf_err := conf.Init_config(toml_target)
	if conf_err != nil {
		log.Fatal("config err", conf_err)
	}

	configuration := conf.Get_config()

	/////set up basic logger ///////
	basic.InitLogger()

	/////set loglevel//////
	loglevel := ilog.ParseLogLevel(configuration.Toml_config.Log.Level)
	basic.Logger.SetLevel(loglevel)
	basic.Logger.Infoln("loglevel used:", ilog.LogLevelToTag(loglevel))
	////////////////////////////////

	var defaultAction = func(clictx *cli.Context) error {
		default_.StartDefault(clictx)
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
				Name:  CMD_NAME_LOG,
				Usage: "print all logs",
				Flags: cmd_log.GetFlags(),
				Action: func(clictx *cli.Context) error {
					cmd_log.StartLog(clictx)
					return nil
				},
			},
			{
				Name:  CMD_NAME_DB,
				Usage: "db command",
				Subcommands: []*cli.Command{
					{
						Name:  "init_data",
						Usage: "create initial data",
						Action: func(clictx *cli.Context) error {
							fmt.Println("======== start of db init_data ========")
							db.InitData()
							fmt.Println("======== end  of  db init_data ========")
							return nil
						},
					},
				},
			},
			{
				Name:  CMD_NAME_GEN_API,
				Usage: "api command",
				Action: func(clictx *cli.Context) error {
					api.GenApiDocs()
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
							configs, _ := conf.Get_config().Read_merge_config()
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
							return config.Cli_set_config(clictx)
						},
					},
				},
			},
		},
	}
}
