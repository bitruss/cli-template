package config

import (
	"github.com/coreservice-io/cli-template/configuration"
	"github.com/urfave/cli/v2"
)

//set your config params types
var stringConfParams = []string{}
var float64ConfParams = []string{}
var boolConfParams = []string{}
var intConfParams = []string{}
var otherConf = []string{}

//get all config flags
func GetFlags() (allflags []cli.Flag) {
	allConfig := configuration.Config.AllSettings()
	mode := "default"
	iMode, exist := allConfig["mode"]
	if exist {
		mode = iMode.(string)
	}

	for k, v := range allConfig[mode].(map[string]interface{}) {
		switch v.(type) {
		case string:
			stringConfParams = append(stringConfParams, k)
		case float64:
			float64ConfParams = append(float64ConfParams, k)
		case int:
			intConfParams = append(intConfParams, k)
		case bool:
			boolConfParams = append(boolConfParams, k)
		}
	}

	for _, name := range stringConfParams {
		allflags = append(allflags, &cli.StringFlag{Name: name, Required: false})
	}

	for _, name := range float64ConfParams {
		allflags = append(allflags, &cli.Float64Flag{Name: name, Required: false})
	}

	for _, name := range intConfParams {
		allflags = append(allflags, &cli.IntFlag{Name: name, Required: false})
	}

	for _, name := range boolConfParams {
		allflags = append(allflags, &cli.BoolFlag{Name: name, Required: false})
	}

	//other custom flags
	allflags = append(allflags, &cli.StringFlag{Name: "addpath", Required: false})
	allflags = append(allflags, &cli.StringFlag{Name: "removepath", Required: false})

	return
}
