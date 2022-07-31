package repo

import (
	"admease/app/system/model"
	"admease/system/datasource"
)

type SystemDictType struct {
	Repository[model.SystemDictType]
}

func NewSystemDictType() *SystemDictType {
	repo := &SystemDictType{}
	repo.SetDB(datasource.GetDB())
	return repo
}
