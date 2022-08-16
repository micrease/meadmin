package repo

import (
	"meadmin/app/giftcard/model"
	"meadmin/system/datasource"
)

type GiftCard struct {
	Repository[model.GiftCard]
}

func NewGiftcard() *GiftCard {
	repo := &GiftCard{}
	repo.SetDB(datasource.GetDB())
	return repo
}
