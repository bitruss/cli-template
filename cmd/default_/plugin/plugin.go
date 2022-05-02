package plugin

import (
	"github.com/coreservice-io/cli-template/basic"
)

//todo: ---
func InitPlugin() {

	err := iniHub()
	if err != nil {
		basic.Logger.Fatalln(err)
	}

	/////////////////////////
	err = initReference()
	if err != nil {
		basic.Logger.Fatalln(err)
	}

	/////////////////////////
	err = initEchoServer()
	if err != nil {
		basic.Logger.Fatalln(err)
	}

}
