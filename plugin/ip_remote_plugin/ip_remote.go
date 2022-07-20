package ip_remote_plugin

import (
	"fmt"

	"github.com/coreservice-io/ip_geo/ipdata"
	"github.com/coreservice-io/log"
	"github.com/coreservice-io/reference"
)

type IpRemote struct {
	Client *ipdata.IpData
}

var instanceMap = map[string]*IpRemote{}

func GetInstance() *IpRemote {
	return instanceMap["default"]
}

func GetInstance_(name string) *IpRemote {
	return instanceMap[name]
}

func Init(Key string, localRef *reference.Reference, redisConfig ipdata.RedisConfig, logger log.Logger) error {
	return Init_("default", Key, localRef, redisConfig, logger)
}

func Init_(name string, Key string, localRef *reference.Reference, redisConfig ipdata.RedisConfig, logger log.Logger) error {
	if name == "" {
		name = "default"
	}

	_, exist := instanceMap[name]
	if exist {
		return fmt.Errorf("ip_geo instance <%s> has already initialized", name)
	}

	ipClient := &IpRemote{}
	//new instance ipDataAndIp2Location
	ipLocalClient, err := ipdata.New(Key, localRef, redisConfig, logger)
	if err != nil {
		return err
	}
	ipClient.Client = ipLocalClient

	instanceMap[name] = ipClient
	return nil
}
