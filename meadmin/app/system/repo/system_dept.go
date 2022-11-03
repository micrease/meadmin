package repo

import (
	"meadmin/app/system/model"
)

type SystemDept struct {
	Repository[model.SystemDept]
}

func NewSystemDept() *SystemDept {
	repo := &SystemDept{}
	repo.initialize()
	return repo
}
