package examples

import (
	"github.com/coreservice-io/CliAppTemplate/basic"
	"github.com/coreservice-io/CliAppTemplate/plugin/cache"
)

//cache example
func Cache_run() {
	cache.GetInstance_("cache1").Set("foo1", "bar1", 10)
	v, _, exist := cache.GetInstance_("cache1").Get("foo1")
	if exist {
		basic.Logger.Debugln(v.(string))
	}

	cache.GetInstance_("cache2").Set("foo2", "bar2", 10)
	v, _, exist = cache.GetInstance_("cache2").Get("foo2")
	if exist {
		basic.Logger.Debugln(v.(string))
	}
}
