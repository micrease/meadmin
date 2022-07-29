package repo

import (
	"admease/app/system/model"
	"admease/system/datasource"
)

type SystemDictData struct {
	Repository[model.SystemDictData]
}

func NewSystemDictData() *SystemDictData {
	repo := &SystemDictData{}
	repo.SetDB(datasource.GetDB())
	return repo
}
