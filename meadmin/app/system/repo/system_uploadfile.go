package repo

import (
	"meadmin/app/system/model"
	"meadmin/system/datasource"
)

type SystemUploadfile struct {
	Repository[model.SystemUploadfile]
}

func NewSystemUploadfile() *SystemUploadfile {
	repo := &SystemUploadfile{}
	repo.SetDB(datasource.GetDB())
	return repo
}
