package geo_ip_plugin

import (
	"fmt"

	"github.com/coreservice-io/geo_ip/lib"
)

type GeoIp struct {
	lib.GeoIpInterface
}

var instanceMap = map[string]*GeoIp{}

func GetInstance() *GeoIp {
	return instanceMap["default"]
}

func GetInstance_(name string) *GeoIp {
	return instanceMap[name]
}

func Init(localDbFile string) error {
	return Init_("default", localDbFile)
}

func Init_(name string, localDbFile string) error {
	if name == "" {
		name = "default"
	}

	_, exist := instanceMap[name]
	if exist {
		return fmt.Errorf("ip_geo instance <%s> has already initialized", name)
	}

	ipClient := &GeoIp{}
	//new instance
	client, err := lib.NewClient(localDbFile)
	if err != nil {
		return err
	}
	ipClient.GeoIpInterface = client

	instanceMap[name] = ipClient
	return nil
}
