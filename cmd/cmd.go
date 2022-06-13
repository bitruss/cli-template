package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/cmd/config"
	"github.com/coreservice-io/cli-template/cmd/default_"
	"github.com/coreservice-io/cli-template/cmd/default_/http/api"
	"github.com/coreservice-io/cli-template/cmd/log"
	"github.com/coreservice-io/cli-template/cmd/service"
	"github.com/coreservice-io/cli-template/configuration"
	"github.com/coreservice-io/cli-template/plugin/daemon_plugin"
	ilog "github.com/coreservice-io/log"
	"github.com/coreservice-io/utils/path_util"
	daemonService "github.com/kardianos/service"
	"github.com/urfave/cli/v2"
)

const CMD_NAME_DEFAULT = "default"
const CMD_NAME_GEN_API = "gen_api"
const CMD_NAME_LOG = "log"
const CMD_NAME_SERVICE = "service"
const CMD_NAME_CONFIG = "config"

type Program struct {
	Clictx *cli.Context
}

func (p *Program) Start(s daemonService.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *Program) run() {
	// Do work here
	default_.StartDefault(p.Clictx)
}
func (p *Program) Stop(s daemonService.Service) error {
	// Stop should not block. Return with a few seconds.
	//basic.Logger.Infoln("service will stop in 5 seconds...")
	//<-time.After(time.Second * 5)
	return nil
}

////////config to do cmd ///////////
func ConfigCmd() *cli.App {
	//check is dev or pro
	confShow := false
	real_args := []string{}

	for _, arg := range os.Args {

		s := strings.ToLower(arg)

		if strings.Contains(s, "-conf=show") || strings.Contains(s, "--conf=show") {
			confShow = true
			continue
		}

		if strings.Contains(s, "-conf=hide") || strings.Contains(s, "--conf=hide") {
			confShow = false
			continue
		}

		real_args = append(real_args, arg)
	}

	os.Args = real_args

	conferr := iniConfig(confShow)
	if conferr != nil {
		basic.Logger.Panicln(conferr)
	}

	daemon_name, err := configuration.Config.GetString("daemon_name", "")
	if err != nil {
		basic.Logger.Fatalln("daemon_name [string] in config error," + err.Error())
	}

	if daemon_name == "" {
		basic.Logger.Fatalln("daemon_name in config should not be vacant")
	}

	p := &Program{}
	err = daemon_plugin.Init(daemon_name, p)
	if err != nil {
		basic.Logger.Fatalln("daemon_plugin.Init error:", err)
	}
	s := daemon_plugin.GetInstance(daemon_name)

	return &cli.App{
		Action: func(clictx *cli.Context) error {
			p.Clictx = clictx
			s.Run()
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
							configs, _ := configuration.Config.GetConfigAsString()
							fmt.Println(configs)
							fmt.Println("======== end  of  config ========")
							return nil
						},
					},
					//set config
					{
						Name:  "set",
						Usage: "set config",
						Flags: config.GetFlags(),
						Action: func(clictx *cli.Context) error {
							config.ConfigSetting(clictx)
							return nil
						},
					},
				},
			},
			{
				Name:  CMD_NAME_SERVICE,
				Usage: "service command",
				Subcommands: []*cli.Command{
					//service install
					{
						Name:  "install",
						Usage: "install service",
						Action: func(clictx *cli.Context) error {
							service.RunServiceCmd(clictx, s)
							return nil
						},
					},
					//service remove
					{
						Name:  "remove",
						Usage: "remove service",
						Action: func(clictx *cli.Context) error {
							service.RunServiceCmd(clictx, s)
							return nil
						},
					},
					//service start
					{
						Name:  "start",
						Usage: "run",
						Action: func(clictx *cli.Context) error {
							service.RunServiceCmd(clictx, s)
							return nil
						},
					},
					//service stop
					{
						Name:  "stop",
						Usage: "stop",
						Action: func(clictx *cli.Context) error {
							service.RunServiceCmd(clictx, s)
							return nil
						},
					},
					//service restart
					{
						Name:  "restart",
						Usage: "restart",
						Action: func(clictx *cli.Context) error {
							service.RunServiceCmd(clictx, s)
							return nil
						},
					},
					//service status
					{
						Name:  "status",
						Usage: "show process status",
						Action: func(clictx *cli.Context) error {
							service.RunServiceCmd(clictx, s)
							return nil
						},
					},
				},
			},
		},
	}
}

////////end config to do app ///////////
func readDefaultConfig(confShow bool) (*configuration.VConfig, string, error) {
	var defaultConfigPath string
	var err error

	c_p, c_p_exist, _ := path_util.SmartPathExist("configs/config.yaml")
	if !c_p_exist {
		basic.Logger.Errorln("no config.yaml under /configs folder")
		return nil, "", err
	}
	defaultConfigPath = c_p

	if confShow {
		basic.Logger.Infoln("using config:", defaultConfigPath)
	}

	config, err := configuration.ReadConfig(defaultConfigPath)
	if err != nil {
		basic.Logger.Errorln("config err", err)
		return nil, "", err
	}

	return config, defaultConfigPath, nil
}

func iniConfig(confShow bool) error {
	//path_util.ExEPathPrintln()
	////read default config
	config, _, err := readDefaultConfig(confShow)
	if err != nil {
		return err
	}

	configuration.Config = config
	logerr := setLoggerLevel()
	if logerr != nil {
		return logerr
	}

	if confShow {
		basic.Logger.Infoln("======== start of config ========")
		configs, _ := config.GetConfigAsString()
		basic.Logger.Infoln(configs)
		basic.Logger.Infoln("======== end  of  config ========")
	}

	return nil
}

func setLoggerLevel() error {
	logLevel := "INFO"
	if configuration.Config != nil {
		var err error
		logLevel, err = configuration.Config.GetString("log_level", "INFO")
		if err != nil {
			return err
		}
	}

	l := ilog.ParseLogLevel(logLevel)
	basic.Logger.SetLevel(l)
	return nil
}
