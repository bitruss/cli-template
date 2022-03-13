package default_

import (
	"context"
	"fmt"
	"time"

	"github.com/coreservice-io/CliAppTemplate/basic"
	"github.com/coreservice-io/CliAppTemplate/configuration"
	"github.com/coreservice-io/CliAppTemplate/plugin/cache"
	"github.com/coreservice-io/CliAppTemplate/plugin/echoServer"
	"github.com/coreservice-io/CliAppTemplate/plugin/hub"
	"github.com/coreservice-io/CliAppTemplate/plugin/redisClient"
	"github.com/coreservice-io/CliAppTemplate/tools/errors"
	uhub "github.com/coreservice-io/UHub"
	"github.com/coreservice-io/UJob"
	"github.com/coreservice-io/USafeGo"
	"github.com/fatih/color"
	goredis "github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/urfave/cli/v2"
)

func StartDefault(clictx *cli.Context) {

	color.Green(basic.Logo)
	basic.Logger.Infoln("hello world , this default app")

	//example
	initComponent()

	//defer func() {
	//	//global.ReleaseResources()
	//}()

	hub_testrun()
	complexconfig_testrun()
	cache_testrun()
	redis_testrun()
	job_safego_testrun()
	httpserver_testrun()
}

//hub example
const testKind uhub.Kind = 1

type testEvent string

func (e testEvent) Kind() uhub.Kind {
	return testKind
}
func hub_testrun() {
	hub.GetInstance().Subscribe(testKind, func(e uhub.Event) {
		fmt.Println("hub callback")
		fmt.Println(string(e.(testEvent)))
	})
	hub.GetInstance().Publish(testEvent("hub message"))
}

//example get complex config
func complexconfig_testrun() {
	provide_folder, err := configuration.Config.GetProvideFolders()
	if err != nil {
		basic.Logger.Errorln(err)
	}
	for _, v := range provide_folder {
		basic.Logger.Debugln("path:", v.AbsPath, "size:", v.SizeGB)
	}
}

//cache example
func cache_testrun() {
	cache.GetInstance_("cache1").Set("foo1", "bar1", 10)
	v, _, exist := cache.GetInstance_("cache1").Get("foo1")
	if exist {
		basic.Logger.Debugln(v.(string))
	}

	cache.GetInstance_("cache2").Set("foo2", "bar2", 10)
	v, _, exist = cache.GetInstance_("cache2").Get("foo2")
	if exist {
		basic.Logger.Debugln(v.(string))
	}
}

//redis example
func redis_testrun() {
	if redisClient.GetInstance() != nil {
		key := redisClient.GetInstance().GenKey("foo")
		redisClient.GetInstance().Set(context.Background(), key, "redis-bar", 100*time.Second)
		str, err := redisClient.GetInstance().Get(context.Background(), "redis-foo").Result()
		if err != nil && err != goredis.Nil {
			basic.Logger.Errorln(err)
		}
		basic.Logger.Debugln(str)
	}
}

//job and safego example
func job_safego_testrun() {
	count := 0
	job := UJob.Start(
		//job process
		"exampleJob",
		func() {
			count++
			basic.Logger.Debugln("Schedule Job running,count", count)
		},
		//onPanic callback
		errors.PanicHandler,
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
		errors.PanicHandler)

	for i := 0; i < 10; i++ {
		basic.Logger.Infoln("running")
		time.Sleep(1 * time.Second)
	}
}

//httpServer example
func httpserver_testrun() {
	httpServer := echoServer.GetInstance()
	//health api
	httpServer.GET("/api/health", func(context echo.Context) error {
		return echoServer.SuccessResp(context, 1, time.Now().Unix(), "")
	})
	httpServer.Start()
}
