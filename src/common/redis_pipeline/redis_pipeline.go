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
const cmd_channel_limit = 20000

var last_exec_time_unixmilli = time.Now().UTC().UnixMilli()

type PipelineCmd struct {
	Ctx       context.Context
	Operation string
	Key       string
	Args      []interface{}
}

var cmdListChannel = make(chan *PipelineCmd, cmd_channel_limit)

func ScheduleRedisPipelineExec() {
	const jobName = "ScheduleRedisPipelineExec"

	for i := 0; i < exec_thread_count; i++ {
		job.Start(
			//job process
			jobName,
			func() {
				for {
					if len(cmdListChannel) < 100 && time.Now().UTC().UnixMilli()-last_exec_time_unixmilli < exec_interval_limit_millisec {
						time.Sleep(250 * time.Millisecond)
						//basic.Logger.Debugln("sleep")

						continue
					}
					exec()
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
}

func exec() {
	//basic.Logger.Debugln("exec redis pipeline")
	//basic.Logger.Debugln("channel length", len(cmdListChannel))

	last_exec_time_unixmilli = time.Now().UTC().UnixMilli()

	pl := redis_plugin.GetInstance().Pipeline()
	cmds := []*PipelineCmd{}

outLoop:
	for i := 0; i < exec_count_limit; i++ {
		select {
		case cmd := <-cmdListChannel:
			switch cmd.Operation {
			case operation_Set:
				pl.Set(cmd.Ctx, cmd.Key, cmd.Args[0], cmd.Args[1].(time.Duration))
				cmds = append(cmds, cmd)

			case operation_ZAdd:
				z := []*redis.Z{}
				for _, v := range cmd.Args {
					z = append(z, v.(*redis.Z))
				}
				pl.ZAdd(cmd.Ctx, cmd.Key, z...)
				cmds = append(cmds, cmd)

			case operation_ZAddNX:
				z := []*redis.Z{}
				for _, v := range cmd.Args {
					z = append(z, v.(*redis.Z))
				}
				pl.ZAddNX(cmd.Ctx, cmd.Key, z...)
				cmds = append(cmds, cmd)

			case operation_HSet:
				pl.HSet(cmd.Ctx, cmd.Key, cmd.Args...)
				cmds = append(cmds, cmd)

			case operation_Expire:
				pl.Expire(cmd.Ctx, cmd.Key, cmd.Args[0].(time.Duration))
				cmds = append(cmds, cmd)

			case operation_ZRemRangeByScore:
				pl.ZRemRangeByScore(cmd.Ctx, cmd.Key, cmd.Args[0].(string), cmd.Args[1].(string))
				cmds = append(cmds, cmd)

			default:
				basic.Logger.Errorln("unsupported cmd:", cmd.Operation)
			}

		default:
			break outLoop
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
}
