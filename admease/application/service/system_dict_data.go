package service

import (
	"giftcard/application/dto"
	models "giftcard/application/model"
	"giftcard/application/repo"
	"giftcard/library/context/api"
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
func (this *SystemDictData) List(ctx *api.Context, req dto.SystemDictDataListReq) ([]models.SystemDictData, error) {
	list, err := this.repo.Where("code", req.Code).List()
	return list, err
}
