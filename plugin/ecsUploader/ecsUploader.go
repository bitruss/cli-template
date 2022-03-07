package ecsUploader

import (
	"fmt"

	"github.com/coreservice-io/UECSUploader/uploader"
)

var instanceMap = map[string]*uploader.Uploader{}

func GetInstance() *uploader.Uploader {
	return instanceMap["default"]
}

func GetInstance_(name string) *uploader.Uploader {
	return instanceMap[name]
}

/*
elasticSearchAddr
elasticSearchUserName
elasticSearchPassword
*/
type Config struct {
	Address  string
	UserName string
	Password string
}

func Init(esConfig Config) error {
	return Init_("default", esConfig)
}

//  Init a new instance.
//  If only need one instance, use empty name "". Use GetDefaultInstance() to get.
//  If you need several instance, run Init() with different <name>. Use GetInstance(<name>) to get.
func Init_(name string, esConfig Config) error {
	if name == "" {
		name = "default"
	}

	_, exist := instanceMap[name]
	if exist {
		return fmt.Errorf("elasticSearch instance <%s> has already initialized", name)
	}

	es, err := uploader.New(esConfig.Address, esConfig.UserName, esConfig.Password)
	if err != nil {
		return err
	}
	instanceMap[name] = es
	return nil
}
