package conf

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/coreservice-io/utils/path_util"
	"github.com/pelletier/go-toml"
)

/////////////////////////////
type Config struct {
	Static_config_tree *toml.Tree
	Static_config_path string

	Custom_config_tree *toml.Tree
	Custom_config_path string

	Merge_config_tree *toml.Tree
	Toml_config       *TomlConfig
}

var config *Config

func Get_config() *Config {
	return config
}

func (config *Config) Read_merge_config() (string, error) {
	config_str, err := toml.Marshal(config.Toml_config)
	if err != nil {
		return "", err
	}

	return string(config_str), nil
}

func (config *Config) Save_custom_config() error {
	custom_config_str, err := config.Custom_config_tree.ToTomlString()
	if err != nil {
		return err
	}

	f, err := os.OpenFile(config.Custom_config_path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}

	defer f.Close()

	err = f.Truncate(0)
	if err != nil {
		return err
	}
	_, err = f.Seek(0, 0)
	if err != nil {
		return err
	}

	_, err = f.Write([]byte(custom_config_str))
	if err != nil {
		return err
	}

	return nil
}

func Init_config(static_config_path string, custom_config_path string) error {

	if config != nil {
		return nil
	}

	var cfg Config
	var err error

	//read static config
	s_c_p, s_c_p_exist, _ := path_util.SmartPathExist(static_config_path)
	if !s_c_p_exist {
		return errors.New("no config file:" + static_config_path)
	}
	cfg.Static_config_path = s_c_p
	cfg.Static_config_tree, err = toml.LoadFile(s_c_p)
	if err != nil {
		return err
	}

	//read custom config
	c_c_p := path_util.ExE_Path(custom_config_path)
	cfg.Custom_config_path = c_c_p
	_, c_c_p_exist, _ := path_util.SmartPathExist(c_c_p)
	if !c_c_p_exist {
		dir := filepath.Dir(c_c_p)
		os.MkdirAll(dir, 0777)
		cfg.Custom_config_tree, err = toml.Load("")
	} else {
		cfg.Custom_config_tree, err = toml.LoadFile(c_c_p)
	}
	if err != nil {
		return err
	}

	cfg.Merge_config_tree, err = mergeConfig(cfg.Custom_config_tree, cfg.Static_config_tree)
	if err != nil {
		return err
	}

	cfg.Toml_config = &TomlConfig{}
	err = cfg.Merge_config_tree.Unmarshal(cfg.Toml_config)
	if err != nil {
		return err
	}

	config = &cfg

	return nil
}

func mergeConfig(src *toml.Tree, base *toml.Tree) (*toml.Tree, error) {
	baseStr, err := base.ToTomlString()
	if err != nil {
		return nil, err
	}

	tree_merge, _ := toml.Load(baseStr)
	flat_map := map[string]interface{}{}
	readToFlat(src, "", flat_map)

	for k, v := range flat_map {
		tree_merge.Set(k, v)
	}

	return tree_merge, nil
}

func readToFlat(tree *toml.Tree, parent_key string, flat_map map[string]interface{}) {
	for _, key := range tree.Keys() {
		newKey := ""
		if parent_key == "" {
			newKey = key
		} else {
			newKey = parent_key + "." + key
		}
		value := tree.Get(key)
		switch value.(type) {
		case *toml.Tree:
			readToFlat(value.(*toml.Tree), newKey, flat_map)
		//case []*toml.Tree:
		//	for _, t := range value.([]*toml.Tree) {
		//		FlatRead(t, newKey, flat_map)
		//	}
		default:
			flat_map[newKey] = value
		}
	}
}
