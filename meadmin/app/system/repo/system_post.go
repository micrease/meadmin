package repo

import (
	"meadmin/app/system/model"
	"meadmin/system/datasource"
)

type SystemPost struct {
	Repository[model.SystemPost]
}

func NewSystemPost() *SystemPost {
	repo := &SystemPost{}
	repo.SetDB(datasource.GetDB())
	return repo
}
