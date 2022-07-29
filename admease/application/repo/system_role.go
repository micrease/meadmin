package repo

import (
	models "admease/application/model"
	"admease/system/datasource"
)

type SystemRole struct {
	Repository[models.SystemRole]
}

func NewSystemRole() *SystemRole {
	repo := &SystemRole{}
	repo.SetDB(datasource.GetDB())
	return repo
}
