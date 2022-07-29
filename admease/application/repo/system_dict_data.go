package repo

import (
	models "giftcard/application/model"
	"giftcard/system/datasource"
)

type SystemDictData struct {
	Repository[models.SystemDictData]
}

func NewSystemDictData() *SystemDictData {
	repo := &SystemDictData{}
	repo.SetDB(datasource.GetDB())
	return repo
}
