package db_cmd

import (
	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/plugin/sqldb_plugin"
	"github.com/coreservice-io/cli-template/src/user_mgr"
)

func Reconfig() {
	/////////////////////
	reconfigAdmin()
}

func reconfigAdmin() {

	var ini_admin_id int64 = 0

	//=====below data can be changed=====
	var ini_admin_email = "admin@coreservice.com"
	var ini_admin_password = "to_be_reset"
	var ini_admin_roles = user_mgr.UserRoles
	var ini_admin_permissions = user_mgr.UserPermissions
	//===================================

	admin_u, err := user_mgr.QueryUser(sqldb_plugin.GetInstance(), &ini_admin_id, nil, nil, nil, nil, 1, 0, false, true)
	if err != nil {
		basic.Logger.Panicln(err)
	}

	if admin_u.Total_count == 0 {
		default_u_admin, err := user_mgr.CreateUser(sqldb_plugin.GetInstance(), ini_admin_email, ini_admin_password,
			user_mgr.RolesToStr(ini_admin_roles...), user_mgr.PermissionsToStr(ini_admin_permissions...))
		if err != nil {
			basic.Logger.Panicln(err)
		} else {
			basic.Logger.Infoln("default admin created:", default_u_admin)
		}

	} else {
		user_mgr.UpdateUser(sqldb_plugin.GetInstance(), map[string]interface{}{
			"email":       ini_admin_email,
			"password":    ini_admin_password,
			"roles":       user_mgr.RolesToStr(ini_admin_roles...),
			"permissions": user_mgr.PermissionsToStr(ini_admin_permissions...),
		}, ini_admin_id)
	}
}
