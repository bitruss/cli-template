package config

import (
	"github.com/coreservice-io/cli-template/configuration"
	"github.com/urfave/cli/v2"
)

var allKeysMap = map[string]string{}

func assignConfigKeys(config interface{}, keysMap map[string]string, keyPre string) {
	if keyPre != "" {
		keyPre = keyPre + "."
	}
	configMap, ok := config.(map[string]interface{})
	if !ok {
		switch config.(type) {
		case string:
			keysMap[keyPre[:len(keyPre)-1]] = "string"
		case float64:
			keysMap[keyPre[:len(keyPre)-1]] = "float64"
		case int:
			keysMap[keyPre[:len(keyPre)-1]] = "int"
		case bool:
			keysMap[keyPre[:len(keyPre)-1]] = "bool"
		}
		return
	} else {
		if keyPre != "" {
			delete(keysMap, keyPre[:len(keyPre)-1])
		}
		for k, v := range configMap {
			keysMap[keyPre+k] = ""
			assignConfigKeys(v, keysMap, keyPre+k)
		}
	}
}

//get all config flags
func GetFlags() (allflags []cli.Flag) {
	allConfig := configuration.Config.AllSettings()

	modeType := map[string]struct{}{}
	for k, _ := range allConfig {
		if k == "mode" {
			continue
		}
		modeType[k] = struct{}{}
	}
	modeType["default"] = struct{}{}

	keysMap := map[string]string{}
	for mType := range modeType {
		configsInType, exist := allConfig[mType]
		if !exist {
			continue
		}

		assignConfigKeys(configsInType, keysMap, "")
	}

	for k := range modeType {
		for key, valueType := range keysMap {
			allKeysMap[k+"."+key] = valueType
			allflags = append(allflags, &cli.StringFlag{Name: k + "." + key, Required: false})
		}
	}

	allKeysMap["mode"] = "string"
	allflags = append(allflags, &cli.StringFlag{Name: "mode", Required: false})

	return
}
