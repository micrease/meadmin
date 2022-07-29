package repo

import (
	models "giftcard/application/model"
	"giftcard/system/datasource"
)

type SystemDictType struct {
	Repository[models.SystemDictType]
}

func NewSystemDictType() *SystemDictType {
	repo := &SystemDictType{}
	repo.SetDB(datasource.GetDB())
	return repo
}
