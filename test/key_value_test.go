package test

import (
	"log"
	"testing"

	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/plugin/redis_plugin"
	"github.com/coreservice-io/cli-template/plugin/reference_plugin"
	"github.com/coreservice-io/cli-template/src/examples/data_redis"
)

func initialize_kv() {
	basic.InitLogger()

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

func Test_peer(t *testing.T) {
	initialize_kv()
	//
	p := &data_redis.PeerInfo{
		Tag:      "abcd",
		Location: "USA",
		IP:       "127.0.0.1",
	}
	tag := "abcd"

	err := data_redis.SetPeer(p, tag)
	if err != nil {
		log.Fatalln("SetPeer err", err, "tag", tag)
	}

	pp, err := data_redis.GetPeer(tag, false)
	log.Println(pp, err)

	data_redis.DeletePeer(tag)

	pp, err = data_redis.GetPeer(tag, false)
	log.Println(pp, err)
}
