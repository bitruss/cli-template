package user

import (
	"context"
	"errors"
	"strconv"

	"github.com/universe-30/CliAppTemplate/basic"
	"github.com/universe-30/CliAppTemplate/cmd/default_/global"
	"github.com/universe-30/CliAppTemplate/tools"
)

//example for GormDB and tools cache
type UserModel struct {
	Id               int
	Status           string
	Name             string
	Email            string
	Created_unixtime int64
}

func GetUserById(userid int, forceupdate bool) (*UserModel, error) {
	key := "finance:user:" + strconv.Itoa(userid)
	if !forceupdate {
		localvalue, _, syncOk := tools.SmartCheck_LocalCache_Redis(context.Background(), global.Redis, global.Cache, key)
		if syncOk {
			if localvalue == nil {
				return nil, nil
			} else {
				result, ok := localvalue.(*UserModel)
				if ok {
					return result, nil
				} else {
					tools.SmartDel_LocalCache_Redis(context.Background(), global.Redis, global.Cache, key)
					basic.Logger.Errorln("GetUserById convert error")
					return nil, errors.New("convert error")
				}
			}
		}
	}

	//after cache miss ,try from remote database
	var userList []*UserModel
	err := global.DB.Table("user").Where("id = ?", userid).Find(&userList).Error

	if err != nil {
		basic.Logger.Errorln("GetUserById err :", err)
		return nil, err
	} else {
		if len(userList) == 0 {
			tools.SmartSet_LocalCache_Redis(context.Background(), global.Redis, global.Cache, key, nil, 300)
			return nil, nil
		} else {
			tools.SmartSet_LocalCache_Redis(context.Background(), global.Redis, global.Cache, key, userList[0], 300)
			return userList[0], nil
		}
	}
}
