package models

import "time"

type SystemUploadfile struct {
	ID          uint64    `json:"id"`           // 主键
	StorageMode string    `json:"storage_mode"` // 状态 (1 本地 2 阿里云 3 七牛云 4 腾讯云)
	OriginName  string    `json:"origin_name"`  // 原文件名
	ObjectName  string    `json:"object_name"`  // 新文件名
	MimeType    string    `json:"mime_type"`    // 资源类型
	StoragePath string    `json:"storage_path"` // 存储目录
	Suffix      string    `json:"suffix"`       // 文件后缀
	SizeByte    uint64    `json:"size_byte"`    // 字节数
	SizeInfo    string    `json:"size_info"`    // 文件大小
	Url         string    `json:"url"`          // url地址
	CreatedBy   uint64    `json:"created_by"`   // 创建者
	UpdatedBy   uint64    `json:"updated_by"`   // 更新者
	CreatedAt   time.Time `json:"created_at"`   // 创建时间
	UpdatedAt   time.Time `json:"updated_at"`   // 更新时间
	Remark      string    `json:"remark"`       // 备注
}

func (model SystemUploadfile) GetID() uint64 {
	return model.ID
}

func (model SystemUploadfile) TableName() string {
	return "system_uploadfile"
}
