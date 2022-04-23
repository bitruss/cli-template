package examples

import (
	"github.com/coreservice-io/CliAppTemplate/basic"
	"github.com/coreservice-io/CliAppTemplate/plugin/hub"
	"github.com/coreservice-io/UHub"
)

//hub example
const testKind UHub.Kind = 1

type testEvent string

func (e testEvent) Kind() UHub.Kind {
	return testKind
}
func Hub_run() {
	hub.GetInstance().Subscribe(testKind, func(e UHub.Event) {
		basic.Logger.Debugln("hub callback")
		basic.Logger.Debugln(string(e.(testEvent)))
	})
	hub.GetInstance().Publish(testEvent("hub message"))
}
