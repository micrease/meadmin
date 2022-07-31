package repo

import (
	"admease/app/system/model"
	"admease/system/datasource"
)

type SystemMenu struct {
	Repository[model.SystemMenu]
}

func NewSystemMenu() *SystemMenu {
	repo := &SystemMenu{}
	repo.SetDB(datasource.GetDB())
	return repo
}
