package plugin

import "github.com/coreservice-io/CliAppTemplate/plugin/hub"

func iniHub() error {
	return hub.Init()
}
