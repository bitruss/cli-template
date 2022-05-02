package test

import (
	"context"
	"log"
	"testing"

	"github.com/coreservice-io/CliAppTemplate/plugin/redis_plugin"
	"github.com/coreservice-io/CliAppTemplate/plugin/reference_plugin"
	"github.com/coreservice-io/CliAppTemplate/tools/smart_cache"
)

func initialize_smc() {
	//redis
	err := redis_plugin.Init(redis_plugin.Config{
		Address:   "127.0.0.1",
		UserName:  "",
		Password:  "",
		Port:      6379,
		KeyPrefix: "userTest:",
		UseTLS:    false,
	})
	if err != nil {
		log.Fatalln("redis init err", err)
	}

	//reference
	err = reference_plugin.Init()
	if err != nil {
		log.Fatalln("reference init err", err)
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
		log.Println("RR_Set error", err)
	}
	r := smart_cache.Ref_Get(reference_plugin.GetInstance(), key)
	log.Println(r.(*int))
	var rInt int
	smart_cache.Redis_Get(context.Background(), redis_plugin.GetInstance().ClusterClient, false, key, &rInt)
	log.Println(rInt)
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
		log.Println("RR_Set error", err)
	}
	r := smart_cache.Ref_Get(reference_plugin.GetInstance(), key)
	log.Println(r.(*person))
	var p person
	smart_cache.Redis_Get(context.Background(), redis_plugin.GetInstance().ClusterClient, true, key, &p)
	log.Println(p)
}
