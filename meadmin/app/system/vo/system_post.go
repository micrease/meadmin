package vo

type SystemPostSaveReq struct {
	ID     uint64 `form:"id" json:"id"`
	Name   string `form:"name" json:"name"`
	Code   string `form:"code" json:"code"`
	Remark string `form:"remark" json:"remark"`
	Sort   int    `form:"sort" json:"sort"`
	Status string `form:"status" json:"status"`
}
