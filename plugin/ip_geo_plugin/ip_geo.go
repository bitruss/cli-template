package ip_geo_plugin

import (
	"fmt"

	"github.com/coreservice-io/ip_geo"
	"github.com/coreservice-io/ip_geo/ipstack_ip2l"
	"github.com/coreservice-io/log"
	"github.com/coreservice-io/reference"
)

type IpLocation struct {
	Client ip_geo.IIpGeo
}

var instanceMap = map[string]*IpLocation{}

func GetInstance() *IpLocation {
	return instanceMap["default"]
}

func GetInstance_(name string) *IpLocation {
	return instanceMap[name]
}

func Init(ipStackKey string, localDbFile string, ip2LUpgradeUrl string, upgradeInterval int64,
	localRef *reference.Reference, redisConfig ip_geo.RedisConfig, logger log.Logger, panicHandler func(interface{})) error {
	return Init_("default", ipStackKey, localDbFile, ip2LUpgradeUrl, upgradeInterval, localRef, redisConfig, logger, panicHandler)
}

func Init_(name string, ipStackKey string, localDbFile string, ip2LUpgradeUrl string,
	upgradeInterval int64, localRef *reference.Reference, redisConfig ip_geo.RedisConfig,
	logger log.Logger, panicHandler func(interface{})) error {
	if name == "" {
		name = "default"
	}

	_, exist := instanceMap[name]
	if exist {
		return fmt.Errorf("mail sender instance <%s> has already initialized", name)
	}

	ipClient := &IpLocation{}
	//new instance ipStackAndIp2Location
	ipStackIp2LClient, err := ipstack_ip2l.New(ipStackKey, localDbFile, ip2LUpgradeUrl, upgradeInterval, localRef, redisConfig, logger, panicHandler)
	if err != nil {
		return err
	}
	ipClient.Client = ipStackIp2LClient

	instanceMap[name] = ipClient
	return nil
}
