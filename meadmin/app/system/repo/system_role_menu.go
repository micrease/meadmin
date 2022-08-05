package repo

import (
	"meadmin/app/system/model"
	"meadmin/system/datasource"
)

type SystemRoleMenu struct {
	Repository[model.SystemRoleMenu]
}

func NewSystemRoleMenu() *SystemRoleMenu {
	repo := &SystemRoleMenu{}
	repo.SetDB(datasource.GetDB())
	return repo
}
