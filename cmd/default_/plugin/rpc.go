package plugin

import (
	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/plugin/rpc_plugin"
)

func initRpcServer() error {
	return rpc_plugin.ServerInstanceInit(basic.Logger)
}

func initRpcClient() error {
	return rpc_plugin.ClientInstanceInit(basic.Logger)
}
