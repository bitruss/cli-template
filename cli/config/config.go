package config

import (
	"errors"
	"io/ioutil"

	"github.com/spf13/viper"
)

type VConfig struct {
	*viper.Viper
	configPath string
}

//func NewViperConfig() *VConfig {
//	return &VConfig{
//		viper.New(),
//		"",
//	}
//}

func ReadConfig(configPath string) (*VConfig, error) {
	c := &VConfig{viper.New(), ""}
	c.SetConfigFile(configPath)
	err := c.ReadInConfig()
	if err != nil {
		return nil, err
	}
	c.configPath = configPath
	return c, nil
}

//GetBool(key string) : bool
//GetFloat64(key string) : float64
//GetInt(key string) : int
//GetIntSlice(key string) : []int
//GetString(key string) : string
//GetStringMap(key string) : map[string]interface{}
//GetStringMapString(key string) : map[string]string
//GetStringSlice(key string) : []string
//GetTime(key string) : time.Time
//GetDuration(key string) : time.Duration
//AllSettings() : map[string]interface{}

func (c *VConfig) Get(key string, defaultValue interface{}) interface{} {
	if !c.Viper.IsSet(key) {
		return defaultValue
	}
	return c.Viper.Get(key)
}

func (c *VConfig) GetBool(key string, defaultValue bool) (bool, error) {
	if !c.Viper.IsSet(key) {
		return defaultValue, nil
	}

	v := c.Viper.Get(key)
	value, ok := v.(bool)
	if !ok {
		return false, errors.New("type error")
	}
	return value, nil
}

func (c *VConfig) GetFloat64(key string, defaultValue float64) (float64, error) {
	if !c.Viper.IsSet(key) {
		return defaultValue, nil
	}

	v := c.Viper.Get(key)
	value, ok := v.(float64)
	if !ok {
		return 0, errors.New("type error")
	}
	return value, nil
}

func (c *VConfig) GetInt(key string, defaultValue int) (int, error) {
	if !c.Viper.IsSet(key) {
		return defaultValue, nil
	}

	v := c.Viper.Get(key)
	value, ok := v.(int)
	if !ok {
		return 0, errors.New("type error")
	}
	return value, nil
}

func (c *VConfig) GetIntSlice(key string, defaultValue []int) ([]int, error) {
	if !c.Viper.IsSet(key) {
		return defaultValue, nil
	}

	v := c.Viper.Get(key)
	value, ok := v.([]int)
	if !ok {
		return nil, errors.New("type error")
	}
	return value, nil
}

func (c *VConfig) GetString(key string, defaultValue string) (string, error) {
	if !c.Viper.IsSet(key) {
		return defaultValue, nil
	}

	v := c.Viper.Get(key)
	value, ok := v.(string)
	if !ok {
		return "", errors.New("type error")
	}
	return value, nil
}

func (c *VConfig) GetStringMap(key string, defaultValue map[string]interface{}) (map[string]interface{}, error) {
	if !c.Viper.IsSet(key) {
		return defaultValue, nil
	}

	v := c.Viper.Get(key)
	value, ok := v.(map[string]interface{})
	if !ok {
		return nil, errors.New("type error")
	}
	return value, nil
}

func (c *VConfig) GetStringSlice(key string, defaultValue []string) ([]string, error) {
	if !c.Viper.IsSet(key) {
		return defaultValue, nil
	}

	v := c.Viper.Get(key)
	value, ok := v.([]string)
	if !ok {
		return nil, errors.New("type error")
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
