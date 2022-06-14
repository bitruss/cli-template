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
	ilog "github.com/coreservice-io/log"
	"github.com/kardianos/service"
	"github.com/urfave/cli/v2"
)

const CMD_NAME_DEFAULT = "default"
const CMD_NAME_GEN_API = "gen_api"
const CMD_NAME_LOG = "log"
const CMD_NAME_SERVICE = "service"
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

	conf_err := basic.Init_config(toml_conf_path)
	if conf_err != nil {
		basic.Logger.Fatalln("config err", conf_err)
	}

	conf := basic.Get_config()

	/////set loglevel//////
	basic.Logger.SetLevel(ilog.ParseLogLevel(conf.Toml_config.Log_level))
	////////////////////////////////

	return &cli.App{
		Action: func(clictx *cli.Context) error {
			OS_service_start(conf.Toml_config.Daemon_name, "run", func() {
				default_.StartDefault(clictx)
			})
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
							configs, _ := basic.Get_config().Read_config_file()
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
			{
				Name:  CMD_NAME_SERVICE,
				Usage: "service command",
				Subcommands: []*cli.Command{
					//service install
					{
						Name:  "install",
						Usage: "install service",
						Action: func(clictx *cli.Context) error {
							OS_service_start(conf.Toml_config.Daemon_name, "install", nil)
							return nil
						},
					},
					//service remove
					{
						Name:  "remove",
						Usage: "remove service",
						Action: func(clictx *cli.Context) error {
							OS_service_start(conf.Toml_config.Daemon_name, "remove", nil)
							return nil
						},
					},
					//service start
					{
						Name:  "start",
						Usage: "run",
						Action: func(clictx *cli.Context) error {
							OS_service_start(conf.Toml_config.Daemon_name, "start", nil)
							return nil
						},
					},
					//service stop
					{
						Name:  "stop",
						Usage: "stop",
						Action: func(clictx *cli.Context) error {
							OS_service_start(conf.Toml_config.Daemon_name, "stop", nil)
							return nil
						},
					},
					//service restart
					{
						Name:  "restart",
						Usage: "restart",
						Action: func(clictx *cli.Context) error {
							OS_service_start(conf.Toml_config.Daemon_name, "restart", nil)
							return nil
						},
					},
					//service status
					{
						Name:  "status",
						Usage: "show process status",
						Action: func(clictx *cli.Context) error {
							OS_service_start(conf.Toml_config.Daemon_name, "status", nil)
							return nil
						},
					},
				},
			},
		},
	}
}

/////////service/////////////

func OS_service_start(name string, action string, exe_func func()) {
	os_service_conf := &service.Config{
		Name:        name,
		DisplayName: name,
		Description: name + ":description",

		Option: map[string]interface{}{
			"OnFailure":              "restart",
			"OnFailureDelayDuration": "15s",
			"SystemdScript":          systemdScript,
			"Restart":                "on-failure", // or use "always"
		},
	}

	oss, err := service.New(OS_service_program{Exe_func: exe_func}, os_service_conf)
	if err != nil {
		basic.Logger.Fatalln(err)
	}

	os_service_run(&oss, action)
}

type OS_service_program struct {
	Exe_func func()
}

func (p OS_service_program) Start(s service.Service) error {
	if p.Exe_func != nil {
		go p.Exe_func()
	}
	return nil
}

func (p OS_service_program) Stop(s service.Service) error {
	return nil
}

func os_service_run(s *service.Service, action string) {
	switch action {
	case "install":
		err := (*s).Install()
		if err != nil {
			basic.Logger.Fatalln("install service error:", err)
		} else {
			basic.Logger.Infoln("service installed")
		}
	case "remove":
		err := (*s).Uninstall()
		if err != nil {
			basic.Logger.Fatalln("remove service error:", err)
		} else {
			basic.Logger.Infoln("service removed")
		}
	case "start":
		err := (*s).Start()
		if err != nil {
			basic.Logger.Fatalln("start service error:", err)
		} else {
			basic.Logger.Infoln("service started")
		}
	case "run":
		err := (*s).Run()
		if err != nil {
			basic.Logger.Fatalln("run service error:", err)
		} else {
			basic.Logger.Infoln("service run")
		}
	case "stop":
		err := (*s).Stop()
		if err != nil {
			basic.Logger.Fatalln("stop service error:", err)
		} else {
			basic.Logger.Infoln("service stopped")
		}
	case "restart":
		err := (*s).Restart()
		if err != nil {
			basic.Logger.Fatalln("restart service error:", err)
		} else {
			basic.Logger.Infoln("service restarted")
		}
	case "status":
		status, err := (*s).Status()
		if err != nil {
			basic.Logger.Fatalln(err)
		}
		switch status {
		case service.StatusRunning:
			basic.Logger.Infoln("service status:", "RUNNING")
		case service.StatusStopped:
			basic.Logger.Infoln("service status:", "STOPPED")
		default:
			basic.Logger.Infoln("service status:", "UNKNOWN")
		}
	default:
		basic.Logger.Warnln("no sub command")
		return
	}
}

const systemdScript = `[Unit]
Description={{.Description}}
ConditionFileIsExecutable={{.Path|cmdEscape}}
{{range $i, $dep := .Dependencies}} 
{{$dep}} {{end}}

[Service]
StartLimitInterval=15
StartLimitBurst=15
ExecStart={{.Path|cmdEscape}}{{range .Arguments}} {{.|cmd}}{{end}}
{{if .ChRoot}}RootDirectory={{.ChRoot|cmd}}{{end}}
{{if .WorkingDirectory}}WorkingDirectory={{.WorkingDirectory|cmdEscape}}{{end}}
{{if .UserName}}User={{.UserName}}{{end}}
{{if .ReloadSignal}}ExecReload=/bin/kill -{{.ReloadSignal}} "$MAINPID"{{end}}
{{if .PIDFile}}PIDFile={{.PIDFile|cmd}}{{end}}
{{if and .LogOutput .HasOutputFileSupport -}}
StandardOutput=file:/var/log/{{.Name}}.out
StandardError=file:/var/log/{{.Name}}.err
{{- end}}
{{if gt .LimitNOFILE -1 }}LimitNOFILE={{.LimitNOFILE}}{{end}}
{{if .Restart}}Restart={{.Restart}}{{end}}
{{if .SuccessExitStatus}}SuccessExitStatus={{.SuccessExitStatus}}{{end}}
RestartSec=60
EnvironmentFile=-/etc/sysconfig/{{.Name}}

[Install]
WantedBy=multi-user.target
`
