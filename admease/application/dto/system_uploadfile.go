package dto

type UploadFileReq struct {
	Path  string `form:"path" json:"path"`
	Image string `form:"image" json:"image"`
}
