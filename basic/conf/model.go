package conf

type TomlConfig struct {
	Log           Log           `toml:"log"`
	Http          HttpConfig    `toml:"http"`
	Https         HttpsConfig   `toml:"https"`
	Auto_cert     AutoCert      `toml:"auto_cert"`
	Api           API           `toml:"api"`
	Redis         Redis         `toml:"redis"`
	Db            DB            `toml:"db"`
	Elasticsearch ElasticSearch `toml:"elasticsearch"`
	IpLocal       IpLocal       `toml:"ip_local"`
	IpRemote      IpRemote      `toml:"ip_remote"`
	Leveldb       LevelDB       `toml:"leveldb"`
	Smtp          SMTP          `toml:"smtp"`
	Sqlite        Sqlite        `toml:"sqlite"`
}

type Log struct {
	Level string `toml:"level"`
}

type API struct {
	Doc_gen_search_dir string `toml:"doc_gen_search_dir"`
	Doc_gen_mainfile   string `toml:"doc_gen_mainfile"`
	Doc_gen_output_dir string `toml:"doc_gen_output_dir"`
}

type HttpConfig struct {
	Enable bool `toml:"enable"`
	Port   int  `toml:"port"`
}

type HttpsConfig struct {
	Enable   bool   `toml:"enable"`
	Port     int    `toml:"port"`
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

type IpLocal struct {
	Enable           bool   `toml:"enable"`
	Db_path          string `toml:"db_path"`
	Upgrade_interval int    `toml:"upgrade_interval"`
	Upgrade_url      string `toml:"upgrade_url"`
}

type IpRemote struct {
	Enable      bool           `toml:"enable"`
	Ipstack_key string         `toml:"ipstack_key"`
	Redis       IpRemote_Redis `toml:"redis"`
}

type IpRemote_Redis struct {
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
