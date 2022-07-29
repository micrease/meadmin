package repo

import (
	"admease/app/system/model"
	"admease/system/datasource"
)

type SystemUserRole struct {
	Repository[model.SystemUserRole]
}

func NewSystemUserRole() *SystemUserRole {
	repo := &SystemUserRole{}
	repo.SetDB(datasource.GetDB())
	return repo
}
