package repo

import (
	"meadmin/app/giftcard/model"
)

type Merchant struct {
	Repository[model.Merchant]
}

func NewMerchant() *Merchant {
	repo := &Merchant{}
	repo.initialize()
	return repo
}
