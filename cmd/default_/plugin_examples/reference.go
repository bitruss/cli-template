package examples

import (
	"github.com/coreservice-io/CliAppTemplate/basic"
	"github.com/coreservice-io/CliAppTemplate/plugin/reference"
)

//cache example
func Reference_run() {
	reference.GetInstance_("cache1").Set("foo1", "bar1", 10)
	v, _, exist := reference.GetInstance_("cache1").Get("foo1")
	if exist {
		basic.Logger.Debugln(v.(string))
	}

	reference.GetInstance_("cache2").Set("foo2", "bar2", 10)
	v, _, exist = reference.GetInstance_("cache2").Get("foo2")
	if exist {
		basic.Logger.Debugln(v.(string))
	}
}
