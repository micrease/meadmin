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

type GiftCard struct {
	Service
	repo *repo.GiftCard
}

func NewGiftCard(ctx *api.Context) *GiftCard {
	service := GiftCard{}
	service.Init(ctx)
	service.repo = repo.NewGiftcard()
	return &service
}

func (this *GiftCard) PageList(req dto.GiftCardPageListReq) (*gorme.PageResult[model.GiftCard], error) {
	pageList, err := this.repo.Paginate(req.PageNo, req.PageSize)
	return pageList, err
}

func (this *GiftCard) Detail(id uint) (model.GiftCard, error) {
	model, err := this.repo.FindById(id)
	return model, err
}

func (this *GiftCard) ChangeStatus(req dto.GiftCardStatusReq) *result.Error {
	model, err := this.repo.FindById(req.ID)
	if err != nil {
		return this.Error(err)
	}

	if model.ID == 0 {
		return this.Message("礼品卡不存在")
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

func (this *GiftCard) Save(req dto.GiftCardSaveReq) *result.Error {
	this.Log.Info("GiftCard.save", zap.Any("GiftCardSaveReq", req))
	model := this.repo.NewModel()
	err := copier.Copy(&model, &req)
	if err != nil {
		return this.Error(err)
	}

	model.Status = enum.StatusEnable
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

func (g *GiftCard) Delete(ids []string) *result.Error {
	err := g.repo.NewQueryBuilder().Delete(&model.GiftCard{}, ids).Error
	if err != nil {
		return g.Error(err)
	}
	return nil
}