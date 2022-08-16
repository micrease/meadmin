package model

import "time"

type GiftCard struct {
	ID           uint64    `json:"id"`
	CardNo       string    `json:"card_no"`
	CardPrefix   string    `json:"card_prefix"`
	CateId       int       `json:"cate_id"`
	Status       int       `json:"status"` // 状态 (1正常 2停用)
	Currency     string    `json:"currency"`
	CVV          string    `json:"cvv"`
	Email        string    `json:"email"`
	FundPassword string    `json:"fund_password"`
	Password     string    `json:"password"`
	Phone        string    `json:"phone"`
	Remark       string    `json:"remark"`     //备注
	CreatedBy    uint64    `json:"created_by"` // 创建者
	UpdatedBy    uint64    `json:"updated_by"` // 更新者
	CreatedAt    time.Time `json:"created_at"` // 创建时间
	UpdatedAt    time.Time `json:"updated_at"` // 更新时间
}

func (model GiftCard) GetID() uint64 {
	return model.ID
}

func (model GiftCard) TableName() string {
	return "giftcard"
}
