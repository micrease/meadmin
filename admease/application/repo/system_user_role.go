package repo

import (
	models "admease/application/model"
	"admease/system/datasource"
)

type SystemUserRole struct {
	Repository[models.SystemUserRole]
}

func NewSystemUserRole() *SystemUserRole {
	repo := &SystemUserRole{}
	repo.SetDB(datasource.GetDB())
	return repo
}
