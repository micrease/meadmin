package repo

import (
	models "admease/application/model"
	"admease/system/datasource"
)

type SystemRoleMenu struct {
	Repository[models.SystemRoleMenu]
}

func NewSystemRoleMenu() *SystemRoleMenu {
	repo := &SystemRoleMenu{}
	repo.SetDB(datasource.GetDB())
	return repo
}
