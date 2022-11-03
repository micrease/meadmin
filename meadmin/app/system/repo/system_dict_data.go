package repo

import (
	"meadmin/app/system/model"
)

type SystemDictData struct {
	Repository[model.SystemDictData]
}

func NewSystemDictData() *SystemDictData {
	repo := &SystemDictData{}
	repo.initialize()
	return repo
}
