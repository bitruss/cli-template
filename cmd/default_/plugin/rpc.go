package plugin

import (
	"github.com/coreservice-io/CliAppTemplate/basic"
	"github.com/coreservice-io/CliAppTemplate/plugin/rpc"
)

func initRpcServer() error {
	return rpc.ServerInstanceInit(basic.Logger)
}

func initRpcClient() error {
	return rpc.ClientInstanceInit(basic.Logger)
}
