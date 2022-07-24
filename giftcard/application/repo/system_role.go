package repo

import (
	models "giftcard/application/model"
	"giftcard/system/datasource"
)

type SystemRole struct {
	Repository[models.SystemRole]
}

func NewSystemRole() *SystemRole {
	repo := &SystemRole{}
	repo.SetDB(datasource.GetDB())
	return repo
}
