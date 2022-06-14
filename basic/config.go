package basic

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/coreservice-io/utils/path_util"
	"github.com/creasty/defaults"
	"github.com/pelletier/go-toml/v2"
)

type TomlConfig struct {
	Daemon_name   string        `toml:"daemon_name"`
	Log_level     string        `toml:"log_level" default:"INFO"`
	Http          HttpConfig    `toml:"http"`
	Https         HttpsConfig   `toml:"https"`
	Auto_cert     AutoCert      `toml:"auto_cert"`
	Api           API           `toml:"api"`
	Redis         Redis         `toml:"redis"`
	Db            DB            `toml:"db"`
	Elasticsearch ElasticSearch `toml:"elasticsearch"`
	Ip_geo        IpGeo         `toml:"ip_geo"`
	Leveldb       LevelDB       `toml:"leveldb"`
	Smtp          SMTP          `toml:"smtp"`
	Sqlite        Sqlite        `toml:"sqlite"`
}

type API struct {
	Doc_gen_search_dir string `toml:"doc_gen_search_dir" default:"cmd/default_/http/api"`
	Doc_gen_mainfile   string `toml:"doc_gen_mainfile" default:"api.go"`
	Doc_gen_output_dir string `toml:"doc_gen_output_dir" default:"cmd/default_/http/api_docs"`
}

type HttpConfig struct {
	Enable bool `toml:"enable"`
	Port   int  `toml:"port" default:"80"`
}

type HttpsConfig struct {
	Enable   bool   `toml:"enable"`
	Port     int    `toml:"port" default:"443"`
	Crt_path string `toml:"crt_path" `
	Key_path string `toml:"key_path"`
	Html_dir string `toml:"html_dir"`
}

type AutoCert struct {
	Enable         bool   `toml:"enable"`
	Check_interval int    `toml:"check_interval"`
	Crt_path       string `toml:"crt_path"`
	Init_download  bool   `toml:"init_download"`
	Key_path       string `toml:"key_path"`
	Url            string `toml:"url"`
}

type Redis struct {
	Enable   bool   `toml:"enable"`
	Use_tls  bool   `toml:"use_tls"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Prefix   string `toml:"prefix"`
}

type DB struct {
	Enable   bool   `toml:"enable"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Name     string `toml:"name"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}

type ElasticSearch struct {
	Enable   bool   `toml:"enable"`
	Host     string `toml:"host"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}

type IpGeo struct {
	Enable      bool        `toml:"enable"`
	Ipstack_key string      `toml:"ipstack_key"`
	Ip2l        IpGeo_Ip2l  `toml:"ip2l"`
	Redis       IpGeo_Redis `toml:"redis"`
}

type IpGeo_Ip2l struct {
	Db_path          string `toml:"db_path"`
	Upgrade_interval int    `toml:"upgrade_interval"`
	Upgrade_url      string `toml:"upgrade_url"`
}

type IpGeo_Redis struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Prefix   string `toml:"prefix"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Use_tls  bool   `toml:"use_tls"`
}

type LevelDB struct {
	Enable bool   `toml:"enable"`
	Path   string `toml:"path"`
}

type SMTP struct {
	Enable   bool   `toml:"enable"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Password string `toml:"password"`
	Username string `toml:"username"`
}

type Sqlite struct {
	Enable bool   `toml:"enable"`
	Path   string `toml:"path"`
}

/////////////////////////////
type Config struct {
	Toml_config *TomlConfig
	Abs_path    string
}

var config *Config

func Get_config() *Config {
	return config
}

func (config *Config) Read_config_file() (string, error) {

	doc, err := ioutil.ReadFile(config.Abs_path)
	if err != nil {
		return "", err
	}

	return string(doc), nil
}

func (config *Config) Save_config() error {

	result, err := toml.Marshal(config.Toml_config)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(config.Abs_path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
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

	_, err = f.Write(result)
	if err != nil {
		return err
	}

	return nil
}

func Init_config(config_path string) error {

	if config != nil {
		return nil
	}

	c_p, c_p_exist, _ := path_util.SmartPathExist(config_path)
	if !c_p_exist {
		return errors.New("no config file:" + config_path)
	}

	var cfg Config
	cfg.Abs_path = c_p
	cfg.Toml_config = &TomlConfig{}

	//default value
	if err := defaults.Set(cfg.Toml_config); err != nil {
		return err
	}

	config_str, err := cfg.Read_config_file()
	if err != nil {
		return err
	}

	err = toml.Unmarshal([]byte(config_str), cfg.Toml_config)
	if err != nil {
		return err
	}

	Logger.Infoln("using config:", cfg.Abs_path)

	config = &cfg

	return nil
}
