package repo

import (
	models "giftcard/application/model"
	"giftcard/system/datasource"
)

type SystemDept struct {
	Repository[models.SystemDept]
}

func NewSystemDept() *SystemDept {
	repo := &SystemDept{}
	repo.SetDB(datasource.GetDB())
	return repo
}
