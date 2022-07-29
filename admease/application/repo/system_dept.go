package repo

import (
	models "admease/application/model"
	"admease/system/datasource"
)

type SystemDept struct {
	Repository[models.SystemDept]
}

func NewSystemDept() *SystemDept {
	repo := &SystemDept{}
	repo.SetDB(datasource.GetDB())
	return repo
}
