package plugin

// import (
// 	"errors"

// 	"github.com/coreservice-io/cli-template/basic"
// 	"github.com/coreservice-io/cli-template/configuration"
// 	"github.com/coreservice-io/cli-template/plugin/ecs_uploader_plugin"
// )

// func initEcsUploader() error {
// 	elasticSearchAddr, err := configuration.Config.GetString("elasticsearch.host", "")
// 	if err != nil {
// 		return errors.New("elasticsearch.host [string] in config error," + err.Error())
// 	}

// 	elasticSearchUserName, err := configuration.Config.GetString("elasticsearch.username", "")
// 	if err != nil {
// 		return errors.New("elasticsearch.username err [string] in config error," + err.Error())
// 	}

// 	elasticSearchPassword, err := configuration.Config.GetString("elasticsearch.password", "")
// 	if err != nil {
// 		return errors.New("elasticsearch.password [string] in config error," + err.Error())
// 	}

// 	return ecs_uploader_plugin.Init(&ecs_uploader_plugin.Config{
// 		Address:  elasticSearchAddr,
// 		UserName: elasticSearchUserName,
// 		Password: elasticSearchPassword}, basic.Logger)

// }
