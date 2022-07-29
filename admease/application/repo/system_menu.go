package repo

import (
	models "giftcard/application/model"
	"giftcard/system/datasource"
)

type SystemMenu struct {
	Repository[models.SystemMenu]
}

func NewSystemMenu() *SystemMenu {
	repo := &SystemMenu{}
	repo.SetDB(datasource.GetDB())
	return repo
}
