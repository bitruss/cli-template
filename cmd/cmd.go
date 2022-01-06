package cmd

import (
	"github.com/universe-30/CliAppTemplate/basic"
	"github.com/universe-30/CliAppTemplate/cmd/config"
	"github.com/universe-30/CliAppTemplate/cmd/default_"
	"github.com/universe-30/CliAppTemplate/cmd/log"
	"github.com/universe-30/CliAppTemplate/cmd/service"
	"github.com/universe-30/UUtils/path_util"
	"github.com/urfave/cli/v2"
)

const CMD_NAME_DEFAULT = "default"
const CMD_NAME_LOG = "logs"
const CMD_NAME_SERVICE = "service"
const CMD_NAME_CONFIG = "config"

////////config to do cmd ///////////
func ConfigCmd() *cli.App {

	return &cli.App{
		Flags: default_.GetFlags(),
		Action: func(clictx *cli.Context) error {
			path_util.ExEPathPrintln()
			conferr := iniConfig(clictx)
			if conferr != nil {
				return conferr
			}
			logerr := iniLogger()
			if logerr != nil {
				return logerr
			}
			default_.StartDefault(clictx)
			return nil
		},

		Commands: []*cli.Command{
			{
				Name: CMD_NAME_LOG,
				//Aliases: []string{CMD_NAME_LOG},
				Usage: "print all logs",
				Flags: log.GetFlags(),
				Action: func(clictx *cli.Context) error {
					path_util.ExEPathPrintln()
					logerr := iniLogger()
					if logerr != nil {
						return logerr
					}
					log.StartLog(clictx)
					return nil
				},
			},
			{
				Name: CMD_NAME_CONFIG,
				//Aliases: []string{CMD_NAME_CONFIG},
				Usage: "config command",
				Flags: config.GetFlags(),
				Action: func(clictx *cli.Context) error {
					path_util.ExEPathPrintln()
					conferr := iniConfig(clictx)
					if conferr != nil {
						return conferr
					}
					logerr := iniLogger()
					if logerr != nil {
						return logerr
					}
					config.ConfigSetting(clictx)
					return nil
				},
			},
			{
				Name:    CMD_NAME_SERVICE,
				Aliases: []string{CMD_NAME_SERVICE},
				Usage:   "service command",
				Subcommands: []*cli.Command{
					//service install
					{
						Name:  "install",
						Usage: "install meson node in service",
						Action: func(clictx *cli.Context) error {
							logerr := iniLogger()
							if logerr != nil {
								return logerr
							}
							service.RunServiceCmd(clictx)
							return nil
						},
					},
					//service remove
					{
						Name:  "remove",
						Usage: "remove meson node from service",
						Action: func(clictx *cli.Context) error {
							logerr := iniLogger()
							if logerr != nil {
								return logerr
							}
							service.RunServiceCmd(clictx)
							return nil
						},
					},
					//service start
					{
						Name:  "start",
						Usage: "run",
						Action: func(clictx *cli.Context) error {
							logerr := iniLogger()
							if logerr != nil {
								return logerr
							}
							service.RunServiceCmd(clictx)
							return nil
						},
					},
					//service stop
					{
						Name:  "stop",
						Usage: "stop",
						Action: func(clictx *cli.Context) error {
							logerr := iniLogger()
							if logerr != nil {
								return logerr
							}
							service.RunServiceCmd(clictx)
							return nil
						},
					},
					//service restart
					{
						Name:  "restart",
						Usage: "restart",
						Action: func(clictx *cli.Context) error {
							logerr := iniLogger()
							if logerr != nil {
								return logerr
							}
							service.RunServiceCmd(clictx)
							return nil
						},
					},
					//service status
					{
						Name:  "status",
						Usage: "show process status",
						Action: func(clictx *cli.Context) error {
							logerr := iniLogger()
							if logerr != nil {
								return logerr
							}
							service.RunServiceCmd(clictx)
							return nil
						},
					},
				},
			},
		},
	}
}

////////end config to do app ///////////
func readDefaultConfig(isDev bool) (*basic.VConfig, string, error) {
	var defaultConfigPath string
	if isDev {
		basic.Logger.Infoln("======== using dev mode ========")
		defaultConfigPath = path_util.GetAbsPath("configs/dev.json")
	} else {
		basic.Logger.Infoln("======== using pro mode ========")
		defaultConfigPath = path_util.GetAbsPath("configs/pro.json")
	}

	basic.Logger.Infoln("config file:", defaultConfigPath)

	config, err := basic.ReadConfig(defaultConfigPath)
	if err != nil {
		basic.Logger.Errorln("no pro.json under /configs folder , use --dev=true to run dev mode")
		return nil, "", err
	} else {
		return config, defaultConfigPath, nil
	}
}

func iniConfig(clictx *cli.Context) error {
	path_util.ExEPathPrintln()
	////read default config
	config, _, err := readDefaultConfig(clictx.Bool("dev"))
	if err != nil {
		return err
	}
	basic.Logger.Infoln("======== start of config ========")
	configs, _ := config.GetConfigAsString()
	basic.Logger.Infoln(configs)
	basic.Logger.Infoln("======== end  of  config ========")
	basic.Config = config
	return nil
}

func iniLogger() error {
	logLevel := "INFO"
	if basic.Config != nil {
		var err error
		logLevel, err = basic.Config.GetString("local_log_level", "INFO")
		if err != nil {
			return err
		}
	}
	basic.SetLogLevel(logLevel)
	return nil
}

// ManualInitAppConfig init app config when use go test
// func ManualInitCmdConfig(configPath string) {
// 	basic.Logger.Infoln("configPath:", configPath)
// 	config, err := basic.ReadConfig(configPath)
// 	if err != nil {
// 		panic("Manual read config err " + err.Error())
// 	}

// 	//ConfigFile = configPath
// 	basic.Config = config

// 	logLevel := "INFO"
// 	logLevel, err = basic.Config.GetString("local_log_level", "DEBU")
// 	if err != nil {
// 		panic("local_log_level [string] in config err:" + err.Error())
// 	}
// 	basic.SetLogLevel(logLevel)
// }
