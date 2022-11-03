package repo

import (
	"meadmin/app/system/model"
)

type SystemMenu struct {
	Repository[model.SystemMenu]
}

func NewSystemMenu() *SystemMenu {
	repo := &SystemMenu{}
	repo.initialize()
	return repo
}
