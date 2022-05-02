package rpc_plugin

import (
	"fmt"

	"github.com/coreservice-io/bi_rpc"
	"github.com/coreservice-io/log"
)

var server_Map = map[string]*bi_rpc.Server{}

func GetServerInstance() *bi_rpc.Server {
	return server_Map["default"]
}

func GetServerInstance_(name string) *bi_rpc.Server {
	return server_Map[name]
}

func ServerInstanceInit(logger log.Logger) error {
	return ServerInstanceInit_("default", logger)
}

// Init a new instance.
// If only need one instance, use empty name "". Use GetDefaultInstance() to get.
// If you need several instance, run Init() with different <name>. Use GetInstance(<name>) to get.
func ServerInstanceInit_(name string, logger log.Logger) error {
	if name == "" {
		name = "default"
	}

	_, exist := server_Map[name]
	if exist {
		return fmt.Errorf("rpc_server instance <%s> has already been initialized", name)
	}

	server_Map[name] = bi_rpc.NewServer(logger)
	return nil
}
