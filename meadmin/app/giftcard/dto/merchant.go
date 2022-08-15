package dto

import "meadmin/app/system/dto"

type MerchantSaveReq struct {
	ID           uint   `json:"id"`
	Username     string `json:"username"`
	MerchantName string `json:"merchant_name"`
	Password     string `json:"password"`
	FundPassword string `json:"fund_password"`
	Email        string `json:"email"`
	LoginIp      string `json:"login_ip"`
	LoginTime    string `json:"login_time"`
	Phone        string `json:"phone"`
	Remark       string `json:"remark"` //备注
	Status       int    `json:"status"`
}

type MerchantStatusReq struct {
	ID     uint `json:"id"`
	Status int  `json:"status"`
}

type MerchantPageListReq struct {
	dto.PageQuery
	ID           uint   `json:"id"`
	Username     string `json:"username"`
	MerchantName string `json:"merchant_name"`
	Phone        string `json:"phone"`
	Status       int    `json:"status"`
}
