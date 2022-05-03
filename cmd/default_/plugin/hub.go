package plugin

import "github.com/coreservice-io/cli-template/plugin/hub_plugin"

func initHub() error {
	return hub_plugin.Init()
}
