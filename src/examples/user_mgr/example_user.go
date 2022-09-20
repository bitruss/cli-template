package user_mgr

import (
	"context"
	"time"

	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/plugin/redis_plugin"
	"github.com/coreservice-io/cli-template/plugin/reference_plugin"
	"github.com/coreservice-io/cli-template/plugin/sqldb_plugin"
	"github.com/coreservice-io/cli-template/src/common/smart_cache"
	"github.com/go-redis/redis/v8"
)

//example for GormDB and tools cache
type ExampleUserModel struct {
	Id               int64  `json:"id"`
	Status           string `json:"status"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	Updated_unixtime int64  `json:"updated_unixtime" gorm:"autoUpdateTime"`
	Created_unixtime int64  `json:"created_unixtime" gorm:"autoCreateTime"`
}

func CreateUser(userInfo *ExampleUserModel) (*ExampleUserModel, error) {
	if err := sqldb_plugin.GetInstance().Table("example_user_models").Create(userInfo).Error; err != nil {
		return nil, err
	}
	return userInfo, nil
}

func DeleteUser(id int64) error {
	user := &ExampleUserModel{Id: id}
	if err := sqldb_plugin.GetInstance().Table("example_user_models").Delete(user).Error; err != nil {
		return err
	}
	//delete cache if necessary
	QueryUser(&id, nil, nil, nil, 0, 0, false, true)
	return nil
}

func UpdateUser(newData map[string]interface{}, id int64) error {
	newData["updated"] = time.Now().UTC().Unix()
	result := sqldb_plugin.GetInstance().Table("example_user_models").Where("id=?", id).Updates(newData)
	if result.Error != nil {
		return result.Error
	}
	//refresh cache if necessary
	QueryUser(&id, nil, nil, nil, 0, 0, false, true)
	return nil
}

//query
type QueryUserResult struct {
	Users       []*ExampleUserModel `json:"users"`
	Total_count int64               `json:"total_count"`
}

func QueryUser(id *int64, status *string, name *string, email *string, limit int, offset int, fromCache bool, updateCache bool) (*QueryUserResult, error) {
	//gen_key
	ck := smart_cache.NewConnectKey("user")
	ck.C_Int64_Ptr("id", id).C_Str_Ptr("status", status).
		C_Str_Ptr("name", name).C_Str_Ptr("email", email).C_Int(limit).C_Int(offset)

	key := redis_plugin.GetInstance().GenKey(ck.String())

	if fromCache {
		// try to get from reference
		result := smart_cache.Ref_Get(reference_plugin.GetInstance(), key)
		if result != nil {
			basic.Logger.Debugln("QueryUser hit from reference")
			return result.(*QueryUserResult), nil
		}

		// try to get from redis
		redis_result := &QueryUserResult{
			Users:       []*ExampleUserModel{},
			Total_count: 0,
		}
		err := smart_cache.Redis_Get(context.Background(), redis_plugin.GetInstance().ClusterClient, true, key, redis_result)
		if err == nil {
			basic.Logger.Debugln("QueryUser hit from redis")
			smart_cache.Ref_Set(reference_plugin.GetInstance(), key, redis_result)
			return redis_result, nil
		} else if err == redis.Nil {
			//continue to get from db part
		} else if err == smart_cache.ErrTempNil {
			//won't happen actually unless you set a nil pointer of queryResult when update
			basic.Logger.Errorln("QueryUser smart_cache.TempNil")
			return nil, smart_cache.ErrTempNil
		} else {
			//redis may broken, just return to keep db safe
			return redis_result, err
		}
	}

	//after cache miss ,try from remote database
	basic.Logger.Debugln("QueryUser try from database")

	queryResult := &QueryUserResult{
		Users:       []*ExampleUserModel{},
		Total_count: 0,
	}

	query := sqldb_plugin.GetInstance().Table("example_user_models")
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

	query.Count(&queryResult.Total_count)
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
			smart_cache.RR_Set(context.Background(), redis_plugin.GetInstance().ClusterClient, reference_plugin.GetInstance(), true, key, queryResult, 300)
		}
		return queryResult, nil
	}
}
