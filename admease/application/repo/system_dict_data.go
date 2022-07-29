package repo

import (
	models "admease/application/model"
	"admease/system/datasource"
)

type SystemDictData struct {
	Repository[models.SystemDictData]
}

func NewSystemDictData() *SystemDictData {
	repo := &SystemDictData{}
	repo.SetDB(datasource.GetDB())
	return repo
}
