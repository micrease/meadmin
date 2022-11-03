package repo

import (
	"meadmin/app/system/model"
)

type SystemUser struct {
	Repository[model.SystemUser]
}

func NewSystemUser() *SystemUser {
	repo := &SystemUser{}
	repo.initialize()
	return repo
}
