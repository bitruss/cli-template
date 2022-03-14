package lrcache

import (
	"context"
	"log"
	"testing"

	"github.com/coreservice-io/CliAppTemplate/plugin/cache"
	"github.com/coreservice-io/CliAppTemplate/plugin/redisClient"
)

func init() {
	//db
	//err := sqldb.Init(sqldb.Config{
	//	Host:     "127.0.0.1",
	//	Port:     3306,
	//	DbName:   "testdb",
	//	UserName: "root",
	//	Password: "123456",
	//})
	//if err != nil {
	//	log.Fatalln(err)
	//}

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
		log.Fatalln(err)
	}

	//cache
	err = cache.Init()
	if err != nil {
		log.Fatalln(err)
	}
}

type ExampleStruct struct {
	ID      int
	Status  string
	Name    string
	Email   string
	Updated int64 `gorm:"autoUpdateTime"`
	Created int64 `gorm:"autoCreateTime"`
}

func Test_(t *testing.T) {
	LRC_Set(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:int", 789, 300)
	r := LRC_Get(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:int")
	log.Println(r.(int))

	LRC_Set(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:float", 789.111, 300)
	r = LRC_Get(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:float")
	log.Println(r.(float64))

	LRC_Set(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:string", "string test", 300)
	r = LRC_Get(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:string")
	log.Println(r.(string))

	LRC_Set(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:bool", true, 300)
	r = LRC_Get(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:bool")
	log.Println(r.(bool))

	LRC_Set(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:array", []int{1, 2, 3, 4, 5, 6}, 300)
	r = LRC_Get(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:array")
	log.Println(r.([]int))

	data := ExampleStruct{
		ID:     1,
		Status: "running",
		Name:   "bruce",
		Email:  "bruce@mail.com",
	}
	LRC_Set(context.Background(), redisClient.GetInstance(), cache.GetInstance(), true, "test:json", data, 300)
	r = LRC_Get(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:json")
	log.Println(r.(ExampleStruct))

	dataArray := []ExampleStruct{{
		ID:     1,
		Status: "running",
		Name:   "bruce",
		Email:  "bruce@mail.com",
	}, {
		ID:     2,
		Status: "running2",
		Name:   "bruce2",
		Email:  "bruce2@mail.com",
	}}
	LRC_Set(context.Background(), redisClient.GetInstance(), cache.GetInstance(), true, "test:json", dataArray, 300)
	r = LRC_Get(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:json")
	log.Println(r.([]ExampleStruct))
}
