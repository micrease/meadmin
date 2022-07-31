package repo

import (
	"admease/app/system/model"
	"admease/system/datasource"
)

type SystemDept struct {
	Repository[model.SystemDept]
}

func NewSystemDept() *SystemDept {
	repo := &SystemDept{}
	repo.SetDB(datasource.GetDB())
	return repo
}
