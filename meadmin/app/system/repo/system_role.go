package repo

import (
	"meadmin/app/system/model"
	"meadmin/system/datasource"
)

type SystemRole struct {
	Repository[model.SystemRole]
}

func NewSystemRole() *SystemRole {
	repo := &SystemRole{}
	repo.SetDB(datasource.GetDB())
	return repo
}
