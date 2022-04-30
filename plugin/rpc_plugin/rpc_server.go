package rpc_plugin

import (
	"fmt"

	"github.com/coreservice-io/UBiRpc"
	"github.com/coreservice-io/ULog"
)

var server_Map = map[string]*UBiRpc.Server{}

func GetServerInstance() *UBiRpc.Server {
	return server_Map["default"]
}

func GetServerInstance_(name string) *UBiRpc.Server {
	return server_Map[name]
}

func ServerInstanceInit(logger ULog.Logger) error {
	return ServerInstanceInit_("default", logger)
}

// Init a new instance.
// If only need one instance, use empty name "". Use GetDefaultInstance() to get.
// If you need several instance, run Init() with different <name>. Use GetInstance(<name>) to get.
func ServerInstanceInit_(name string, logger ULog.Logger) error {
	if name == "" {
		name = "default"
	}

	_, exist := server_Map[name]
	if exist {
		return fmt.Errorf("server instance <%s> has already initialized", name)
	}

	server_Map[name] = UBiRpc.NewServer(logger)
	return nil
}
