package examples

// import (
// 	"context"
// 	"time"

// 	"github.com/coreservice-io/cli-template/basic"
// 	"github.com/coreservice-io/cli-template/plugin/redis_plugin"
// 	goredis "github.com/go-redis/redis/v8"
// )

// //redis example
// func Redis_run() {
// 	if redis_plugin.GetInstance() != nil {
// 		key := redis_plugin.GetInstance().GenKey("foo")
// 		redis_plugin.GetInstance().Set(context.Background(), key, "redis-bar", 100*time.Second)
// 		str, err := redis_plugin.GetInstance().Get(context.Background(), "redis-foo").Result()
// 		if err != nil && err != goredis.Nil {
// 			basic.Logger.Errorln(err)
// 		}
// 		basic.Logger.Debugln(str)
// 	}
// }
