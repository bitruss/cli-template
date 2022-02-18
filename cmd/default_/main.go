package default_

import (
	"errors"
	"time"

	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
	"github.com/universe-30/CliAppTemplate/basic"
	"github.com/universe-30/CliAppTemplate/configuration"
	"github.com/universe-30/CliAppTemplate/plugin/cache"
	"github.com/universe-30/CliAppTemplate/plugin/echoServer"
	"github.com/universe-30/CliAppTemplate/plugin/es"
	"github.com/universe-30/CliAppTemplate/plugin/redisClient"
	"github.com/universe-30/CliAppTemplate/plugin/sprMgr"
	"github.com/universe-30/CliAppTemplate/plugin/sqldb"
	"github.com/universe-30/CliAppTemplate/tools"
	"github.com/universe-30/UJob"
	"github.com/urfave/cli/v2"

	"github.com/universe-30/USafeGo"
)

func initEchoServer() error {
	http_port, err := configuration.Config.GetInt("http_port", 8080)
	if err != nil {
		return errors.New("http_port [int] in config error," + err.Error())
	}

	http_static_rel_folder, err := configuration.Config.GetString("http_static_rel_folder", "")
	if err != nil {
		return errors.New("http_static_rel_folder [string] in config error," + err.Error())
	}

	return echoServer.Init("", echoServer.Config{Port: http_port, StaticFolder: http_static_rel_folder})
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

	return es.Init("", es.Config{
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

	return redisClient.Init("", redisClient.Config{
		Address:  redis_addr,
		UserName: redis_username,
		Password: redis_password,
		Port:     redis_port,
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

	return sprMgr.Init("", sprMgr.Config{
		Address:  redis_addr,
		UserName: redis_username,
		Password: redis_password,
		Port:     redis_port,
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

	return sqldb.Init("", sqldb.Config{
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
	err := cache.Init("")
	if err != nil {
		return err
	}

	// cache1 instance
	err = cache.Init("cache1")
	if err != nil {
		return err
	}

	// cache2 instance
	err = cache.Init("cache2")
	if err != nil {
		return err
	}

	return nil
}

//todo: ---
func InitComponent() {
	err := initEchoServer()
	if err != nil {
		basic.Logger.Fatalln(err)
	}

	//err = initElasticSearch()
	//if err != nil {
	//	basic.Logger.Fatalln(err)
	//}
	//
	//err = initRedis()
	//if err != nil {
	//	basic.Logger.Fatalln(err)
	//}
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

func StartDefault(clictx *cli.Context) {

	color.Green(basic.Logo)
	basic.Logger.Infoln("hello world , this default app")

	//example
	InitComponent()
	//defer func() {
	//	//global.ReleaseResources()
	//}()

	//example get complex config
	provide_folder, err := configuration.Config.GetProvideFolders()
	if err != nil {
		basic.Logger.Errorln(err)
	}
	for _, v := range provide_folder {
		basic.Logger.Debugln("path:", v.AbsPath, "size:", v.SizeGB)
	}

	//cache example
	cache.GetInstance("cache1").Set("foo1", "bar1", 10)
	v, _, exist := cache.GetInstance("cache1").Get("foo1")
	if exist {
		basic.Logger.Debugln(v.(string))
	}

	cache.GetInstance("cache2").Set("foo2", "bar2", 10)
	v, _, exist = cache.GetInstance("cache2").Get("foo2")
	if exist {
		basic.Logger.Debugln(v.(string))
	}

	//redis example
	//if redisClient.GetDefaultInstance() != nil {
	//	redisClient.GetDefaultInstance().Set(context.Background(), "redis-foo", "redis-bar", 10*time.Second)
	//	str, err := redisClient.GetDefaultInstance().Get(context.Background(), "redis-foo").Result()
	//	if err != nil && err != goredis.Nil {
	//		basic.Logger.Errorln(err)
	//	}
	//	basic.Logger.Debugln(str)
	//}

	//schedule job
	count := 0
	job := UJob.Start(
		//job process
		"exampleJob",
		func() {
			count++
			basic.Logger.Debugln("Schedule Job running,count", count)
		},
		//onPanic callback
		tools.PanicHandler,
		2,
		// job type
		// UJob.TYPE_PANIC_REDO  auto restart if panic
		// UJob.TYPE_PANIC_RETURN  stop if panic
		UJob.TYPE_PANIC_REDO,
		// check continue callback, the job will stop running if return false
		// the job will keep running if this callback is nil
		func(job *UJob.Job) bool {
			return true
		},
		// onFinish callback
		func(inst *UJob.Job) {
			basic.Logger.Debugln("finish", "cycle", inst.Cycles)
		},
	)

	//safeGo
	USafeGo.Go(
		//process
		func(args ...interface{}) {
			basic.Logger.Debugln("example of USafeGo")
			time.Sleep(10 * time.Second)
			job.SetToCancel()
		},
		//onPanic callback
		tools.PanicHandler)

	for i := 0; i < 10; i++ {
		basic.Logger.Infoln("running")
		time.Sleep(1 * time.Second)
	}

	//httpServer example
	httpServer := echoServer.GetDefaultInstance()
	httpServer.GET("/test", func(context echo.Context) error {
		return context.String(200, "test success")
	})
	httpServer.Start()
}
