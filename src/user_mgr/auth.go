package user_mgr

import (
	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/src/common/data"
	"github.com/coreservice-io/cli-template/src/common/json"
)

const (
	USER_ROLE_ADMIN    = "admin"
	USER_ROLE_USER     = "user"
	USER_ROLE_READONLY = "read_only"
)

var UserRoles = []string{USER_ROLE_ADMIN, USER_ROLE_USER, USER_ROLE_READONLY}
var UserPermissions = []string{}

// only return true if  permissions are all defined
func PermissionsDefined(permissions []string) bool {
	for _, v := range permissions {
		if !data.InArray(v, UserPermissions) {
			return false
		}
	}
	return true
}

// only return true if  roles are all defined
func RolesDefined(roles []string) bool {
	for _, v := range roles {
		if !data.InArray(v, UserRoles) {
			return false
		}
	}
	return true
}

// user must have all the roles specified
func (u *UserModel) HasRoles(roles []string) bool {
	var userRoles []string
	err := json.Unmarshal([]byte(u.Roles), &userRoles)
	if err != nil {
		basic.Logger.Errorln("HasOneOfRoles json.Unmarshal error:", err)
		return false
	}

	for _, role := range roles {
		if !data.InArray(role, userRoles) {
			return false
		}
	}
	return true
}

func (u *UserModel) HasOneOfRoles(roles []string) bool {
	if len(roles) == 0 {
		return true
	}

	var userRoles []string
	err := json.Unmarshal([]byte(u.Roles), &userRoles)
	if err != nil {
		basic.Logger.Errorln("HasOneOfRoles json.Unmarshal error:", err)
		return false
	}

	for _, role := range roles {
		if data.InArray(role, userRoles) {
			return true
		}
	}
	return false
}

func (u *UserModel) HasOneOfPermissions(permissions []string) bool {
	if len(permissions) == 0 {
		return true
	}

	var userPermissions []string
	err := json.Unmarshal([]byte(u.Permissions), &userPermissions)
	if err != nil {
		basic.Logger.Errorln("HasOneOfPermissions json.Unmarshal error:", err)
		return false
	}

	for _, permission := range permissions {
		if data.InArray(permission, userPermissions) {
			return true
		}
	}
	return false
}
