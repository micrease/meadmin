package service

import (
	"giftcard/application/dto"
	models "giftcard/application/model"
	"giftcard/application/repo"
)

type SystemPost struct {
	repo *repo.SystemPost
}

func NewSystemPost() SystemPost {
	service := SystemPost{}
	service.repo = repo.NewSystemPost()
	return service
}

func (this SystemPost) PageList() (*dto.SystemPageListResp[models.SystemPost], error) {
	postPageList, err := this.repo.Where("status=?", models.StatusEnable).Paginate(1, 10)
	if err != nil {
		return nil, err
	}
	respPage := ToSystemPage[models.SystemPost](postPageList)
	return &respPage, nil
}
