package ip_local_plugin

import (
	"fmt"

	"github.com/coreservice-io/ip_geo/ip2l"
	"github.com/coreservice-io/log"
)

type IpLocal struct {
	Client *ip2l.Ip2LClient
}

var instanceMap = map[string]*IpLocal{}

func GetInstance() *IpLocal {
	return instanceMap["default"]
}

func GetInstance_(name string) *IpLocal {
	return instanceMap[name]
}

func Init(localDbFile string, ip2LUpgradeUrl string, upgradeInterval int64, logger log.Logger, panicHandler func(interface{})) error {
	return Init_("default", localDbFile, ip2LUpgradeUrl, upgradeInterval, logger, panicHandler)
}

func Init_(name string, localDbFile string, ip2LUpgradeUrl string, upgradeInterval int64, logger log.Logger, panicHandler func(interface{})) error {
	if name == "" {
		name = "default"
	}

	_, exist := instanceMap[name]
	if exist {
		return fmt.Errorf("ip_geo instance <%s> has already initialized", name)
	}

	ipClient := &IpLocal{}
	//new instance IpDataAndIp2Location
	ipLocalClient, err := ip2l.New(localDbFile, ip2LUpgradeUrl, upgradeInterval, logger, panicHandler)
	if err != nil {
		return err
	}
	ipClient.Client = ipLocalClient

	instanceMap[name] = ipClient
	return nil
}
