package plugin

import "github.com/coreservice-io/CliAppTemplate/plugin/hub_plugin"

func iniHub() error {
	return hub_plugin.Init()
}
