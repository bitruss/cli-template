package exampleUser

import (
	"context"
	"errors"
	"strconv"

	"github.com/coreservice-io/CliAppTemplate/basic"
	"github.com/coreservice-io/CliAppTemplate/plugin/cache"
	"github.com/coreservice-io/CliAppTemplate/plugin/redisClient"
	"github.com/coreservice-io/CliAppTemplate/plugin/sqldb"
	lrc "github.com/coreservice-io/CliAppTemplate/tools/lrcache"
)

//example for GormDB and tools cache
type ExampleUserModel struct {
	ID      int
	Status  string
	Name    string
	Email   string
	Updated int64 `gorm:"autoUpdateTime"`
	Created int64 `gorm:"autoCreateTime"`
}

func CreateUser(userInfo *ExampleUserModel) (*ExampleUserModel, error) {
	//userInfo in param data
	//&ExampleUserModel{
	//	Status: "normal",
	//	Name:"userName",
	//	Email:"mail@email.com",
	//}

	if err := sqldb.GetInstance().Create(userInfo).Error; err != nil {
		return nil, err
	}
	GetUserById(userInfo.ID, true)
	return userInfo, nil
}

func DeleteUser(id int) error {
	user := &ExampleUserModel{ID: id}
	if err := sqldb.GetInstance().Delete(user).Error; err != nil {
		return err
	}

	//delete cache
	key := redisClient.GetInstance().GenKey("user", strconv.Itoa(id))
	lrc.LRC_Del(context.Background(), redisClient.GetInstance(), cache.GetInstance(), key)

	return nil
}

func UpdateUser(newData map[string]interface{}, id int) error {
	user := &ExampleUserModel{ID: id}

	//newData in param data
	//newData= map[string]interface{}{
	//	"status":"error",
	//	"name":"userName2",
	//	"email":"mail2@email.com",
	//}

	result := sqldb.GetInstance().Model(user).Updates(newData)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("record not exist")
	}

	//refresh cache
	GetUserById(id, true)

	return nil
}

func GetUserById(userid int, forceupdate bool) (*ExampleUserModel, error) {
	key := redisClient.GetInstance().GenKey("user", strconv.Itoa(userid))
	if !forceupdate {
		result := lrc.LRC_Get(context.Background(), redisClient.GetInstance(), cache.GetInstance(), true, key)
		if result != nil {
			value, ok := result.(*ExampleUserModel)
			if ok {
				return value, nil
			} else {
				nullValue, ok := result.(string)
				if ok && nullValue == lrc.TEMP_NULL {
					return nil, nil
				}
				lrc.LRC_Del(context.Background(), redisClient.GetInstance(), cache.GetInstance(), key)
				basic.Logger.Errorln("GetUsers convert error, result:", result)
				return nil, errors.New("GetUsers convert error")
			}
		}
	}

	//after cache miss ,try from remote database
	var userList []*ExampleUserModel
	err := sqldb.GetInstance().Model(&ExampleUserModel{}).Where("id = ?", userid).Find(&userList).Error

	if err != nil {
		basic.Logger.Errorln("GetUserById err :", err)
		return nil, err
	} else {
		if len(userList) == 0 {
			lrc.LRC_Set(context.Background(), redisClient.GetInstance(), cache.GetInstance(), false, key, lrc.TEMP_NULL, 300)
			return nil, nil
		} else {
			lrc.LRC_Set(context.Background(), redisClient.GetInstance(), cache.GetInstance(), true, key, userList[0], 300)
			return userList[0], nil
		}
	}
}

func GetUsers(username string, forceupdate bool) ([]*ExampleUserModel, error) {
	key := redisClient.GetInstance().GenKey("getusers", username)
	if !forceupdate {
		result := lrc.LRC_Get(context.Background(), redisClient.GetInstance(), cache.GetInstance(), true, key)
		if result != nil {
			value, ok := result.([]*ExampleUserModel)
			if ok {
				return value, nil
			} else {
				lrc.LRC_Del(context.Background(), redisClient.GetInstance(), cache.GetInstance(), key)
				basic.Logger.Errorln("GetUsers convert error, result:", result)
				return nil, errors.New("GetUsers convert error")
			}
		}
	}

	//after cache miss ,try from remote database
	var userList []*ExampleUserModel
	err := sqldb.GetInstance().Model(&ExampleUserModel{}).Where("name = ?", username).Find(&userList).Error

	if err != nil {
		basic.Logger.Errorln("GetUsers err :", err)
		return nil, err
	} else {
		lrc.LRC_Set(context.Background(), redisClient.GetInstance(), cache.GetInstance(), true, key, userList, 300)
		return userList, nil
	}
}
