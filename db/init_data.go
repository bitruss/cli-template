package db

import (
	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/src/user_mgr"
)

func InitData() {
	//create your own data here

	//init default admin
	ini_admin_email := "admin@admin.com"
	ini_admin_password := "adminxwiqweru"
	ini_admin_roles := user_mgr.UserRoles
	ini_admin_permissions := user_mgr.UserPermissions

	default_u_admin, err := user_mgr.CreateUser(ini_admin_email, ini_admin_password,
		user_mgr.RolesToStr(ini_admin_roles...), user_mgr.PermissionsToStr(ini_admin_permissions...))
	if err != nil {
		panic(err)
	} else {
		basic.Logger.Infoln("default admin created:", default_u_admin)
	}

}
