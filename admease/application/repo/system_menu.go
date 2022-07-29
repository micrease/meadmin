package repo

import (
	models "admease/application/model"
	"admease/system/datasource"
)

type SystemMenu struct {
	Repository[models.SystemMenu]
}

func NewSystemMenu() *SystemMenu {
	repo := &SystemMenu{}
	repo.SetDB(datasource.GetDB())
	return repo
}
