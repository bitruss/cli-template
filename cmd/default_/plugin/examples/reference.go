package examples

import (
	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/plugin/reference_plugin"
)

//cache example
func Reference_run() {
	bar1 := "bar1"
	err1 := reference_plugin.GetInstance_("ref1").Set("foo1", &bar1, 10)
	if err1 != nil {
		basic.Logger.Errorln(err1)
	}
	v, _ := reference_plugin.GetInstance_("ref1").Get("foo1")
	if v != nil {
		basic.Logger.Debugln(*(v.(*string)))
	}

	bar2 := "bar2"
	err2 := reference_plugin.GetInstance_("ref1").Set("foo2", &bar2, 10)
	if err2 != nil {
		basic.Logger.Errorln(err2)
	}
	v, _ = reference_plugin.GetInstance_("ref1").Get("foo2")
	if v != nil {
		basic.Logger.Debugln(*(v.(*string)))
	}
}
