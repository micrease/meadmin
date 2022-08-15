package service

import (
	"github.com/jinzhu/copier"
	"github.com/micrease/gorme"
	"go.uber.org/zap"
	"meadmin/app/giftcard/dto"
	"meadmin/app/giftcard/enum"
	"meadmin/app/giftcard/model"
	"meadmin/app/giftcard/repo"
	"meadmin/library/context/api"
	"meadmin/library/context/result"
	"time"
)

type Merchant struct {
	Service
	repo *repo.Merchant
}

func NewMerchant(ctx *api.Context) *Merchant {
	service := Merchant{}
	service.Init(ctx)
	service.repo = repo.NewMerchant()
	return &service
}

func (this *Merchant) PageList(req dto.MerchantPageListReq) (*gorme.PageResult[model.Merchant], error) {
	pageList, err := this.repo.Paginate(req.PageNo, req.PageSize)
	return pageList, err
}

func (this *Merchant) Detail(id uint) (model.Merchant, error) {
	model, err := this.repo.FindById(id)
	return model, err
}

func (this *Merchant) ChangeStatus(req dto.MerchantStatusReq) *result.Error {
	model, err := this.repo.FindById(req.ID)
	if err != nil {
		return this.Error(err)
	}

	if model.ID == 0 {
		return this.Message("商户不存在")
	}

	if req.Status != enum.StatusEnable {
		req.Status = enum.StatusDisable
	}

	err = this.repo.Select("status").Save(model).Error
	if err != nil {
		return this.Error(err)
	}
	return nil
}

func (this *Merchant) Save(req dto.MerchantSaveReq) *result.Error {
	this.Log.Info("merchant.save", zap.Any("MerchantSaveReq", req))
	model := this.repo.NewModel()
	err := copier.Copy(&model, &req)
	if err != nil {
		return this.Error(err)
	}

	model.Status = enum.StatusEnable
	model.LoginTime = time.Now()
	if req.ID == 0 {
		model.CreatedBy = this.ctx.JwtClaimData.UserId
	} else {
		old, err := this.repo.FindById(req.ID)
		if err != nil {
			return this.Error(err)
		}
		model.UpdatedBy = this.ctx.JwtClaimData.UserId
		model.UpdatedAt = time.Now()
		model.CreatedAt = old.CreatedAt
		model.CreatedBy = old.CreatedBy
	}

	err = this.repo.Save(&model).Error
	if err != nil {
		return this.Error(err)
	}
	return nil
}
