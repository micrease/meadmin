package repo

import (
	"meadmin/app/system/model"
	"meadmin/system/datasource"
)

type SystemDictData struct {
	Repository[model.SystemDictData]
}

func NewSystemDictData() *SystemDictData {
	repo := &SystemDictData{}
	repo.SetDB(datasource.GetDB())
	return repo
}
