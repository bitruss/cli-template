package test

import (
	"context"
	"log"
	"testing"

	"github.com/coreservice-io/CliAppTemplate/plugin/redisClient"
	"github.com/coreservice-io/CliAppTemplate/plugin/reference"
	"github.com/coreservice-io/CliAppTemplate/tools/smartCache"
)

func init() {
	//redis
	err := redisClient.Init(redisClient.Config{
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
	err = reference.Init()
	if err != nil {
		log.Fatalln("reference init err", err)
	}
}

type person struct {
	Name string
	Age  int
}

func Test_BuildInType(t *testing.T) {
	key := "test:111"
	v := 7
	err := smartCache.RR_Set(context.Background(), redisClient.GetInstance().ClusterClient, reference.GetInstance(), false, key, &v, 300)
	if err != nil {
		log.Println("RR_Set error", err)
	}
	r := smartCache.Ref_Get(reference.GetInstance(), key)
	log.Println(r.(*int))
	var rInt int
	smartCache.Redis_Get(context.Background(), redisClient.GetInstance().ClusterClient, false, key, &rInt)
	log.Println(rInt)
}

func Test_Struct(t *testing.T) {
	key := "test:111"
	v := &person{
		Name: "Jack",
		Age:  10,
	}
	err := smartCache.RR_Set(context.Background(), redisClient.GetInstance().ClusterClient, reference.GetInstance(), true, key, v, 300)
	if err != nil {
		log.Println("RR_Set error", err)
	}
	r := smartCache.Ref_Get(reference.GetInstance(), key)
	log.Println(r.(*person))
	var p person
	smartCache.Redis_Get(context.Background(), redisClient.GetInstance().ClusterClient, true, key, &p)
	log.Println(p)
}
