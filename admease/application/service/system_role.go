package service

import (
	"giftcard/application/dto"
	models "giftcard/application/model"
	"giftcard/application/repo"
	"giftcard/library/context/api"
	"github.com/spf13/cast"
)

type SystemRole struct {
	repo *repo.SystemRole
}

func NewSystemRole() SystemRole {
	service := SystemRole{}
	service.repo = repo.NewSystemRole()
	return service
}

func (this SystemRole) Save(ctx *api.Context, req dto.SystemRoleSaveReq) error {
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

func (this SystemRole) GetPageList(ctx *api.Context) (*dto.SystemPageListResp[models.SystemRole], error) {
	pageList, err := this.repo.Paginate(1, 10)
	if err != nil {
		return nil, err
	}
	resp := ToSystemPage[models.SystemRole](pageList)
	return &resp, nil
}

func (this SystemRole) GetList(ctx *api.Context) ([]models.SystemRole, error) {
	list, err := this.repo.List()
	if err != nil {
		return nil, err
	}
	return list, nil
}
