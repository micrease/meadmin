package handler

import (
	"meadmin/app/system/dto"
	"meadmin/app/system/repo"
	"meadmin/library/context/api"
	"meadmin/library/context/result"
	"meadmin/library/validate"
)

//系统管理
type SystemConfig struct {
	Handler
}

//获取系统配置信息
func (this *SystemConfig) GetSysConfig(ctx *api.Context) *result.Result {
	list, _ := repo.NewSystemConfig().List()
	return result.Success(list)
}

//根据Key获取配置
func (this *SystemConfig) GetConfigByKey(ctx *api.Context) *result.Result {
	var req dto.SystemConfigKeyReq
	validate.BindWithPanic(ctx, &req)
	model, _ := repo.NewSystemConfig().Where("`key`=?", req.Key).First()
	return result.Success(model)
}
