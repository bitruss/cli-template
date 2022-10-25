package db_cmd

import (
	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/component"
	"github.com/coreservice-io/cli-template/config"
)

func StartDBComponent() {

	toml_conf := config.Get_config().Toml_config

	/////////////////////////
	if err := component.InitReference(); err != nil {
		basic.Logger.Fatalln(err)
	}
	/////////////////////////
	if err := component.InitDB(toml_conf); err != nil {
		basic.Logger.Fatalln(err)
	}
	/////////////////////////
	if err := component.InitRedis(toml_conf); err != nil {
		basic.Logger.Fatalln(err)
	}
}
