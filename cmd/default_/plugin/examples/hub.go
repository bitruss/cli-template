package examples

import (
	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/plugin/hub_plugin"
	"github.com/coreservice-io/hub"
)

//hub example
const testKind hub.Kind = 1

type testEvent string

func (e testEvent) Kind() hub.Kind {
	return testKind
}
func Hub_run() {
	hub_plugin.GetInstance().Subscribe(testKind, func(e hub.Event) {
		basic.Logger.Debugln("hub callback")
		basic.Logger.Debugln(string(e.(testEvent)))
	})
	hub_plugin.GetInstance().Publish(testEvent("hub message"))
}
