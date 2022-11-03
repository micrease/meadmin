package repo

import (
	"meadmin/app/system/model"
)

type SystemPost struct {
	Repository[model.SystemPost]
}

func NewSystemPost() *SystemPost {
	repo := &SystemPost{}
	repo.initialize()
	return repo
}
