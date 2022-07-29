package repo

import (
	models "giftcard/application/model"
	"giftcard/system/datasource"
)

type SystemConfig struct {
	Repository[models.SettingConfig]
}

func NewSystemConfig() *SystemConfig {
	repo := &SystemConfig{}
	repo.SetDB(datasource.GetDB())
	return repo
}
