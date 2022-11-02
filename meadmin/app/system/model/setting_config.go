package model

// 参数配置信息表
type SettingConfig struct {
	ID        uint64 `json:"id"`         // 主键
	Key       string `json:"key"`        // 配置键名
	Value     string `json:"value"`      // 配置值
	Name      string `json:"name"`       // 配置名称
	GroupName string `json:"group_name"` // 组名称
	Sort      int    `json:"sort"`       // 排序
	Remark    string `json:"remark"`     // 备注
}

func (model SettingConfig) GetID() any {
	return model.ID
}

func (model SettingConfig) TableName() string {
	return "setting_config"
}
