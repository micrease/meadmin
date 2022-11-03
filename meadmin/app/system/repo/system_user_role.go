package repo

import (
	"meadmin/app/system/model"
)

type SystemUserRole struct {
	Repository[model.SystemUserRole]
}

func NewSystemUserRole() *SystemUserRole {
	repo := &SystemUserRole{}
	repo.initialize()
	return repo
}
