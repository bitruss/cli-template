package ip_remote_plugin

import (
	"fmt"

	"github.com/coreservice-io/ip_geo/ipstack"
	"github.com/coreservice-io/log"
	"github.com/coreservice-io/reference"
)

type IpRemote struct {
	Client *ipstack.IpStack
}

var instanceMap = map[string]*IpRemote{}

func GetInstance() *IpRemote {
	return instanceMap["default"]
}

func GetInstance_(name string) *IpRemote {
	return instanceMap[name]
}

func Init(ipStackKey string, localRef *reference.Reference, redisConfig ipstack.RedisConfig, logger log.Logger) error {
	return Init_("default", ipStackKey, localRef, redisConfig, logger)
}

func Init_(name string, ipStackKey string, localRef *reference.Reference, redisConfig ipstack.RedisConfig, logger log.Logger) error {
	if name == "" {
		name = "default"
	}

	_, exist := instanceMap[name]
	if exist {
		return fmt.Errorf("ip_geo instance <%s> has already initialized", name)
	}

	ipClient := &IpRemote{}
	//new instance ipStackAndIp2Location
	ipLocalClient, err := ipstack.New(ipStackKey, localRef, redisConfig, logger)
	if err != nil {
		return err
	}
	ipClient.Client = ipLocalClient

	instanceMap[name] = ipClient
	return nil
}
