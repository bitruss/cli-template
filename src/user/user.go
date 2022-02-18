package user

import (
	"context"
	"errors"
	"strconv"

	"github.com/universe-30/CliAppTemplate/basic"
	"github.com/universe-30/CliAppTemplate/plugin/cache"
	"github.com/universe-30/CliAppTemplate/plugin/redisClient"
	"github.com/universe-30/CliAppTemplate/plugin/sqldb"
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
		localvalue, _, syncOk := tools.LCR_Check(context.Background(), redisClient.GetInstance(), cache.GetInstance(), key)
		if syncOk {
			if localvalue == nil {
				return nil, nil
			} else {
				result, ok := localvalue.(*UserModel)
				if ok {
					return result, nil
				} else {
					tools.LCR_Del(context.Background(), redisClient.GetInstance(), cache.GetInstance(), key)
					basic.Logger.Errorln("GetUserById convert error")
					return nil, errors.New("convert error")
				}
			}
		}
	}

	//after cache miss ,try from remote database
	var userList []*UserModel
	err := sqldb.GetInstance().Table("user").Where("id = ?", userid).Find(&userList).Error

	if err != nil {
		basic.Logger.Errorln("GetUserById err :", err)
		return nil, err
	} else {
		if len(userList) == 0 {
			tools.LCR_Set(context.Background(), redisClient.GetInstance(), cache.GetInstance(), key, nil, 300)
			return nil, nil
		} else {
			tools.LCR_Set(context.Background(), redisClient.GetInstance(), cache.GetInstance(), key, userList[0], 300)
			return userList[0], nil
		}
	}
}
