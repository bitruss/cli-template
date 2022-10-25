package test

import (
	"context"
	"testing"

	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/component"
	"github.com/coreservice-io/cli-template/config"
	"github.com/coreservice-io/cli-template/plugin/redis_plugin"
	"github.com/coreservice-io/cli-template/plugin/reference_plugin"
	"github.com/coreservice-io/cli-template/src/common/smart_cache"
)

func initialize_smc() {

	config.ConfigBasic("test")

	config_toml := config.Get_config().Toml_config

	/////////////////////////
	if err := component.InitReference(); err != nil {
		basic.Logger.Fatalln(err)
	}

	/////////////////////////
	if err := component.InitRedis(config_toml); err != nil {
		basic.Logger.Fatalln(err)
	}

}

type person struct {
	Name string
	Age  int
}

func Test_BuildInType(t *testing.T) {
	initialize_smc()
	key := "test:111"
	v := 7
	err := smart_cache.RR_Set(context.Background(), redis_plugin.GetInstance().ClusterClient, reference_plugin.GetInstance(), false, key, &v, 300)
	if err != nil {
		basic.Logger.Infoln("RR_Set error", err)
	}
	r := smart_cache.Ref_Get(reference_plugin.GetInstance(), key)
	if r != nil {
		basic.Logger.Infoln(r.(*int))
	}
	var rInt int
	smart_cache.Redis_Get(context.Background(), redis_plugin.GetInstance().ClusterClient, false, key, &rInt)
	basic.Logger.Infoln(rInt)
}

func Test_Struct(t *testing.T) {
	initialize_smc()
	key := "test:111"
	v := &person{
		Name: "Jack",
		Age:  10,
	}
	err := smart_cache.RR_Set(context.Background(), redis_plugin.GetInstance().ClusterClient, reference_plugin.GetInstance(), true, key, v, 300)
	if err != nil {
		basic.Logger.Infoln("RR_Set error", err)
	}
	r := smart_cache.Ref_Get(reference_plugin.GetInstance(), key)
	if r != nil {
		basic.Logger.Infoln(r.(*person))
	}
	var p person
	smart_cache.Redis_Get(context.Background(), redis_plugin.GetInstance().ClusterClient, true, key, &p)
	basic.Logger.Infoln(p)
}
