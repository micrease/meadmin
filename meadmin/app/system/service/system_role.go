package service

import (
	"github.com/spf13/cast"
	"meadmin/app/system/model"
	"meadmin/app/system/repo"
	"meadmin/app/system/vo"
	"meadmin/library/context/api"
	"meadmin/library/context/result"
)

type SystemRole struct {
	Service
	repo *repo.SystemRole
}

func NewSystemRole() *SystemRole {
	service := &SystemRole{}
	service.repo = repo.NewSystemRole()
	return service
}

func (this *SystemRole) Save(ctx *api.Context, req vo.SystemRoleSaveReq) error {
	model := model.SystemRole{}
	var err error
	if req.ID > 0 {
		model, err = this.repo.GetById(cast.ToUint(req.ID))
		if err != nil {
			return err
		}
		model.UpdatedBy = ctx.JwtClaimData.UserId
	} else {
		model.CreatedBy = ctx.JwtClaimData.UserId
	}

	model.ID = req.ID
	model.Name = req.Name
	model.Sort = req.Sort
	model.Code = req.Code
	model.Remark = req.Remark
	model.Status = req.Status
	model.DataScope = "0"
	return this.repo.Save(&model).Error
}

func (this *SystemRole) GetPageList(ctx *api.Context) (*vo.SystemPageListResp[model.SystemRole], error) {
	pageList, err := this.repo.Paginate(1, 10)
	if err != nil {
		return nil, err
	}
	resp := ToSystemPage[model.SystemRole](pageList)
	return &resp, nil
}

func (this *SystemRole) GetList(ctx *api.Context) ([]model.SystemRole, error) {
	list, err := this.repo.List()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (this *SystemRole) GetRoutersByIds(roleIds []any) ([]model.SystemRole, error) {
	list, err := this.repo.Where("status=0").Where("id in(?)", roleIds).Order("sort desc").List()
	if err != nil {
		return nil, err
	}
	return list, nil
}

// ChangeStatus 设置用户状态
func (this *SystemRole) ChangeStatus(userId uint64, status string) *result.Error {
	builder := this.repo.Where("id=?", userId)
	err := builder.UpdateColumn("status", status).Error
	if err != nil {
		return this.Error(err)
	}
	return nil
}

func (this *SystemRole) Delete(ids []string) *result.Error {
	err := this.repo.DeleteSoft(&model.SystemUser{}, ids).Error
	if err != nil {
		return this.Error(err)
	}
	return nil
}
