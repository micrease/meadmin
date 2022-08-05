package repo

import (
	"meadmin/app/system/model"
	"meadmin/system/datasource"
)

type SystemUserRole struct {
	Repository[model.SystemUserRole]
}

func NewSystemUserRole() *SystemUserRole {
	repo := &SystemUserRole{}
	repo.SetDB(datasource.GetDB())
	return repo
}
