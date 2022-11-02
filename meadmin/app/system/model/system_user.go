package model

import "time"

type SystemUser struct {
	ID             uint64    `json:"id"`              // 用户ID，主键
	Username       string    `json:"username"`        // 用户名
	Password       string    `json:"password"`        // 密码
	UserType       string    `json:"user_type"`       // 用户类型：(100系统用户)
	Nickname       string    `json:"nickname"`        // 用户昵称
	Phone          string    `json:"phone"`           // 手机
	Email          string    `json:"email"`           // 用户邮箱
	Avatar         string    `json:"avatar"`          // 用户头像
	Signed         string    `json:"signed"`          // 个人签名
	Dashboard      string    `json:"dashboard"`       // 后台首页类型
	DeptId         uint64    `json:"dept_id"`         // 部门ID
	Status         string    `json:"status"`          // 状态 (0正常 1停用)
	LoginIp        string    `json:"login_ip"`        // 最后登陆IP
	LoginTime      time.Time `json:"login_time"`      // 最后登陆时间
	BackendSetting string    `json:"backend_setting"` // 后台设置数据
	CreatedBy      uint64    `json:"created_by"`      // 创建者
	UpdatedBy      uint64    `json:"updated_by"`      // 更新者
	CreatedAt      time.Time `json:"created_at"`      // 创建时间
	UpdatedAt      time.Time `json:"updated_at"`      // 更新时间
	DeletedAt      time.Time `json:"deleted_at"`      // 删除时间
	Remark         string    `json:"remark"`          // 备注
}

func (model SystemUser) GetID() any {
	return model.ID
}

func (model SystemUser) TableName() string {
	return "system_user"
}
