package userMgr

import (
	"context"
	"strconv"
	"time"

	"github.com/coreservice-io/CliAppTemplate/basic"
	"github.com/coreservice-io/CliAppTemplate/plugin/redisClient"
	"github.com/coreservice-io/CliAppTemplate/plugin/reference"
	"github.com/coreservice-io/CliAppTemplate/plugin/sqldb"
	"github.com/coreservice-io/CliAppTemplate/tools/smartCache"
	"github.com/go-redis/redis/v8"
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
	//GetUserById(userInfo.ID, true)
	return userInfo, nil
}

func DeleteUser(id int) error {
	user := &ExampleUserModel{ID: id}
	if err := sqldb.GetInstance().Table("example_user_models").Delete(user).Error; err != nil {
		return err
	}

	//delete cache
	key := redisClient.GetInstance().GenKey("user", strconv.Itoa(id))
	smartCache.RR_Del(context.Background(), redisClient.GetInstance(), reference.GetInstance(), key)

	return nil
}

func UpdateUser(newData map[string]interface{}, id int) error {
	//newData in param data
	//newData= map[string]interface{}{
	//	"status":"error",
	//	"name":"userName2",
	//	"email":"mail2@email.com",
	//}

	newData["updated"] = time.Now().UTC().Unix()
	result := sqldb.GetInstance().Table("example_user_models").Where("id=?", id).Updates(newData)
	if result.Error != nil {
		return result.Error
	}

	//refresh cache
	GetUserById(id, true)

	return nil
}

func GetUserById(userid int, forceupdate bool) (*ExampleUserModel, error) {
	key := redisClient.GetInstance().GenKey("user", strconv.Itoa(userid))
	if !forceupdate {
		// try to get from reference
		result := smartCache.Ref_Get(reference.GetInstance(), key)
		if result != nil {
			return result.(*ExampleUserModel), nil
		}

		// try to get from redis
		redis_result := &ExampleUserModel{}
		err := smartCache.Redis_Get(context.Background(), redisClient.GetInstance(), true, key, redis_result)
		if err == redis.Nil {
			return nil, nil
		}
		if err == nil {
			return redis_result, nil
		}
	}

	//after cache miss ,try from remote database
	var userList []*ExampleUserModel
	err := sqldb.GetInstance().Table("example_user_models").Where("id = ?", userid).Find(&userList).Error
	if err != nil {
		basic.Logger.Errorln("GetUserById err :", err)
		return nil, err
	} else {
		if len(userList) == 0 {
			smartCache.RR_Set(context.Background(), redisClient.GetInstance(), reference.GetInstance(), false, key, nil, 300)
			return nil, nil
		} else {
			smartCache.RR_Set(context.Background(), redisClient.GetInstance(), reference.GetInstance(), true, key, userList[0], 300)
			return userList[0], nil
		}
	}
}
