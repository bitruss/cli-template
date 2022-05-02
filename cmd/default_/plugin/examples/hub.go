package examples

import (
	"github.com/coreservice-io/UHub"
	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/plugin/hub_plugin"
)

//hub example
const testKind UHub.Kind = 1

type testEvent string

func (e testEvent) Kind() UHub.Kind {
	return testKind
}
func Hub_run() {
	hub_plugin.GetInstance().Subscribe(testKind, func(e UHub.Event) {
		basic.Logger.Debugln("hub callback")
		basic.Logger.Debugln(string(e.(testEvent)))
	})
	hub_plugin.GetInstance().Publish(testEvent("hub message"))
}
