package cliCmd

import (
	"errors"

	"github.com/universe-30/UUtils/path_util"
	"github.com/urfave/cli/v2"
)

type Cmd struct {
	CmdName    string
	CliContext *cli.Context
}

var CmdToDo *Cmd

const CMD_NAME_DEFAULT = "default"
const CMD_NAME_LOG = "logs"
const CMD_NAME_SERVICE = "service"
const CMD_NAME_CONFIG = "config"

////////config to do cmd ///////////
func configCliCmd() *cli.App {

	var todoerr error

	return &cli.App{
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "dev", Required: false},
		},
		Action: func(c *cli.Context) error {
			CmdToDo, todoerr = getCmdToDo(CMD_NAME_DEFAULT, true, c)
			if todoerr != nil {
				return todoerr
			}
			return nil
		},

		Commands: []*cli.Command{
			{
				Name:    CMD_NAME_LOG,
				Aliases: []string{CMD_NAME_LOG},
				Usage:   "print all logs ",
				Flags: []cli.Flag{
					&cli.IntFlag{Name: "num", Required: false},
					&cli.BoolFlag{Name: "onlyerr", Required: false},
				},
				Action: func(c *cli.Context) error {
					CmdToDo, todoerr = getCmdToDo(CMD_NAME_LOG, false, c)
					if todoerr != nil {
						return todoerr
					}
					return nil
				},
			},
			{
				Name:    CMD_NAME_CONFIG,
				Aliases: []string{CMD_NAME_CONFIG},
				Usage:   "config command",
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "dev", Required: false},
				},
				Action: func(c *cli.Context) error {
					CmdToDo, todoerr = getCmdToDo(CMD_NAME_CONFIG, true, c)
					if todoerr != nil {
						return todoerr
					}
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
						Action: func(c *cli.Context) error {
							CmdToDo, todoerr = getCmdToDo(CMD_NAME_SERVICE, false, c)
							if todoerr != nil {
								return todoerr
							}
							return nil
						},
					},
					//service remove
					{
						Name:  "remove",
						Usage: "remove meson node from service",
						Action: func(c *cli.Context) error {
							CmdToDo, todoerr = getCmdToDo(CMD_NAME_SERVICE, false, c)
							if todoerr != nil {
								return todoerr
							}
							return nil
						},
					},
					//service start
					{
						Name:  "start",
						Usage: "run",
						Action: func(c *cli.Context) error {
							CmdToDo, todoerr = getCmdToDo(CMD_NAME_SERVICE, false, c)
							if todoerr != nil {
								return todoerr
							}
							return nil
						},
					},
					//service stop
					{
						Name:  "stop",
						Usage: "stop",
						Action: func(c *cli.Context) error {
							CmdToDo, todoerr = getCmdToDo(CMD_NAME_SERVICE, false, c)
							if todoerr != nil {
								return todoerr
							}
							return nil
						},
					},
					//service restart
					{
						Name:  "restart",
						Usage: "restart",
						Action: func(c *cli.Context) error {
							CmdToDo, todoerr = getCmdToDo(CMD_NAME_SERVICE, false, c)
							if todoerr != nil {
								return todoerr
							}
							return nil
						},
					},
					//service status
					{
						Name:  "status",
						Usage: "show process status",
						Action: func(c *cli.Context) error {
							CmdToDo, todoerr = getCmdToDo(CMD_NAME_SERVICE, false, c)
							if todoerr != nil {
								return todoerr
							}
							return nil
						},
					},
				},
			},
		},
	}
}

////////end config to do app ///////////
func readDefaultConfig(isDev bool) (*VConfig, string, error) {
	var defaultConfigPath string
	if isDev {
		Logger.Infoln("======== using dev mode ========")
		defaultConfigPath = path_util.GetAbsPath("configs/dev.json")
	} else {
		Logger.Infoln("======== using pro mode ========")
		defaultConfigPath = path_util.GetAbsPath("configs/pro.json")
	}

	Logger.Infoln("config file:", defaultConfigPath)

	config, err := ReadConfig(defaultConfigPath)
	if err != nil {
		Logger.Errorln("no pro.json under /configs folder , use --dev=true to run dev mode")
		return nil, "", err
	} else {
		return config, defaultConfigPath, nil
	}
}

func getCmdToDo(cmdName string, needconfig bool, c *cli.Context) (*Cmd, error) {
	path_util.ExEPathPrintln()

	app := &Cmd{
		CmdName:    cmdName,
		CliContext: c,
	}

	if needconfig {
		////read default config
		config, _, err := readDefaultConfig(c.Bool("dev"))
		if err != nil {
			return nil, err
		}
		Logger.Infoln("======== start of config ========")
		configs, _ := config.GetConfigAsString()
		Logger.Infoln(configs)
		Logger.Infoln("======== end of config ========")

		//ConfigFile = defaultConfigPath
		Config = config
	}

	logLevel := "INFO"
	if Config != nil {
		var err error
		logLevel, err = Config.GetString("local_log_level", "INFO")
		if err != nil {
			return nil, errors.New("local_log_level [string] in config err:" + err.Error())
		}
	}
	SetLogLevel(logLevel)
	return app, nil
}

// ManualInitAppConfig init app config when use go test
func ManualInitAppConfig(configPath string) {
	Logger.Infoln("configPath:", configPath)
	config, err := ReadConfig(configPath)
	if err != nil {
		panic("Manual read config err " + err.Error())
	}

	//ConfigFile = configPath
	Config = config

	logLevel := "INFO"
	logLevel, err = Config.GetString("local_log_level", "DEBU")
	if err != nil {
		panic("local_log_level [string] in config err:" + err.Error())
	}
	SetLogLevel(logLevel)
}
