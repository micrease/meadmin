package dto

import (
	"meadmin/app/system/dto"
	"time"
)

type CardCate struct {
	ID       int    `json:"id"`
	CateName string `json:"cate_name"`
}

type CardPrefix struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Currency struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type CardAttrs struct {
	CateList    []CardCate   `json:"cate_list"`
	Currencies  []Currency   `json:"currencies"`
	CardPrefixs []CardPrefix `json:"card_prefixs"`
}

type GiftCardSaveReq struct {
	ID           uint      `json:"id"`
	CardNo       string    `json:"card_no"`
	CardPrefix   string    `json:"card_prefix"`
	CateId       int       `json:"cate_id"`
	Currency     string    `json:"currency"`
	CVV          string    `json:"cvv"`
	Email        string    `json:"email"`
	FundPassword string    `json:"fund_password"`
	Password     string    `json:"password"`
	Phone        string    `json:"phone"`
	Remark       string    `json:"remark"`      //备注
	ExpireTime   time.Time `json:"expire_time"` // 过期时间

}

type GiftCardPageListReq struct {
	dto.PageQuery
	ID         uint64 `json:"id"`
	Username   string `json:"username"`
	CardNo     string `json:"card_no"`
	CardPrefix string `json:"card_prefix"`
	CateId     int    `json:"cate_id"`
	Currency   string `json:"currency"`
	Phone      string `json:"phone"`
	Status     int    `json:"status"`
}

type GiftCardStatusReq struct {
	ID     uint `json:"id"`
	Status int  `json:"status"`
}
