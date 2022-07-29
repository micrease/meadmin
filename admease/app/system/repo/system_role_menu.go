package repo

import (
	"admease/app/system/model"
	"admease/system/datasource"
)

type SystemRoleMenu struct {
	Repository[model.SystemRoleMenu]
}

func NewSystemRoleMenu() *SystemRoleMenu {
	repo := &SystemRoleMenu{}
	repo.SetDB(datasource.GetDB())
	return repo
}
