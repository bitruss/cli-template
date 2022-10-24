package db

import (
	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/component"
)

func startDBComponent() {
	/////////////////////////
	if err := component.InitReference(); err != nil {
		basic.Logger.Fatalln(err)
	}
	/////////////////////////
	if err := component.InitDB(); err != nil {
		basic.Logger.Fatalln(err)
	}
	/////////////////////////
	if err := component.InitRedis(); err != nil {
		basic.Logger.Fatalln(err)
	}
}
