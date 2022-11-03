package repo

import (
	"meadmin/app/system/model"
)

type SystemUploadfile struct {
	Repository[model.SystemUploadfile]
}

func NewSystemUploadfile() *SystemUploadfile {
	repo := &SystemUploadfile{}
	repo.initialize()
	return repo
}
