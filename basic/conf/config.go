package conf

import (
	"errors"
	"os"
	"path"
	"path/filepath"

	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/utils/path_util"
	"github.com/pelletier/go-toml"
)

// ///////////////////////////
type Config struct {
	Root_config_tree *toml.Tree
	Root_config_path string

	User_config_tree *toml.Tree
	User_config_path string

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

func (config *Config) Save_user_config() error {
	user_config_str, err := config.User_config_tree.ToTomlString()
	if err != nil {
		return err
	}

	f, err := os.OpenFile(config.User_config_path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
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

	_, err = f.Write([]byte(user_config_str))
	if err != nil {
		return err
	}

	return nil
}

func Init_config(conf_target string) error {

	if config != nil {
		return nil
	}

	var cfg Config
	var err error

	//read root config
	root_conf_toml_rel_path := path.Join("root_conf", conf_target+".toml")
	r_c_p, r_c_p_exist, _ := path_util.SmartPathExist(root_conf_toml_rel_path)
	if !r_c_p_exist {
		return errors.New("no root config file:" + root_conf_toml_rel_path)
	}
	cfg.Root_config_path = r_c_p
	cfg.Root_config_tree, err = toml.LoadFile(r_c_p)
	if err != nil {
		return err
	}

	basic.Logger.Infoln("using root config toml file:", r_c_p)

	basic.WORK_DIR = path.Dir(path.Dir(r_c_p))

	basic.Logger.Infoln("--------------------------------------")
	basic.Logger.Infoln("working dir:", basic.WORK_DIR)
	basic.Logger.Infoln("--------------------------------------")

	//read user config

	user_conf_toml_rel_path := path.Join("user_conf", conf_target+".toml")
	user_conf_toml_abs_path, u_c_p_exist, err := basic.PathExist(user_conf_toml_rel_path)
	if err != nil {
		return err
	}

	cfg.User_config_path = user_conf_toml_abs_path
	basic.Logger.Infoln("using user config toml file:", user_conf_toml_abs_path)

	if !u_c_p_exist {
		dir := filepath.Dir(user_conf_toml_abs_path)
		os.MkdirAll(dir, 0777)
		cfg.User_config_tree, err = toml.Load("")
	} else {
		cfg.User_config_tree, err = toml.LoadFile(user_conf_toml_abs_path)
	}
	if err != nil {
		return err
	}

	cfg.Merge_config_tree, err = mergeConfig(cfg.User_config_tree, cfg.Root_config_tree)
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
		//todo
		switch value.(type) {
		case *toml.Tree:
			readToFlat(value.(*toml.Tree), newKey, flat_map)

		default:
			flat_map[newKey] = value
		}
	}
}
