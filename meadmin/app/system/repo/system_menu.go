package repo

import (
	"meadmin/app/system/model"
	"meadmin/system/datasource"
)

type SystemMenu struct {
	Repository[model.SystemMenu]
}

func NewSystemMenu() *SystemMenu {
	repo := &SystemMenu{}
	repo.SetDB(datasource.GetDB())
	return repo
}
