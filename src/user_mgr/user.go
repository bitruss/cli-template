package user_mgr

import (
	"errors"

	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/plugin/redis_plugin"
	"github.com/coreservice-io/cli-template/plugin/sqldb_plugin"
	"github.com/coreservice-io/cli-template/src/common/data"
	"github.com/coreservice-io/cli-template/src/common/smart_cache"
	"github.com/coreservice-io/utils/hash_util"
	"github.com/coreservice-io/utils/rand_util"
)

func CreateUser(email string, passwd string, roles string, permissions string) (*UserModel, error) {
	sha256_passwd := hash_util.SHA256String(passwd)
	token := rand_util.GenRandStr(24)

	user := &UserModel{
		Email:       email,
		Password:    sha256_passwd,
		Token:       token,
		Roles:       roles,
		Permissions: permissions,
		Forbidden:   false,
	}
	if err := sqldb_plugin.GetInstance().Table("users").Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(updateData map[string]interface{}, id int64) error {
	queryResult, err := QueryUser(&id, nil, nil, nil, nil, 1, 0, false, false)
	if err != nil {
		basic.Logger.Errorln("UpdateUser queryUsers error:", err, "id:", id)
		return err
	}
	if len(queryResult.Users) == 0 {
		return errors.New("user not exist")
	}

	err = sqldb_plugin.GetInstance().Table("users").Where("id =?", id).Updates(updateData).Error
	if err != nil {
		return err
	}

	//update cache , for fast api middleware token auth
	QueryUser(nil, &queryResult.Users[0].Token, nil, nil, nil, 1, 0, false, true)
	return nil
}

type QueryUserResult struct {
	Users       []*UserModel
	Total_count int64
}

func QueryUser(id *int64, token *string, emailPattern *string, email *string, forbidden *bool, limit int, offset int, fromCache bool, updateCache bool) (*QueryUserResult, error) {

	if emailPattern != nil && email != nil {
		return &QueryUserResult{
			Users:       []*UserModel{},
			Total_count: 0,
		}, errors.New("emailPattern ,email :can't be set at the same time")
	}

	//gen_key
	ck := smart_cache.NewConnectKey("users")
	ck.C_Int64_Ptr("id", id).
		C_Str_Ptr("token", token).
		C_Str_Ptr("emailPattern", emailPattern).
		C_Str_Ptr("email", email).
		C_Bool_Ptr("forbidden", forbidden).
		C_Int(limit).
		C_Int(offset)

	key := redis_plugin.GetInstance().GenKey(ck.String())

	/////
	resultHolderAlloc := func() interface{} {
		return &QueryUserResult{
			Users:       []*UserModel{},
			Total_count: 0,
		}
	}

	/////
	query := func(resultHolder interface{}) error {
		queryResult := resultHolder.(*QueryUserResult)

		query := sqldb_plugin.GetInstance().Table("users")
		if emailPattern != nil {
			query.Where("email LIKE ?", "%"+*emailPattern+"%")
		}
		if email != nil {
			query.Where("email = ?", email)
		}

		if id != nil {
			query.Where("id = ?", *id)
		}

		if token != nil {
			query.Where("token = ?", *token)
		}
		if forbidden != nil {
			query.Where("forbidden = ?", *forbidden)
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
	sq_result, sq_err := smart_cache.SmartQuery(key, resultHolderAlloc, true, fromCache, updateCache, 300, query, "QueryUser")

	/////
	if sq_err != nil {
		return nil, sq_err
	} else {
		return sq_result.(*QueryUserResult), nil
	}
}

//return true if all element defined in array is a allowed permission definition
func CheckPermissionList(permissions []string) bool {
	for _, v := range permissions {
		if !data.InArray(v, UserPermissions) {
			return false
		}
	}
	return true
}

//return true if all element defined in array is a allowed role definition
func CheckRoleList(roles []string) bool {
	for _, v := range roles {
		if !data.InArray(v, UserRoles) {
			return false
		}
	}
	return true
}
