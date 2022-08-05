package repo

import (
	"meadmin/app/system/model"
	"meadmin/system/datasource"
)

type SystemDept struct {
	Repository[model.SystemDept]
}

func NewSystemDept() *SystemDept {
	repo := &SystemDept{}
	repo.SetDB(datasource.GetDB())
	return repo
}
