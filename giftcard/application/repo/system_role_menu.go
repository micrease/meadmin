package repo

import (
	models "giftcard/application/model"
	"giftcard/system/datasource"
)

type SystemRoleMenu struct {
	Repository[models.SystemRoleMenu]
}

func NewSystemRoleMenu() *SystemRoleMenu {
	repo := &SystemRoleMenu{}
	repo.SetDB(datasource.GetDB())
	return repo
}
