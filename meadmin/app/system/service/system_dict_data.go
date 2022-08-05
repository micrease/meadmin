package service

import (
	"meadmin/app/system/dto"
	"meadmin/app/system/model"
	"meadmin/app/system/repo"
	"meadmin/library/context/api"
)

type SystemDictData struct {
	repo *repo.SystemDictData
}

func NewSystemDictData() SystemDictData {
	service := SystemDictData{}
	service.repo = repo.NewSystemDictData()
	return service
}

//查询列表
func (this *SystemDictData) List(ctx *api.Context, req dto.SystemDictDataListReq) ([]model.SystemDictData, error) {
	list, err := this.repo.Where("code", req.Code).List()
	return list, err
}
