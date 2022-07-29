package model

type SystemAppApi struct {
	ID    uint64 `json:"id"`     // 主键
	AppId uint64 `json:"app_id"` // 应用ID
	ApiId uint64 `json:"api_id"` // API—ID
}

func (model SystemAppApi) GetID() uint64 {
	return model.ID
}

func (model SystemAppApi) TableName() string {
	return "system_app_api"
}
