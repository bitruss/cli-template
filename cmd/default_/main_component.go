package default_

import (
	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/component"
)

func InitComponent() {

	/////////////////////////
	if err := component.InitReference(); err != nil {
		basic.Logger.Fatalln(err)
	}
	////////////////////////
	if err := component.InitGeoIp(); err != nil {
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
	/////////////////////////
	if err := component.InitAutoCert(); err != nil {
		basic.Logger.Fatalln(err)
	}
	/////////////////////////
	if err := component.InitEchoServer(); err != nil {
		basic.Logger.Fatalln(err)
	}
	/////////////////////////
	if err := component.InitElasticSearch(); err != nil {
		basic.Logger.Fatalln(err)
	}
	////////////////////////
	if err := component.InitEcsUploader(); err != nil {
		basic.Logger.Fatalln(err)
	}
	/////////////////////////
	if err := component.InitLevelDB(); err != nil {
		basic.Logger.Fatalln(err)
	}
	/////////////////////////
	if err := component.InitSmtpMail(); err != nil {
		basic.Logger.Fatalln(err)
	}
	/////////////////////////
	if err := component.InitSpr(); err != nil {
		basic.Logger.Fatalln(err)
	}

	/////////////////////////
	if err := component.InitSqlite(); err != nil {
		basic.Logger.Fatalln(err)
	}

}
