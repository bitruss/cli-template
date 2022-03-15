package examples

import (
	"fmt"

	"github.com/coreservice-io/CliAppTemplate/plugin/hub"
	uhub "github.com/coreservice-io/UHub"
)

//hub example
const testKind uhub.Kind = 1

type testEvent string

func (e testEvent) Kind() uhub.Kind {
	return testKind
}
func Hub_run() {
	hub.GetInstance().Subscribe(testKind, func(e uhub.Event) {
		fmt.Println("hub callback")
		fmt.Println(string(e.(testEvent)))
	})
	hub.GetInstance().Publish(testEvent("hub message"))
}
