package repo

import (
	models "admease/application/model"
	"admease/system/datasource"
)

type SystemPost struct {
	Repository[models.SystemPost]
}

func NewSystemPost() *SystemPost {
	repo := &SystemPost{}
	repo.SetDB(datasource.GetDB())
	return repo
}
