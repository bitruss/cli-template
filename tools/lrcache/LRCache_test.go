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
	var r1 int
	LRC_Set(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:int", 789, 300)
	LRC_Get(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:int", &r1)
	log.Println(r1)

	var r2 int8
	LRC_Set(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:int", 127, 300)
	LRC_Get(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:int", &r2)
	log.Println(r2)

	var r3 int16
	LRC_Set(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:int", -98, 300)
	LRC_Get(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:int", &r3)
	log.Println(r3)

	var r4 int32
	LRC_Set(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:int", -98, 300)
	LRC_Get(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:int", &r4)
	log.Println(r4)

	var r5 int64
	LRC_Set(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:int", -98, 300)
	LRC_Get(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:int", &r5)
	log.Println(r5)

	var r6 uint8
	LRC_Set(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:int", 250, 300)
	LRC_Get(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:int", &r6)
	log.Println(r6)

	var r7 uint16
	LRC_Set(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:int", 345346, 300)
	LRC_Get(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:int", &r7)
	log.Println(r7)

	var r8 uint32
	LRC_Set(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:int", 345346, 300)
	LRC_Get(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:int", &r8)
	log.Println(r8)

	var r9 uint64
	LRC_Set(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:int", 345346, 300)
	LRC_Get(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:int", &r9)
	log.Println(r9)

	var r10 uint
	LRC_Set(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:int", 345346, 300)
	LRC_Get(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:int", &r10)
	log.Println(r10)

	var r11 float32
	LRC_Set(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:int", 345.346, 300)
	LRC_Get(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:int", &r11)
	log.Println(r11)

	var r12 float64
	LRC_Set(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:int", 345.346, 300)
	LRC_Get(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:int", &r12)
	log.Println(r12)

	var r13 bool
	LRC_Set(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:int", true, 300)
	LRC_Get(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:int", &r13)
	log.Println(r13)

	//LRC_Set(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:float", 789.111, 300)
	//r = LRC_Get(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:float")
	//log.Println(r.(float64))
	//
	//LRC_Set(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:string", "string test", 300)
	//r = LRC_Get(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:string")
	//log.Println(r.(string))
	//
	//LRC_Set(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:bool", true, 300)
	//r = LRC_Get(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:bool")
	//log.Println(r.(bool))
	//
	//LRC_Set(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:array", []int{1, 2, 3, 4, 5, 6}, 300)
	//r = LRC_Get(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:array")
	//log.Println(r.([]int))
	//
	//data := ExampleStruct{
	//	ID:     1,
	//	Status: "running",
	//	Name:   "bruce",
	//	Email:  "bruce@mail.com",
	//}
	//LRC_Set(context.Background(), redisClient.GetInstance(), cache.GetInstance(), true, "test:json", data, 300)
	//r = LRC_Get(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:json")
	//log.Println(r.(ExampleStruct))
	//
	//dataArray := []ExampleStruct{{
	//	ID:     1,
	//	Status: "running",
	//	Name:   "bruce",
	//	Email:  "bruce@mail.com",
	//}, {
	//	ID:     2,
	//	Status: "running2",
	//	Name:   "bruce2",
	//	Email:  "bruce2@mail.com",
	//}}
	//LRC_Set(context.Background(), redisClient.GetInstance(), cache.GetInstance(), true, "test:json", dataArray, 300)
	//r = LRC_Get(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, "test:json")
	//log.Println(r.([]ExampleStruct))
}
