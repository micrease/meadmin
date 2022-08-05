package repo

import (
	"meadmin/app/system/model"
	"meadmin/system/datasource"
)

type SystemDictType struct {
	Repository[model.SystemDictType]
}

func NewSystemDictType() *SystemDictType {
	repo := &SystemDictType{}
	repo.SetDB(datasource.GetDB())
	return repo
}
