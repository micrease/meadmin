package repo

import (
	"admease/app/system/model"
	"admease/system/datasource"
)

type SystemRole struct {
	Repository[model.SystemRole]
}

func NewSystemRole() *SystemRole {
	repo := &SystemRole{}
	repo.SetDB(datasource.GetDB())
	return repo
}
