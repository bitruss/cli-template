package examples

// import (
// 	"time"

// 	"github.com/coreservice-io/cli-template/basic"
// 	"github.com/coreservice-io/cli-template/tools/errors"

// 	"github.com/coreservice-io/job"
// 	"github.com/coreservice-io/safe_go"
// )

// //job and safego example
// func Job_Safeo_run() {
// 	count := 0
// 	job := job.Start(
// 		//job process
// 		"exampleJob",
// 		func() {
// 			count++
// 			basic.Logger.Debugln("Schedule Job running,count", count)
// 		},
// 		//onPanic callback
// 		errors.PanicHandler,
// 		2,
// 		// job type
// 		// job.TYPE_PANIC_REDO  auto restart if panic
// 		// job.TYPE_PANIC_RETURN  stop if panic
// 		job.TYPE_PANIC_REDO,
// 		// check continue callback, the job will stop running if return false
// 		// the job will keep running if this callback is nil
// 		func(job *job.Job) bool {
// 			return true
// 		},
// 		// onFinish callback
// 		func(inst *job.Job) {
// 			basic.Logger.Debugln("finish", "cycle", inst.Cycles)
// 		},
// 	)

// 	//safeGo
// 	safe_go.Go(
// 		//process
// 		func(args ...interface{}) {
// 			basic.Logger.Debugln("example of safe_go")
// 			time.Sleep(10 * time.Second)
// 			job.SetToCancel()
// 		},
// 		//onPanic callback
// 		errors.PanicHandler)

// 	for i := 0; i < 10; i++ {
// 		basic.Logger.Debugln("running")
// 		time.Sleep(1 * time.Second)
// 	}
// }
