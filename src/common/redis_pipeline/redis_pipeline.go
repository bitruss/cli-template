package redis_pipeline

import (
	"context"
	"time"

	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/plugin/redis_plugin"
	"github.com/coreservice-io/job"

	"github.com/go-redis/redis/v8"
)

const exec_count_limit = 100
const exec_interval_limit_millisec = 2500
const exec_thread_count = 8

var last_exec_time_unixmilli = time.Now().UTC().UnixMilli()

type PipelineCmd struct {
	Ctx       context.Context
	Operation string
	Key       string
	Args      []interface{}

	StatusCmd_callback func(statusCmd *redis.StatusCmd)
	IntCmd_callback    func(intCmd *redis.IntCmd)
	BoolCmd_callback   func(boolCmd *redis.BoolCmd)

	ResultCmd interface{}
}

var cmdListChannel = make(chan *PipelineCmd, 2000)

func ScheduleRedisPipelineExec() {
	const jobName = "ScheduleRedisPipelineExec"

	job.Start(
		//job process
		jobName,
		func() {
			//limit goroutinue count
			ch := make(chan bool, exec_thread_count)
			for {
				ch <- true
				go func() {
					if len(cmdListChannel) < 100 && time.Now().UTC().UnixMilli()-last_exec_time_unixmilli < exec_interval_limit_millisec {
						time.Sleep(250 * time.Millisecond)
						//basic.Logger.Debugln("sleep")
						<-ch
						return
					}

					exec()
					<-ch
				}()
			}
		},
		//onPanic callback
		func(panic_err interface{}) {
			basic.Logger.Errorln(panic_err)
		},
		1,
		// job type
		// UJob.TYPE_PANIC_REDO  auto restart if panic
		// UJob.TYPE_PANIC_RETURN  stop if panic
		job.TYPE_PANIC_REDO,
		// check continue callback, the job will stop running if return false
		// the job will keep running if this callback is nil
		nil,
		// onFinish callback
		nil,
	)
}

func exec() {
	//basic.Logger.Debugln("exec redis pipeline")
	//basic.Logger.Debugln("channel length", len(cmdListChannel))

	last_exec_time_unixmilli = time.Now().UTC().UnixMilli()

	pl := redis_plugin.GetInstance().Pipeline()
	cmds := []*PipelineCmd{}
	for i := 0; i < exec_count_limit; i++ {
		select {
		case cmd := <-cmdListChannel:
			switch cmd.Operation {
			case operation_Set:
				cmd.ResultCmd = pl.Set(cmd.Ctx, cmd.Key, cmd.Args[0], cmd.Args[1].(time.Duration))
				cmds = append(cmds, cmd)

			case operation_ZAdd:
				z := []*redis.Z{}
				for _, v := range cmd.Args {
					z = append(z, v.(*redis.Z))
				}
				cmd.ResultCmd = pl.ZAdd(cmd.Ctx, cmd.Key, z...)
				cmds = append(cmds, cmd)

			case operation_ZAddNX:
				z := []*redis.Z{}
				for _, v := range cmd.Args {
					z = append(z, v.(*redis.Z))
				}
				cmd.ResultCmd = pl.ZAddNX(cmd.Ctx, cmd.Key, z...)
				cmds = append(cmds, cmd)

			case operation_HSet:
				cmd.ResultCmd = pl.HSet(cmd.Ctx, cmd.Key, cmd.Args...)
				cmds = append(cmds, cmd)

			case operation_Expire:
				cmd.ResultCmd = pl.Expire(cmd.Ctx, cmd.Key, cmd.Args[0].(time.Duration))
				cmds = append(cmds, cmd)

			case operation_ZRemRangeByScore:
				cmd.ResultCmd = pl.ZRemRangeByScore(cmd.Ctx, cmd.Key, cmd.Args[0].(string), cmd.Args[1].(string))
				cmds = append(cmds, cmd)

			default:
				basic.Logger.Errorln("unsupported cmd:", cmd.Operation)
			}

		//todo default is not working
		default:
			break
		}
	}

	if pl.Len() == 0 {
		return
	}

	_, err := pl.Exec(context.Background())
	if err != nil {
		basic.Logger.Errorln("exec pipeline error:", err)
		return
	}
	for _, cmd := range cmds {
		switch cmd.Operation {
		case operation_Set:
			if cmd.StatusCmd_callback != nil {
				cmd.StatusCmd_callback(cmd.ResultCmd.(*redis.StatusCmd))
			}
		case operation_ZAdd, operation_HSet, operation_ZRemRangeByScore:
			if cmd.IntCmd_callback != nil {
				cmd.IntCmd_callback(cmd.ResultCmd.(*redis.IntCmd))
			}
		case operation_Expire:
			if cmd.BoolCmd_callback != nil {
				cmd.BoolCmd_callback(cmd.ResultCmd.(*redis.BoolCmd))
			}

		default:

		}
	}

}
