package configuration

import (
	"errors"
	"io/ioutil"

	"github.com/coreservice-io/cli-template/basic"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

var Config *VConfig

type VConfig struct {
	*viper.Viper
	configPath string
	mode       string
}

func ReadConfig(configPath string) (*VConfig, error) {
	c := &VConfig{viper.New(), "", ""}
	c.SetConfigFile(configPath)
	err := c.ReadInConfig()
	if err != nil {
		return nil, err
	}
	c.configPath = configPath

	//read mode
	mode := "default"
	if c.Viper.IsSet("mode") {
		v := c.Viper.Get("mode")
		value, err := cast.ToStringE(v)
		if err != nil {
			return nil, errors.New("read mode in config err" + err.Error())
		} else {
			if value == "" {
				return nil, errors.New("mode value can not be empty")
			}
			mode = value
		}
	}
	basic.Logger.Infoln("config mode:", mode)
	c.mode = mode
	return c, nil
}

func (c *VConfig) getConfigKey(key string) string {
	vKey := c.mode + "." + key
	if !c.Viper.IsSet(vKey) {
		return "default." + key
	}

	return "default." + key
}

//func (c *VConfig) Set(key string, value interface{}) {
//	vKey := "default." + key
//	if c.mode != "" {
//		vKey = c.mode + "." + key
//	}
//
//	c.Viper.Set(vKey, value)
//}

func (c *VConfig) Get(key string, defaultValue interface{}) interface{} {
	vKey := c.getConfigKey(key)

	if !c.Viper.IsSet(vKey) {
		return defaultValue
	}

	return c.Viper.Get(vKey)
}

func (c *VConfig) GetBool(key string, defaultValue bool) (bool, error) {
	vKey := c.getConfigKey(key)

	if !c.Viper.IsSet(vKey) {
		return defaultValue, nil
	}

	v := c.Viper.Get(vKey)
	value, err := cast.ToBoolE(v)
	if err != nil {
		return false, err
	}
	return value, nil
}

func (c *VConfig) GetFloat64(key string, defaultValue float64) (float64, error) {
	vKey := c.getConfigKey(key)

	if !c.Viper.IsSet(vKey) {
		return defaultValue, nil
	}

	v := c.Viper.Get(vKey)
	value, err := cast.ToFloat64E(v)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func (c *VConfig) GetInt(key string, defaultValue int) (int, error) {
	vKey := c.getConfigKey(key)

	if !c.Viper.IsSet(vKey) {
		return defaultValue, nil
	}

	v := c.Viper.Get(vKey)
	value, err := cast.ToIntE(v)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func (c *VConfig) GetIntSlice(key string, defaultValue []int) ([]int, error) {
	vKey := c.getConfigKey(key)

	if !c.Viper.IsSet(vKey) {
		return defaultValue, nil
	}

	v := c.Viper.Get(vKey)
	value, err := cast.ToIntSliceE(v)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (c *VConfig) GetString(key string, defaultValue string) (string, error) {
	vKey := c.getConfigKey(key)

	if !c.Viper.IsSet(vKey) {
		return defaultValue, nil
	}

	v := c.Viper.Get(vKey)
	value, err := cast.ToStringE(v)
	if err != nil {
		return "", err
	}
	return value, nil
}

func (c *VConfig) GetStringMap(key string, defaultValue map[string]interface{}) (map[string]interface{}, error) {
	vKey := c.getConfigKey(key)

	if !c.Viper.IsSet(vKey) {
		return defaultValue, nil
	}

	v := c.Viper.Get(vKey)
	value, err := cast.ToStringMapE(v)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (c *VConfig) GetStringSlice(key string, defaultValue []string) ([]string, error) {
	vKey := c.getConfigKey(key)

	if !c.Viper.IsSet(vKey) {
		return defaultValue, nil
	}

	v := c.Viper.Get(vKey)
	value, err := cast.ToStringSliceE(v)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (c *VConfig) GetConfigAsString() (string, error) {
	b, err := ioutil.ReadFile(c.configPath)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
