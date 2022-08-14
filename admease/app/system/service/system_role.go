package service

import (
	dto2 "admease/app/system/dto"
	"admease/app/system/model"
	"admease/app/system/repo"
	"admease/library/context/api"
	"admease/library/context/result"
	"admease/library/validate"
	"github.com/spf13/cast"
	"time"
)

type SystemRole struct {
	repo *repo.SystemRole
}

func NewSystemRole() SystemRole {
	service := SystemRole{}
	service.repo = repo.NewSystemRole()
	return service
}

func (this SystemRole) Save(ctx *api.Context, req dto2.SystemRoleSaveReq) error {
	model := this.repo.NewModel()
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

func (this SystemRole) GetPageList(ctx *api.Context) (*dto2.SystemPageListResp[model.SystemRole], error) {
	var req dto2.SystemRoleListReq
	validate.BindWithPanic(ctx, &req)
	builder := this.repo.NewQueryBuilder()
	withBuilder := this.repo.QueryWithBuilder(builder)
	if req.Name != "" {
		withBuilder.Where("name=?", req.Name)
	}

	if req.Code != "" {
		withBuilder.Where("code=?", req.Code)
	}
	if req.Status != "" {
		withBuilder.Where("status=?", req.Status)
	}

	if req.MinDate != "" {
		location, err := time.Parse("2006-01-02", req.MinDate)
		if err == nil {
			withBuilder.Where("created_at>?", location)
		}
	}
	if req.MaxDate != "" {
		location, err := time.Parse("2006-01-02", req.MaxDate)
		if err == nil {
			withBuilder.Where("created_at<?", location)
		}
	}
	withBuilder.Where("deleted_at is null")
	if req.OrderBy != "" {
		OrderType := "DESC"
		if req.OrderType == "ascending" {
			OrderType = "ASC"
		}
		withBuilder.Order(req.OrderBy + " " + OrderType)
	}

	pageList, err := this.repo.Paginate(1, 10)
	if err != nil {
		return nil, err
	}
	resp := ToSystemPage[model.SystemRole](pageList)
	return &resp, nil
}

func (this SystemRole) GetList(ctx *api.Context) ([]model.SystemRole, error) {
	list, err := this.repo.List()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (this SystemRole) GetRoutersByIds(roleIds []any) ([]model.SystemRole, error) {
	list, err := this.repo.Order("sort desc").
		Where("status=?", model.StatusEnable).
		Where("id in(?)", roleIds).
		List()
	if err != nil {
		return nil, err
	}
	return list, nil
}

// ChangeStatus 设置用户状态
func (r SystemRole) ChangeStatus(userId uint64, status string) *result.ErrorMessage {
	builder := r.repo.NewQueryBuilder().Where("id=?", userId)
	err := builder.UpdateColumn("status", status).Error
	if err != nil {
		return result.Error(err)
	}
	return nil
}

func (r SystemRole) Delete(ids []string) error {
	return r.repo.NewQueryBuilder().Delete(&model.SystemUser{}, ids).Error
}