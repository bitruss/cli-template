package plugin

import (
	"github.com/coreservice-io/cli-template/basic"
)

//todo: ---
func InitPlugin() {

	/////////////////////////
	err := initReference()
	if err != nil {
		basic.Logger.Fatalln("initReference err:", err)
	}

	/////////////////////////
	// err = initAutoCert()
	// if err != nil {
	// 	basic.Logger.Fatalln("initAutoCert err:", err)
	// }

	/////////////////////////
	err = initEchoServer()
	if err != nil {
		basic.Logger.Fatalln("initEchoServer err:", err)
	}

}
