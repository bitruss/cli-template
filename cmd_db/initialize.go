package cmd_db

import (
	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/plugin/sqldb_plugin"
	"github.com/coreservice-io/cli-template/src/user_mgr"
)

// =====below data can be changed=====
var ini_admin_email = "admin@coreservice.com"
var ini_admin_password = "to_be_reset"
var ini_admin_roles = user_mgr.UserRoles
var ini_admin_permissions = user_mgr.UserPermissions

func Initialize() {
	StartDBComponent()
	//create your own data here which won't change in the future
	reconfigAdmin()

	//dbkv
}

func reconfigAdmin() {
	default_u_admin, err := user_mgr.CreateUser(sqldb_plugin.GetInstance(), ini_admin_email, ini_admin_password,
		user_mgr.RolesToStr(ini_admin_roles...), user_mgr.PermissionsToStr(ini_admin_permissions...))
	if err != nil {
		basic.Logger.Panicln(err)
	} else {
		basic.Logger.Infoln("default admin created:", default_u_admin)
	}
}
