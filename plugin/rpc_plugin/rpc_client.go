package rpc_plugin

import (
	"fmt"

	"github.com/coreservice-io/bi_rpc"
	"github.com/coreservice-io/log"
)

var client_Map = map[string]*bi_rpc.Client{}

func GetClient() *bi_rpc.Client {
	return client_Map["default"]
}

func GetClient_(name string) *bi_rpc.Client {
	return client_Map[name]
}

func ClientInstanceInit(logger log.Logger) error {
	return ClientInstanceInit_("default", logger)
}

// Init a new instance.
// If only need one instance, use empty name "". Use GetDefaultInstance() to get.
// If you need several instance, run Init() with different <name>. Use GetInstance(<name>) to get.
func ClientInstanceInit_(name string, logger log.Logger) error {
	if name == "" {
		name = "default"
	}

	_, exist := client_Map[name]
	if exist {
		return fmt.Errorf("rpc_client instance <%s> has already been initialized", name)
	}

	client_Map[name] = bi_rpc.NewClient(logger)

	return nil
}
