package repo

import (
	models "giftcard/application/model"
	"giftcard/system/datasource"
)

type SystemUploadfile struct {
	Repository[models.SystemUploadfile]
}

func NewSystemUploadfile() *SystemUploadfile {
	repo := &SystemUploadfile{}
	repo.SetDB(datasource.GetDB())
	return repo
}
