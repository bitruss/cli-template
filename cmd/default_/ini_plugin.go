package default_

import (
	"errors"

	"github.com/coreservice-io/CliAppTemplate/basic"
	"github.com/coreservice-io/CliAppTemplate/configuration"
	"github.com/coreservice-io/CliAppTemplate/plugin/cache"
	"github.com/coreservice-io/CliAppTemplate/plugin/echoServer"
	"github.com/coreservice-io/CliAppTemplate/plugin/es"
	"github.com/coreservice-io/CliAppTemplate/plugin/hub"
	"github.com/coreservice-io/CliAppTemplate/plugin/redisClient"
	"github.com/coreservice-io/CliAppTemplate/plugin/sprMgr"
	"github.com/coreservice-io/CliAppTemplate/plugin/sqldb"
	"github.com/coreservice-io/RedisSpr"
)

func iniHub() error {
	return hub.Init()
}

func initEchoServer() error {
	http_port, err := configuration.Config.GetInt("http_port", 8080)
	if err != nil {
		return errors.New("http_port [int] in config error," + err.Error())
	}

	http_static_rel_folder, err := configuration.Config.GetString("http_static_rel_folder", "")
	if err != nil {
		return errors.New("http_static_rel_folder [string] in config error," + err.Error())
	}

	return echoServer.Init(echoServer.Config{Port: http_port, StaticFolder: http_static_rel_folder})
}

func initElasticSearch() error {
	elasticSearchAddr, err := configuration.Config.GetString("elasticsearch_addr", "")
	if err != nil {
		return errors.New("elasticsearch_addr [string] in config error," + err.Error())
	}

	elasticSearchUserName, err := configuration.Config.GetString("elasticsearch_username", "")
	if err != nil {
		return errors.New("elasticsearch_username_err [string] in config error," + err.Error())
	}

	elasticSearchPassword, err := configuration.Config.GetString("elasticsearch_password", "")
	if err != nil {
		return errors.New("elasticsearch_password [string] in config error," + err.Error())
	}

	return es.Init(es.Config{
		Address:  elasticSearchAddr,
		UserName: elasticSearchUserName,
		Password: elasticSearchPassword})

}

func initRedis() error {
	redis_addr, err := configuration.Config.GetString("redis_addr", "127.0.0.1")
	if err != nil {
		return errors.New("redis_addr [string] in config err," + err.Error())
	}

	redis_username, err := configuration.Config.GetString("redis_username", "")
	if err != nil {
		return errors.New("redis_username [string] in config err," + err.Error())
	}

	redis_password, err := configuration.Config.GetString("redis_password", "")
	if err != nil {
		return errors.New("redis_password [string] in config err," + err.Error())
	}

	redis_port, err := configuration.Config.GetInt("redis_port", 6379)
	if err != nil {
		return errors.New("redis_port [int] in config err," + err.Error())
	}

	redis_prefix, err := configuration.Config.GetString("redis_prefix", "")
	if err != nil {
		return errors.New("redis_prefix [string] in config err," + err.Error())
	}

	redis_useTls, err := configuration.Config.GetBool("redis_useTls", false)
	if err != nil {
		return errors.New("redis_useTls [bool] in config err," + err.Error())
	}

	return redisClient.Init(redisClient.Config{
		Address:   redis_addr,
		UserName:  redis_username,
		Password:  redis_password,
		Port:      redis_port,
		KeyPrefix: redis_prefix,
		UseTLS:    redis_useTls,
	})
}

func initSpr() error {
	redis_addr, err := configuration.Config.GetString("redis_addr", "127.0.0.1")
	if err != nil {
		return errors.New("redis_addr [string] in config.json err," + err.Error())
	}

	redis_username, err := configuration.Config.GetString("redis_username", "")
	if err != nil {
		return errors.New("redis_username [string] in config.json err," + err.Error())
	}

	redis_password, err := configuration.Config.GetString("redis_password", "")
	if err != nil {
		return errors.New("redis_password [string] in config.json err," + err.Error())
	}

	redis_port, err := configuration.Config.GetInt("redis_port", 6379)
	if err != nil {
		return errors.New("redis_port [int] in config.json err," + err.Error())
	}

	redis_prefix, err := configuration.Config.GetString("redis_prefix", "")
	if err != nil {
		return errors.New("redis_prefix [string] in config err," + err.Error())
	}

	redis_useTls, err := configuration.Config.GetBool("redis_useTls", false)
	if err != nil {
		return errors.New("redis_useTls [bool] in config err," + err.Error())
	}

	return sprMgr.Init(&RedisSpr.RedisConfig{
		Addr:     redis_addr,
		UserName: redis_username,
		Password: redis_password,
		Port:     redis_port,
		Prefix:   redis_prefix,
		UseTLS:   redis_useTls,
	})
}

func initDB() error {
	db_host, err := configuration.Config.GetString("db_host", "127.0.0.1")
	if err != nil {
		return errors.New("db_host [string] in config err," + err.Error())
	}

	db_port, err := configuration.Config.GetInt("db_port", 3306)
	if err != nil {
		return errors.New("db_port [int] in config err," + err.Error())
	}

	db_name, err := configuration.Config.GetString("db_name", "dbname")
	if err != nil {
		return errors.New("db_name [string] in config err," + err.Error())
	}

	db_username, err := configuration.Config.GetString("db_username", "username")
	if err != nil {
		return errors.New("db_username [string] in config err," + err.Error())
	}

	db_password, err := configuration.Config.GetString("db_password", "password")
	if err != nil {
		return errors.New("db_password [string] in config err," + err.Error())
	}

	return sqldb.Init(sqldb.Config{
		Host:     db_host,
		Port:     db_port,
		DbName:   db_name,
		UserName: db_username,
		Password: db_password,
	})
}

//example 3 cache instance
func initCache() error {
	//default instance
	err := cache.Init()
	if err != nil {
		return err
	}

	// cache1 instance
	err = cache.Init_("cache1")
	if err != nil {
		return err
	}

	// cache2 instance
	err = cache.Init_("cache2")
	if err != nil {
		return err
	}

	return nil
}

//todo: ---
func initComponent() {

	err := iniHub()
	if err != nil {
		basic.Logger.Fatalln(err)
	}

	err = initEchoServer()
	if err != nil {
		basic.Logger.Fatalln(err)
	}

	//err = initElasticSearch()
	//if err != nil {
	//	basic.Logger.Fatalln(err)
	//}
	//
	// err = initRedis()
	// if err != nil {
	// 	basic.Logger.Fatalln(err)
	// }
	//
	//err = initSpr()
	//if err != nil {
	//	basic.Logger.Fatalln(err)
	//}
	//
	//err = initDB()
	//if err != nil {
	//	basic.Logger.Fatalln(err)
	//}

	err = initCache()
	if err != nil {
		basic.Logger.Fatalln(err)
	}
}
