package repo

import (
	models "giftcard/application/model"
	"giftcard/system/datasource"
)

type SystemUser struct {
	Repository[models.SystemUser]
}

func NewSystemUser() *SystemUser {
	repo := &SystemUser{}
	repo.SetDB(datasource.GetDB())
	return repo
}
