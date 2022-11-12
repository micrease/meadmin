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
	if req.CardNo != "" {
		this.repo.Where("card_no", req.CardNo)
	}
	if req.Currency != "" {
		this.repo.Where("currency", req.Currency)
	}
	if req.Status != 0 {
		this.repo.Where("status", req.Status)
	}
	if req.MinDate != "" {
		location, err := time.Parse("2006-01-02", req.MinDate)
		if err == nil {
			this.repo.Where("created_at", ">", location)
		}
	}
	if req.MaxDate != "" {
		location, err := time.Parse("2006-01-02", req.MaxDate)
		if err == nil {
			this.repo.Where("created_at", "<", location)
		}
	}

	this.repo.Where("deleted_at is null")
	if req.OrderBy != "" {
		OrderType := "DESC"
		if req.OrderType == "ascending" {
			OrderType = "ASC"
		}
		this.repo.Order(req.OrderBy + " " + OrderType)
	}
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

	model.Status = req.Status
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
