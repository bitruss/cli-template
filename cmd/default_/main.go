package default_

import (
	"context"
	"time"

	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
	"github.com/universe-30/CliAppTemplate/basic"
	"github.com/universe-30/CliAppTemplate/cmd/default_/global"
	"github.com/universe-30/CliAppTemplate/tools"
	"github.com/universe-30/UJob"
	"github.com/urfave/cli/v2"

	goredis "github.com/go-redis/redis/v8"
	"github.com/universe-30/USafeGo"
)

func StartDefault(clictx *cli.Context) {

	color.Green(basic.Logo)
	basic.Logger.Infoln("hello world , this default app")

	//example
	global.IniResources()
	defer func() {
		global.ReleaseResources()
	}()

	//cache example
	global.Cache.Set("foo", "bar", 10)
	v, _, exist := global.Cache.Get("foo")
	if exist {
		basic.Logger.Debugln(v.(string))
	}

	//redis example
	global.Redis.Set(context.Background(), "redis-foo", "redis-bar", 10*time.Second)
	str, err := global.Redis.Get(context.Background(), "redis-foo").Result()
	if err != nil && err != goredis.Nil {
		basic.Logger.Errorln(err)
	}
	basic.Logger.Debugln(str)

	//schedule job
	count := 0
	job := UJob.Start(
		//job process
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
	global.HttpServer.GET("/test", func(context echo.Context) error {
		return context.String(200, "test success")
	})
	global.HttpServer.UseJsoniter()
	global.HttpServer.SetPanicHandler(tools.PanicHandler)
	global.HttpServer.Start()
}
