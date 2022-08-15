package model

import "time"

type Merchant struct {
	ID           uint64    `json:"id"`            // 用户ID，主键
	Username     string    `json:"username"`      // 用户名
	Password     string    `json:"password"`      // 密码
	MerchantName string    `json:"merchant_name"` // 商户名称
	Phone        string    `json:"phone"`         // 手机
	Email        string    `json:"email"`         // 用户邮箱
	Status       int       `json:"status"`        // 状态 (1正常 2停用)
	LoginIp      string    `json:"login_ip"`      // 最后登陆IP
	LoginTime    time.Time `json:"login_time"`    // 最后登陆时间
	CreatedBy    uint64    `json:"created_by"`    // 创建者
	UpdatedBy    uint64    `json:"updated_by"`    // 更新者
	CreatedAt    time.Time `json:"created_at"`    // 创建时间
	UpdatedAt    time.Time `json:"updated_at"`    // 更新时间
	Remark       string    `json:"remark"`        // 备注
}

func (model Merchant) GetID() uint64 {
	return model.ID
}

func (model Merchant) TableName() string {
	return "merchant"
}
