package plugin

import (
	"github.com/coreservice-io/CliAppTemplate/basic"
	"github.com/coreservice-io/CliAppTemplate/plugin/rpc_plugin"
)

func initRpcServer() error {
	return rpc_plugin.ServerInstanceInit(basic.Logger)
}

func initRpcClient() error {
	return rpc_plugin.ClientInstanceInit(basic.Logger)
}
