package service

import (
	"admease/app/system/dto"
	"admease/app/system/model"
	"admease/app/system/repo"
	"admease/library/context/api"
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
