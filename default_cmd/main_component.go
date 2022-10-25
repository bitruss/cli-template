package default_cmd

import (
	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/component"
	"github.com/coreservice-io/cli-template/config"
)

func InitComponent() {

	config_toml := config.Get_config().Toml_config

	/////////////////////////
	if err := component.InitReference(); err != nil {
		basic.Logger.Fatalln(err)
	}
	////////////////////////
	if err := component.InitGeoIp(config_toml); err != nil {
		basic.Logger.Fatalln(err)
	}
	/////////////////////////
	if err := component.InitDB(config_toml); err != nil {
		basic.Logger.Fatalln(err)
	}
	/////////////////////////
	if err := component.InitRedis(config_toml); err != nil {
		basic.Logger.Fatalln(err)
	}
	/////////////////////////
	if err := component.InitAutoCert(config_toml); err != nil {
		basic.Logger.Fatalln(err)
	}
	/////////////////////////
	if err := component.InitEchoServer(config_toml); err != nil {
		basic.Logger.Fatalln(err)
	}
	/////////////////////////
	if err := component.InitElasticSearch(config_toml); err != nil {
		basic.Logger.Fatalln(err)
	}
	////////////////////////
	if err := component.InitEcsUploader(config_toml); err != nil {
		basic.Logger.Fatalln(err)
	}
	/////////////////////////
	if err := component.InitLevelDB(config_toml); err != nil {
		basic.Logger.Fatalln(err)
	}
	/////////////////////////
	if err := component.InitSmtpMail(config_toml); err != nil {
		basic.Logger.Fatalln(err)
	}
	/////////////////////////
	if err := component.InitSpr(config_toml); err != nil {
		basic.Logger.Fatalln(err)
	}

	/////////////////////////
	if err := component.InitSqlite(config_toml); err != nil {
		basic.Logger.Fatalln(err)
	}

}
