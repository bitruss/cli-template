package redis_pipeline

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	operation_Set              = "Set"
	operation_ZAdd             = "ZAdd"
	operation_ZAddNX           = "ZAddNX"
	operation_HSet             = "HSet"
	operation_Expire           = "Expire"
	operation_ZRemRangeByScore = "ZRemRangeByScore"
)

func Set(ctx context.Context, callback func(statusCmd *redis.StatusCmd), key string, value interface{}, expiration time.Duration) {
	redisCmd := &PipelineCmd{
		Ctx:                ctx,
		Operation:          operation_Set,
		Key:                key,
		Args:               []interface{}{value, expiration},
		StatusCmd_callback: callback,
	}
	cmdListChannel <- redisCmd
}

func ZAdd(ctx context.Context, callback func(intCmd *redis.IntCmd), key string, members ...*redis.Z) {
	redisCmd := &PipelineCmd{
		Ctx:             ctx,
		Operation:       operation_ZAdd,
		Key:             key,
		Args:            []interface{}{},
		IntCmd_callback: callback,
	}
	for _, v := range members {
		redisCmd.Args = append(redisCmd.Args, v)
	}
	cmdListChannel <- redisCmd
}

func ZAddNX(ctx context.Context, callback func(intCmd *redis.IntCmd), key string, members ...*redis.Z) {
	redisCmd := &PipelineCmd{
		Ctx:             ctx,
		Operation:       operation_ZAddNX,
		Key:             key,
		Args:            []interface{}{},
		IntCmd_callback: callback,
	}
	for _, v := range members {
		redisCmd.Args = append(redisCmd.Args, v)
	}
	cmdListChannel <- redisCmd
}

func HSet(ctx context.Context, callback func(intCmd *redis.IntCmd), key string, values ...interface{}) {
	redisCmd := &PipelineCmd{
		Ctx:             ctx,
		Operation:       operation_HSet,
		Key:             key,
		Args:            []interface{}{},
		IntCmd_callback: callback,
	}
	redisCmd.Args = append(redisCmd.Args, values...)
	cmdListChannel <- redisCmd
}

func Expire(ctx context.Context, callback func(boolCmd *redis.BoolCmd), key string, expiration time.Duration) {
	redisCmd := &PipelineCmd{
		Ctx:              ctx,
		Operation:        operation_Expire,
		Key:              key,
		Args:             []interface{}{expiration},
		BoolCmd_callback: callback,
	}
	cmdListChannel <- redisCmd
}

func ZRemRangeByScore(ctx context.Context, callback func(intCmd *redis.IntCmd), key, min, max string) {
	redisCmd := &PipelineCmd{
		Ctx:             ctx,
		Operation:       operation_ZRemRangeByScore,
		Key:             key,
		Args:            []interface{}{min, max},
		IntCmd_callback: callback,
	}
	cmdListChannel <- redisCmd
}
