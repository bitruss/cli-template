package test

import (
	"log"
	"testing"

	"github.com/coreservice-io/CliAppTemplate/basic"
	"github.com/coreservice-io/CliAppTemplate/plugin/redisClient"
	"github.com/coreservice-io/CliAppTemplate/plugin/reference"
	"github.com/coreservice-io/CliAppTemplate/src/examples/dataInRedis"
)

func initialize_kv() {
	basic.InitLogger()

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

func Test_peer(t *testing.T) {
	initialize_kv()
	//
	p := &dataInRedis.PeerInfo{
		Tag:      "abcd",
		Location: "USA",
		IP:       "127.0.0.1",
	}
	tag := "abcd"

	err := dataInRedis.SetPeer(p, tag)
	if err != nil {
		log.Fatalln("SetPeer err", err, "tag", tag)
	}

	pp, err := dataInRedis.GetPeer(tag, false)
	log.Println(pp, err)

	dataInRedis.DeletePeer(tag)

	pp, err = dataInRedis.GetPeer(tag, false)
	log.Println(pp, err)
}
