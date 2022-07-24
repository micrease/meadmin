package repo

import (
	models "giftcard/application/model"
	"giftcard/system/datasource"
)

type SystemPost struct {
	Repository[models.SystemPost]
}

func NewSystemPost() *SystemPost {
	repo := &SystemPost{}
	repo.SetDB(datasource.GetDB())
	return repo
}
