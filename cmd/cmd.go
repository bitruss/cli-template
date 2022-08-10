package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/basic/conf"
	"github.com/coreservice-io/cli-template/cmd/config"
	"github.com/coreservice-io/cli-template/cmd/default_"
	"github.com/coreservice-io/cli-template/cmd/default_/http/api"
	"github.com/coreservice-io/cli-template/cmd/log"
	daemon "github.com/coreservice-io/daemon/daemon_util"
	ilog "github.com/coreservice-io/log"
	"github.com/urfave/cli/v2"
)

const CMD_NAME_DEFAULT = "default"
const CMD_NAME_GEN_API = "gen_api"
const CMD_NAME_LOG = "log"
const CMD_NAME_CONFIG = "config"

type DaemonProcess struct {
	cliContext *cli.Context
}

func (dp *DaemonProcess) Start() {
	go dp.Run()
	return
}

func (dp *DaemonProcess) Stop() {

}

func (dp *DaemonProcess) Run() {
	// Do work here
	default_.StartDefault(dp.cliContext)
}

type Service struct {
	daemon.Daemon
}

////////config to do cmd ///////////
func ConfigCmd() *cli.App {

	//////////init config/////////////
	toml_conf_path := "configs/default.toml"

	real_args := []string{}
	for _, arg := range os.Args {
		arg_lower := strings.ToLower(arg)
		if strings.HasPrefix(arg_lower, "-conf=") || strings.HasPrefix(arg_lower, "--conf=") {

			toml_target := strings.TrimPrefix(arg_lower, "--conf=")
			toml_target = strings.TrimPrefix(toml_target, "-conf=")
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
	loglevel := ilog.ParseLogLevel(configuration.Toml_config.Log.Level)
	basic.Logger.SetLevel(loglevel)
	basic.Logger.Infoln("loglevel used:", ilog.LogLevelToTag(loglevel))
	////////////////////////////////

	var defaultAction = func(clictx *cli.Context) error {
		//get exe path
		file, err := exec.LookPath(os.Args[0])
		if err != nil {
			basic.Logger.Fatalln("exec.LookPath err", err)
		}
		runPath, err := filepath.Abs(file)
		if err != nil {
			basic.Logger.Fatalln("filepath.Abs err", err)
		}
		//basic.Logger.Infoln("running file:", runPath)
		appName := filepath.Base(runPath)
		//basic.Logger.Infoln("fileName:", appName)

		//new daemon instance
		//daemon kind
		kind := daemon.SystemDaemon
		if runtime.GOOS == "darwin" {
			kind = daemon.UserAgent
		}
		daemonInstance, err := daemon.New(appName, appName+"service", kind)
		if err != nil {
			basic.Logger.Fatalln("daemon.New err", err)
		}
		dp := &DaemonProcess{
			cliContext: clictx,
		}
		s := &Service{
			Daemon: daemonInstance,
		}
		_, err = s.Run(dp)
		if err != nil {
			return err
		}
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
