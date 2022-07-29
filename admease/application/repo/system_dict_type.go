package repo

import (
	models "admease/application/model"
	"admease/system/datasource"
)

type SystemDictType struct {
	Repository[models.SystemDictType]
}

func NewSystemDictType() *SystemDictType {
	repo := &SystemDictType{}
	repo.SetDB(datasource.GetDB())
	return repo
}
