package repo

import (
	"meadmin/app/system/model"
	"meadmin/system/datasource"
)

type SystemUserPost struct {
	Repository[model.SystemUserPost]
}

func NewSystemUserPost() *SystemUserPost {
	repo := &SystemUserPost{}
	repo.SetDB(datasource.GetDB())
	return repo
}
