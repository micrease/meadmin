package repo

import (
	models "admease/application/model"
	"admease/system/datasource"
)

type SystemUploadfile struct {
	Repository[models.SystemUploadfile]
}

func NewSystemUploadfile() *SystemUploadfile {
	repo := &SystemUploadfile{}
	repo.SetDB(datasource.GetDB())
	return repo
}
