package repo

import (
	"meadmin/app/giftcard/model"
)

type GiftCard struct {
	Repository[model.GiftCard]
}

func NewGiftcard() *GiftCard {
	repo := &GiftCard{}
	repo.initialize()
	return repo
}
