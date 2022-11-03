package repo

import (
	"meadmin/app/system/model"
)

type SystemUserPost struct {
	Repository[model.SystemUserPost]
}

func NewSystemUserPost() *SystemUserPost {
	repo := &SystemUserPost{}
	repo.initialize()
	return repo
}
