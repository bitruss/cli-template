package plugin

// import (
// 	"errors"

// 	tool_errors "github.com/coreservice-io/cli-template/tools/errors"

// 	"github.com/coreservice-io/cli-template/basic"
// 	"github.com/coreservice-io/cli-template/configuration"
// 	"github.com/coreservice-io/cli-template/plugin/ip_geo_plugin"
// 	"github.com/coreservice-io/cli-template/plugin/reference_plugin"
// 	"github.com/coreservice-io/ip_geo"
// 	"github.com/coreservice-io/utils/path_util"
// )

// func initIpGeo() error {
// 	ipStackAccessKey, err := configuration.Config.GetString("ip_geo.ipstack_key", "")
// 	if err != nil {
// 		return errors.New("ip_geo.ipstack_key [string] in config.json err," + err.Error())
// 	}
// 	upgradeUrl, err := configuration.Config.GetString("ip_geo.ip2l.upgrade_url", "")
// 	if err != nil {
// 		return errors.New("ip_geo.ip2l.upgrade_url [string] in config.json err," + err.Error())
// 	}
// 	upgradeInterval, err := configuration.Config.GetInt("ip_geo.ip2l.upgrade_interval", 172800)
// 	if err != nil {
// 		return errors.New("ip_geo.ip2l.upgrade_interval [string] in config.json err," + err.Error())
// 	}

// 	dbFilePath, err := configuration.Config.GetString("ip_geo.ip2l.db_path", "")
// 	if err != nil {
// 		return errors.New("ip_geo.ip2l.db_path [string] in config.json err," + err.Error())
// 	}

// 	redis_addr, err := configuration.Config.GetString("ip_geo.redis.host", "127.0.0.1")
// 	if err != nil {
// 		return errors.New("ip_geo.redis.host [string] in config.json err," + err.Error())
// 	}

// 	redis_username, err := configuration.Config.GetString("ip_geo.redis.username", "")
// 	if err != nil {
// 		return errors.New("ip_geo.redis.username [string] in config.json err," + err.Error())
// 	}

// 	redis_password, err := configuration.Config.GetString("ip_geo.redis.password", "")
// 	if err != nil {
// 		return errors.New("ip_geo.redis.password [string] in config.json err," + err.Error())
// 	}

// 	redis_port, err := configuration.Config.GetInt("ip_geo.redis.port", 6379)
// 	if err != nil {
// 		return errors.New("ip_geo.redis.port [int] in config.json err," + err.Error())
// 	}

// 	redis_prefix, err := configuration.Config.GetString("ip_geo.redis.prefix", "")
// 	if err != nil {
// 		return errors.New("ip_geo.redis.prefix [string] in config err," + err.Error())
// 	}

// 	redis_useTls, err := configuration.Config.GetBool("ip_geo.redis.use_tls", false)
// 	if err != nil {
// 		return errors.New("ip_geo.redis.use_tls [bool] in config err," + err.Error())
// 	}

// 	dbFilePath_abs, dbFilePath_abs_exist, _ := path_util.SmartPathExist(dbFilePath)
// 	if !dbFilePath_abs_exist {
// 		return errors.New("ip2Location db file path error," + dbFilePath)
// 	}
// 	ip_geo_redis_config := ip_geo.RedisConfig{
// 		Addr:     redis_addr,
// 		UserName: redis_username,
// 		Password: redis_password,
// 		Port:     redis_port,
// 		Prefix:   redis_prefix,
// 		UseTLS:   redis_useTls,
// 	}

// 	reference_plugin.Init_("ip_geo")
// 	return ip_geo_plugin.Init(ipStackAccessKey, dbFilePath_abs, upgradeUrl, int64(upgradeInterval),
// 		reference_plugin.GetInstance_("ip_geo"), ip_geo_redis_config, basic.Logger, tool_errors.PanicHandler)
// }
