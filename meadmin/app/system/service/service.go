package service

import (
	"github.com/micrease/gorme"
	"meadmin/app/system/dto"
)

func ToSystemPage[T any](result *gorme.PageResult[T]) dto.SystemPageListResp[T] {
	var t dto.SystemPageListResp[T]
	t.PageInfo = dto.PageInfo{
		CurrentPage: result.PageNo,
		TotalPage:   result.TotalPage,
		Total:       result.TotalSize,
	}
	t.Items = result.List
	return t
}
