package configuration

import (
	"errors"
)

type ProvideFolder struct {
	AbsPath string `json:"abs_path"`
	SizeGB  int    `json:"size_GB"`
}

//example read provide_folder
func (c *VConfig) GetProvideFolders() ([]ProvideFolder, error) {
	key := "provide_folder"
	if !c.Viper.IsSet(key) {
		return nil, errors.New("provide_folder not find in config")
	}

	provide_folder := c.Viper.Get("provide_folder")
	iArray, ok := provide_folder.([]interface{})
	if !ok {
		return nil, errors.New("provide_folder in config type error")
	}

	provideFolders := []ProvideFolder{}
	for _, v := range iArray {
		m, ok := v.(map[string]interface{})
		if !ok {
			return nil, errors.New("provide_folder content type error")
		}
		pf := ProvideFolder{
			AbsPath: m["abs_path"].(string),
			SizeGB:  int(m["size_GB"].(float64)),
		}
		provideFolders = append(provideFolders, pf)
	}
	return provideFolders, nil
}

func SetProvideFolders(pf []ProvideFolder) {
	Config.Set("provide_folder", pf)
}
