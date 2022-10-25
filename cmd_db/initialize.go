package cmd_db

import (
	"github.com/coreservice-io/cli-template/basic"
	"github.com/coreservice-io/cli-template/plugin/sqldb_plugin"
	"github.com/coreservice-io/cli-template/src/common/dbkv"
	"github.com/coreservice-io/cli-template/src/user_mgr"
	"gorm.io/gorm"
)

// =====below data can be changed=====
var ini_admin_email = "admin@coreservice.com"
var ini_admin_password = "to_be_reset"
var ini_admin_roles = user_mgr.UserRoles
var ini_admin_permissions = user_mgr.UserPermissions

func Initialize() {
	StartDBComponent()

	key := "db_initialized"
	result, _ := dbkv.GetDBKV(sqldb_plugin.GetInstance(), nil, &key, false, false)
	if result != nil {
		initialized, _ := result.ToBool()
		if initialized {
			basic.Logger.Infoln("db already initialized")
			return
		}
	}

	err := sqldb_plugin.GetInstance().Transaction(func(tx *gorm.DB) error {
		// create your own data here which won't change in the future
		reconfigAdmin(tx)

		// dbkv
		return dbkv.SetDBKV_Bool(tx, key, true, "db initialized sign")
	})
	if err != nil {
		basic.Logger.Errorln("db initialize error:", err)
		return
	}

	basic.Logger.Infoln("db initialized")
}

func reconfigAdmin(tx *gorm.DB) {
	default_u_admin, err := user_mgr.CreateUser(tx, ini_admin_email, ini_admin_password,
		user_mgr.RolesToStr(ini_admin_roles...), user_mgr.PermissionsToStr(ini_admin_permissions...))
	if err != nil {
		basic.Logger.Panicln(err)
	} else {
		basic.Logger.Infoln("default admin created:", default_u_admin)
	}
}
