package repo

import (
	"meadmin/app/system/model"
)

type SystemRoleMenu struct {
	Repository[model.SystemRoleMenu]
}

func NewSystemRoleMenu() *SystemRoleMenu {
	repo := &SystemRoleMenu{}
	repo.initialize()
	return repo
}
