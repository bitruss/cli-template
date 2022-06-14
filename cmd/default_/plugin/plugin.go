package plugin

import (
	"github.com/coreservice-io/cli-template/basic"
)

//todo: ---
func InitPlugin() {

	/////////////////////////
	err := initAutoCert()
	if err != nil {
		basic.Logger.Fatalln("initAutoCert err:", err)
	}

	/////////////////////////
	err = initDB()
	if err != nil {
		basic.Logger.Fatalln("initDB err:", err)
	}

	/////////////////////////
	err = initEchoServer()
	if err != nil {
		basic.Logger.Fatalln("initEchoServer err:", err)
	}

	/////////////////////////
	err = initElasticSearch()
	if err != nil {
		basic.Logger.Fatalln("initElasticSearch err:", err)
	}

	////////////////////////
	err = initEcsUploader()
	if err != nil {
		basic.Logger.Fatalln("initEcsUploader err:", err)
	}

	////////////////////////
	err = initIpGeo()
	if err != nil {
		basic.Logger.Fatalln("initIpGeo err:", err)
	}

	/////////////////////////
	err = initLevelDB()
	if err != nil {
		basic.Logger.Fatalln("initLevelDB err:", err)
	}

	/////////////////////////
	err = initSmtpMail()
	if err != nil {
		basic.Logger.Fatalln("initSmtpMail err:", err)
	}

	/////////////////////////
	err = initRedis()
	if err != nil {
		basic.Logger.Fatalln("initRedis err:", err)
	}

	/////////////////////////
	err = initReference()
	if err != nil {
		basic.Logger.Fatalln("initReference err:", err)
	}

	/////////////////////////
	err = initSqlite()
	if err != nil {
		basic.Logger.Fatalln("initSqlite err:", err)
	}

}
