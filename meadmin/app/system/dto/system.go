package dto

import "time"

type SystemConfigKeyReq struct {
	Key string `form:"key" json:"key"`
}

type SystemLoginReq struct {
	Code     string `form:"code" json:"code"`         //验证码
	Username string `form:"username" json:"username"` //用户名称
	Password string `form:"password" json:"password"` //密码
}

type SystemLoginResp struct {
	Token string `json:"token"` //token
}

type User struct {
	ID             uint      `json:"id"`
	Avatar         string    `json:"avatar"`
	BackendSetting string    `json:"backend_setting"`
	Username       string    `json:"username"`
	Nickname       string    `json:"nickname"`
	UserType       string    `json:"user_type"`
	Dashboard      string    `json:"dashboard"`
	DeptId         uint      `json:"dept_id"`
	Email          string    `json:"email"`
	LoginIp        string    `json:"login_ip"`
	LoginTime      string    `json:"login_time"`
	Phone          string    `json:"phone"`
	Remark         string    `json:"remark"` //备注
	Signed         string    `json:"signed"` //个人签名
	Status         int       `json:"status"`
	CreatedBy      uint      `json:"created_by"` //创建人ID
	CreatedAt      time.Time `json:"created_at"`
	UpdatedBy      uint      `json:"updated_by"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type RouterMeta struct {
	Hidden           bool   `json:"hidden"`
	HiddenBreadcrumb bool   `json:"hiddenBreadcrumb"`
	Icon             string `json:"icon"`
	Title            string `json:"title"`
	Type             string `json:"type"`
}

type RouterTree struct {
	ID        uint         `json:"id"`
	ParentId  uint         `json:"parent_id"`
	Path      string       `json:"path"`
	Component string       `json:"component"`
	Name      string       `json:"name"`
	Redirect  string       `json:"redirect"`
	Meta      RouterMeta   `json:"meta"`
	Children  []RouterTree `json:"children"`
}

type MenuTree struct {
	ID        uint64     `json:"id"`         // 主键
	ParentId  uint64     `json:"parent_id"`  // 父ID
	Level     string     `json:"level"`      // 组级集合
	Name      string     `json:"name"`       // 菜单名称
	Code      string     `json:"code"`       // 菜单标识代码
	Icon      string     `json:"icon"`       // 菜单图标
	Route     string     `json:"route"`      // 路由地址
	Component string     `json:"component"`  // 组件路径
	Redirect  string     `json:"redirect"`   // 跳转地址
	IsHidden  int        `json:"is_hidden"`  // 是否隐藏 (0是 1否)
	Type      string     `json:"type"`       // 菜单类型, (M菜单 B按钮 L链接 I iframe)
	Status    string     `json:"status"`     // 状态 (0正常 1停用)
	Sort      int        `json:"sort"`       // 排序
	CreatedBy uint64     `json:"created_by"` // 创建者
	UpdatedBy uint64     `json:"updated_by"` // 更新者
	CreatedAt time.Time  `json:"created_at"` // 创建时间
	UpdatedAt time.Time  `json:"updated_at"` // 更新时间
	DeletedAt time.Time  `json:"deleted_at"` // 删除时间
	Remark    string     `json:"remark"`     // 备注
	Children  []MenuTree `json:"children"`
}

type DeptTree struct {
	ID       uint64     `json:"id"`
	Label    string     `json:"label"`
	ParentId uint64     `json:"parent_id"`
	Value    uint64     `json:"value"`
	Children []DeptTree `json:"children"`
}

type DeptModelTree struct {
	ID        uint64          `json:"id"`         // 主键
	ParentId  uint64          `json:"parent_id"`  // 父ID
	Level     string          `json:"level"`      // 组级集合
	Name      string          `json:"name"`       // 部门名称
	Leader    string          `json:"leader"`     // 负责人
	Phone     string          `json:"phone"`      // 联系电话
	Status    string          `json:"status"`     // 状态 (0正常 1停用)
	Sort      int             `json:"sort"`       // 排序
	CreatedBy uint64          `json:"created_by"` // 创建者
	UpdatedBy uint64          `json:"updated_by"` // 更新者
	CreatedAt time.Time       `json:"created_at"` // 创建时间
	UpdatedAt time.Time       `json:"updated_at"` // 更新时间
	Remark    string          `json:"remark"`     // 备注
	Children  []DeptModelTree `json:"children"`
}

type SystemInfoResp struct {
	Roles   []string     `json:"roles"`
	Codes   []string     `json:"codes"`
	User    User         `json:"user"`
	Routers []RouterTree `json:"routers"`
}
