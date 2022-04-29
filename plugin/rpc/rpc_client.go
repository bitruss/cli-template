package rpc

import (
	"fmt"

	"github.com/coreservice-io/UBiRpc"
	"github.com/coreservice-io/ULog"
)

var client_Map = map[string]*UBiRpc.Client{}

func GetClient() *UBiRpc.Client {
	return client_Map["default"]
}

func GetClient_(name string) *UBiRpc.Client {
	return client_Map[name]
}

func ClientInstanceInit(logger ULog.Logger) error {
	return ClientInstanceInit_("default", logger)
}

// Init a new instance.
// If only need one instance, use empty name "". Use GetDefaultInstance() to get.
// If you need several instance, run Init() with different <name>. Use GetInstance(<name>) to get.
func ClientInstanceInit_(name string, logger ULog.Logger) error {
	if name == "" {
		name = "default"
	}

	_, exist := client_Map[name]
	if exist {
		return fmt.Errorf("client instance <%s> has already initialized", name)
	}

	client_Map[name] = UBiRpc.NewClient(logger)

	return nil
}
