package repo

import (
	"meadmin/app/system/model"
)

type SystemRole struct {
	Repository[model.SystemRole]
}

func NewSystemRole() *SystemRole {
	repo := &SystemRole{}
	repo.initialize()
	return repo
}
