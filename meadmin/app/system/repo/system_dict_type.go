package repo

import (
	"meadmin/app/system/model"
)

type SystemDictType struct {
	Repository[model.SystemDictType]
}

func NewSystemDictType() *SystemDictType {
	repo := &SystemDictType{}
	repo.initialize()
	return repo
}
