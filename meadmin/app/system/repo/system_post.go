package repo

import (
	"admease/app/system/model"
	"admease/system/datasource"
)

type SystemPost struct {
	Repository[model.SystemPost]
}

func NewSystemPost() *SystemPost {
	repo := &SystemPost{}
	repo.SetDB(datasource.GetDB())
	return repo
}
