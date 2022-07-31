package service

import (
	"admease/app/system/dto"
	"admease/app/system/model"
	"admease/app/system/repo"
	"admease/library/context/api"
	"github.com/spf13/cast"
)

type SystemPost struct {
	repo *repo.SystemPost
}

func NewSystemPost() SystemPost {
	service := SystemPost{}
	service.repo = repo.NewSystemPost()
	return service
}

func (this SystemPost) PageList() (*dto.SystemPageListResp[model.SystemPost], error) {
	postPageList, err := this.repo.Where("status=?", model.StatusEnable).Paginate(1, 10)
	if err != nil {
		return nil, err
	}
	respPage := ToSystemPage[model.SystemPost](postPageList)
	return &respPage, nil
}

func (this SystemPost) PostList() ([]model.SystemPost, error) {
	postList, err := this.repo.Where("status=?", model.StatusEnable).List()
	return postList, err
}

func (this SystemPost) Save(ctx *api.Context, req dto.SystemPostSaveReq) error {
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
	return this.repo.Save(&model).Error
}
