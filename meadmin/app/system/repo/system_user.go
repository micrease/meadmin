package repo

import (
	"meadmin/app/system/model"
	"meadmin/system/datasource"
)

type SystemUser struct {
	Repository[model.SystemUser]
}

func NewSystemUser() *SystemUser {
	repo := &SystemUser{}
	repo.SetDB(datasource.GetDB())
	return repo
}
