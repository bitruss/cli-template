package smartCache

import (
	"context"
	"log"
	"testing"

	"github.com/coreservice-io/CliAppTemplate/plugin/redisClient"
	"github.com/coreservice-io/CliAppTemplate/plugin/reference"
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
	err := RR_Set(context.Background(), redisClient.GetInstance(), reference.GetInstance(), false, key, &v, 300)
	if err != nil {
		log.Println("RR_Set error", err)
	}
	r := Ref_Get(reference.GetInstance(), key)
	log.Println(r.(*int))
	var rInt int
	Redis_Get(context.Background(), redisClient.GetInstance(), false, key, &rInt)
	log.Println(rInt)
}

func Test_Struct(t *testing.T) {
	key := "test:111"
	v := &person{
		Name: "Jack",
		Age:  10,
	}
	err := RR_Set(context.Background(), redisClient.GetInstance(), reference.GetInstance(), true, key, v, 300)
	if err != nil {
		log.Println("RR_Set error", err)
	}
	r := Ref_Get(reference.GetInstance(), key)
	log.Println(r.(*person))
	var p person
	Redis_Get(context.Background(), redisClient.GetInstance(), true, key, &p)
	log.Println(p)
}
