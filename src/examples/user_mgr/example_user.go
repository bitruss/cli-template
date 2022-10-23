package user_mgr

import (
	"time"

	"github.com/coreservice-io/cli-template/plugin/redis_plugin"
	"github.com/coreservice-io/cli-template/plugin/sqldb_plugin"
	"github.com/coreservice-io/cli-template/src/common/smart_cache"
)

// example for GormDB and tools cache
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

// query
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

	/////
	resultHolderAlloc := func() interface{} {
		return &QueryUserResult{
			Users:       []*ExampleUserModel{},
			Total_count: 0,
		}
	}

	/////
	query := func(resultHolder interface{}) error {
		queryResult := resultHolder.(*QueryUserResult)

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
		return query.Find(&queryResult.Users).Error
	}

	/////
	sq_result, sq_err := smart_cache.SmartQuery(key, resultHolderAlloc, fromCache, updateCache, 300, query, "UserQuery")

	/////
	if sq_err != nil {
		return nil, sq_err
	} else {
		return sq_result.(*QueryUserResult), nil
	}

}
