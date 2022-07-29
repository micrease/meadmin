package admin

import (
	"admease/application/controller"
	"admease/application/dto"
	"admease/application/repo"
	"admease/library/context/api"
	"admease/library/context/result"
	"admease/library/validate"
)

//系统管理
type SystemConfig struct {
	controller.BaseController
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
