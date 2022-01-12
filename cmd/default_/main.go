package default_

import (
	"time"

	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
	"github.com/universe-30/CliAppTemplate/basic"
	"github.com/universe-30/CliAppTemplate/components/cache"
	"github.com/universe-30/CliAppTemplate/components/echoServer"
	"github.com/universe-30/CliAppTemplate/configuration"
	"github.com/universe-30/CliAppTemplate/tools"
	"github.com/universe-30/UJob"
	"github.com/urfave/cli/v2"

	"github.com/universe-30/USafeGo"
)

//todo: ---
func InitComponent() {
	cache.Init()
	//echoServer.Init()
	//redisClient.Init()
	//sprMgr.Init()
	//sqldb.Init()
	//es.Init()
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
	// cache.GetSingleInstance().Set("foo", "bar", 10)
	// v, _, exist := cache.GetSingleInstance().Get("foo")
	// if exist {
	// 	basic.Logger.Debugln(v.(string))
	// }

	//redis example
	// if redisClient.GetSingleInstance() != nil {
	// 	redisClient.GetSingleInstance().Set(context.Background(), "redis-foo", "redis-bar", 10*time.Second)
	// 	str, err := redisClient.GetSingleInstance().Get(context.Background(), "redis-foo").Result()
	// 	if err != nil && err != goredis.Nil {
	// 		basic.Logger.Errorln(err)
	// 	}
	// 	basic.Logger.Debugln(str)
	// }

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
	httpServer := echoServer.GetSingleInstance()
	httpServer.GET("/test", func(context echo.Context) error {
		return context.String(200, "test success")
	})
	httpServer.UseJsoniter()
	httpServer.SetPanicHandler(tools.PanicHandler)
	httpServer.Start()
}
