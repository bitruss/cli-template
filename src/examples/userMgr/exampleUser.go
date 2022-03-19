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
			basic.Logger.Debugln("GetUserById hit from reference")
			return result.(*ExampleUserModel), nil
		}

		// try to get from redis
		redis_result := &ExampleUserModel{}
		err := smartCache.Redis_Get(context.Background(), redisClient.GetInstance().ClusterClient, true, key, redis_result)
		if err == nil {
			basic.Logger.Debugln("GetUserById hit from redis")
			smartCache.Ref_Set(reference.GetInstance(), key, redis_result)
			return redis_result, nil
		} else if err == smartCache.TempNil {
			return nil, nil
		} else if err == redis.Nil {
			//continue to get from db part
		} else {
			//redis may broken, just return to keep db safe
			return nil, err
		}
	}

	//after cache miss ,try from remote database
	basic.Logger.Debugln("GetUserById try from db")
	var userList []*ExampleUserModel
	err := sqldb.GetInstance().Table("example_user_models").Where("id = ?", userid).Find(&userList).Error
	if err != nil {
		basic.Logger.Errorln("GetUserById err :", err)
		return nil, err
	} else {
		if len(userList) == 0 {
			smartCache.RR_Set(context.Background(), redisClient.GetInstance().ClusterClient, reference.GetInstance(), false, key, nil, 300)
			return nil, nil
		} else {
			smartCache.RR_Set(context.Background(), redisClient.GetInstance().ClusterClient, reference.GetInstance(), true, key, userList[0], 300)
			return userList[0], nil
		}
	}
}

func GetUsersByStatus(status string, forceupdate bool) ([]*ExampleUserModel, error) {
	key := redisClient.GetInstance().GenKey("users", "status", status)
	if !forceupdate {
		// try to get from reference
		result := smartCache.Ref_Get(reference.GetInstance(), key)
		if result != nil {
			basic.Logger.Debugln("GetUsersByStatus hit from reference")
			return result.([]*ExampleUserModel), nil
		}

		// try to get from redis
		redis_result := []*ExampleUserModel{}
		err := smartCache.Redis_Get(context.Background(), redisClient.GetInstance().ClusterClient, true, key, &redis_result)
		if err == nil {
			basic.Logger.Debugln("GetUsersByStatus hit from redis")
			smartCache.Ref_Set(reference.GetInstance(), key, redis_result)
			return redis_result, nil
		} else if err == redis.Nil {
			//continue to get from db part
		} else {
			//redis may broken, just return to keep db safe
			return nil, err
		}
	}

	//after cache miss ,try from remote database
	basic.Logger.Debugln("GetUsersByStatus try from database")
	var userList []*ExampleUserModel
	err := sqldb.GetInstance().Table("example_user_models").Where("status = ?", status).Find(&userList).Error
	if err != nil {
		basic.Logger.Errorln("GetUsersByStatus err :", err)
		return nil, err
	} else {
		smartCache.RR_Set(context.Background(), redisClient.GetInstance().ClusterClient, reference.GetInstance(), true, key, userList, 300)
		return userList, nil
	}
}

// not recommended usage
func GetUserNameById(userid int, forceupdate bool) (string, error) {
	key := redisClient.GetInstance().GenKey("user", "name", strconv.Itoa(userid))
	if !forceupdate {
		// try to get from reference
		result := smartCache.Ref_Get(reference.GetInstance(), key)
		if result != nil {
			basic.Logger.Debugln("GetUserNameById hit from reference")
			return *result.(*string), nil
		}

		// try to get from redis
		var redis_result string
		err := smartCache.Redis_Get(context.Background(), redisClient.GetInstance().ClusterClient, false, key, &redis_result)
		if err == nil {
			basic.Logger.Debugln("GetUserNameById hit from redis")
			smartCache.Ref_Set(reference.GetInstance(), key, redis_result)
			return redis_result, nil
		} else if err == smartCache.TempNil {
			return "", nil
		} else if err == redis.Nil {
			//continue to get from db part
		} else {
			//redis may broken, just return to keep db safe
			return "", err
		}
	}

	//after cache miss ,try from remote database
	basic.Logger.Debugln("GetUserNameById try from db")
	var userName []string
	err := sqldb.GetInstance().Table("example_user_models").Select("name").Where("id = ?", userid).Find(&userName).Error
	if err != nil {
		basic.Logger.Errorln("GetUserById err :", err)
		return "", err
	} else {
		if len(userName) == 0 {
			smartCache.RR_Set(context.Background(), redisClient.GetInstance().ClusterClient, reference.GetInstance(), false, key, nil, 300)
			return "", nil
		} else {
			smartCache.RR_Set(context.Background(), redisClient.GetInstance().ClusterClient, reference.GetInstance(), false, key, &userName[0], 300)
			return userName[0], nil
		}
	}
}
