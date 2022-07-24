package repo

import (
	models "giftcard/application/model"
	"giftcard/system/datasource"
)

type SystemUserRole struct {
	Repository[models.SystemUserRole]
}

func NewSystemUserRole() *SystemUserRole {
	repo := &SystemUserRole{}
	repo.SetDB(datasource.GetDB())
	return repo
}
