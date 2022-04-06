package userMgr

import (
	"context"
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
	if err := sqldb.GetInstance().Table("example_user_models").Create(userInfo).Error; err != nil {
		return nil, err
	}
	return userInfo, nil
}

func DeleteUser(id int) error {
	user := &ExampleUserModel{ID: id}
	if err := sqldb.GetInstance().Table("example_user_models").Delete(user).Error; err != nil {
		return err
	}
	//delete cache if necessary
	QueryUser(&id, nil, nil, nil, 0, 0, false, true)
	return nil
}

func UpdateUser(newData map[string]interface{}, id int) error {
	newData["updated"] = time.Now().UTC().Unix()
	result := sqldb.GetInstance().Table("example_user_models").Where("id=?", id).Updates(newData)
	if result.Error != nil {
		return result.Error
	}
	//refresh cache if necessary
	QueryUser(&id, nil, nil, nil, 0, 0, false, true)
	return nil
}

//query
type QueryUserResult struct {
	Users      []*ExampleUserModel
	TotalCount int64
}

func QueryUser(id *int, status *string, name *string, email *string, limit int, offset int, fromCache bool, updateCache bool) (*QueryUserResult, error) {
	//gen_key
	ck := smartCache.ConnectKey{}
	ck.C_Int_Ptr("id", id).C_Str_Ptr("status", status).
		C_Str_Ptr("name", name).C_Str_Ptr("email", email).C_Int(limit).C_Int(offset)

	key := redisClient.GetInstance().GenKey(ck.String())

	if fromCache {
		// try to get from reference
		result := smartCache.Ref_Get(reference.GetInstance(), key)
		if result != nil {
			basic.Logger.Debugln("QueryUser hit from reference")
			return result.(*QueryUserResult), nil
		}

		// try to get from redis
		redis_result := &QueryUserResult{
			Users:      []*ExampleUserModel{},
			TotalCount: 0,
		}
		err := smartCache.Redis_Get(context.Background(), redisClient.GetInstance().ClusterClient, true, key, redis_result)
		if err == nil {
			basic.Logger.Debugln("QueryUser hit from redis")
			smartCache.Ref_Set(reference.GetInstance(), key, redis_result)
			return redis_result, nil
		} else if err == redis.Nil {
			//continue to get from db part
		} else {
			//redis may broken, just return to keep db safe
			return redis_result, err
		}
	}

	//after cache miss ,try from remote database
	basic.Logger.Debugln("QueryUser try from database")

	queryResult := &QueryUserResult{
		Users:      []*ExampleUserModel{},
		TotalCount: 0,
	}

	query := sqldb.GetInstance().Table("example_user_models")
	if id != nil {
		query.Where("id = ?", *id)
	}
	if status != nil {
		query.Where("status = ?", status)
	}
	if name != nil {
		query.Where("name = ?", name)
	}
	if email != nil {
		query.Where("email = ?", email)
	}

	query.Count(&queryResult.TotalCount)
	if limit > 0 {
		query.Limit(limit)
	}
	if offset > 0 {
		query.Offset(offset)
	}

	err := query.Find(&queryResult.Users).Error
	if err != nil {
		basic.Logger.Errorln("QueryUser err :", err)
		return nil, err
	} else {
		if updateCache {
			smartCache.RR_Set(context.Background(), redisClient.GetInstance().ClusterClient, reference.GetInstance(), true, key, queryResult, 300)
		}
		return queryResult, nil
	}
}
