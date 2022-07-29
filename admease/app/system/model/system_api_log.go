package model

import "time"

type SystemApiLog struct {
	ID           uint64    `json:"id"`            // 主键
	ApiId        uint64    `json:"api_id"`        // api ID
	ApiName      string    `json:"api_name"`      // 接口名称
	AccessName   string    `json:"access_name"`   // 接口访问名称
	RequestData  string    `json:"request_data"`  // 请求数据
	ResponseCode string    `json:"response_code"` // 响应状态码
	ResponseData string    `json:"response_data"` // 响应数据
	Ip           string    `json:"ip"`            // 访问IP地址
	IpLocation   string    `json:"ip_location"`   // IP所属地
	AccessTime   time.Time `json:"access_time"`   // 访问时间
	Remark       string    `json:"remark"`        // 备注
}

func (model SystemApiLog) GetID() uint64 {
	return model.ID
}

func (model SystemApiLog) TableName() string {
	return "system_api_log"
}
