package repo

import (
	"meadmin/app/giftcard/model"
	"meadmin/system/datasource"
)

type Merchant struct {
	Repository[model.Merchant]
}

func NewMerchant() *Merchant {
	repo := &Merchant{}
	repo.SetDB(datasource.GetDB())
	return repo
}
