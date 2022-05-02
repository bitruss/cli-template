package plugin

import "github.com/coreservice-io/cli-template/plugin/hub_plugin"

func iniHub() error {
	return hub_plugin.Init()
}
