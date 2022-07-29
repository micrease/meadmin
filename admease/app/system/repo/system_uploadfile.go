package repo

import (
	"admease/app/system/model"
	"admease/system/datasource"
)

type SystemUploadfile struct {
	Repository[model.SystemUploadfile]
}

func NewSystemUploadfile() *SystemUploadfile {
	repo := &SystemUploadfile{}
	repo.SetDB(datasource.GetDB())
	return repo
}
