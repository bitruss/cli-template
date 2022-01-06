package config

import "github.com/urfave/cli/v2"

//set your config params types
var IntConfParams = []string{"http_port", "http_host"}
var StringConfParams = []string{"db_host", "db_name"}

//get all config flags
func GetFlags() (allflags []cli.Flag) {

	allflags = append(allflags, &cli.IntFlag{Name: "dev", Required: false})
	for _, name := range IntConfParams {
		allflags = append(allflags, &cli.IntFlag{Name: name, Required: false})
	}

	for _, name := range StringConfParams {
		allflags = append(allflags, &cli.StringFlag{Name: name, Required: false})
	}
	return
}
