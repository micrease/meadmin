package repo

import (
	"admease/app/system/model"
	"admease/system/datasource"
)

type SystemUserPost struct {
	Repository[model.SystemUserPost]
}

func NewSystemUserPost() *SystemUserPost {
	repo := &SystemUserPost{}
	repo.SetDB(datasource.GetDB())
	return repo
}
