package repo

import (
	models "admease/application/model"
	"admease/system/datasource"
)

type SystemUser struct {
	Repository[models.SystemUser]
}

func NewSystemUser() *SystemUser {
	repo := &SystemUser{}
	repo.SetDB(datasource.GetDB())
	return repo
}
