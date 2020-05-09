package action

import (
	_ "github.com/go-sql-driver/mysql"
)

func RegisterAcsModels() {
	registerMenuModels()
	registerApplicationModels()
	registerPermissionModels()
	registerRoleModels()
	registerUserModels()
	registerGroupModels()
}
