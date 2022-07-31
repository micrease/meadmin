package repo

import (
	"admease/app/system/model"
	"admease/system/datasource"
)

type SystemUser struct {
	Repository[model.SystemUser]
}

func NewSystemUser() *SystemUser {
	repo := &SystemUser{}
	repo.SetDB(datasource.GetDB())
	return repo
}
