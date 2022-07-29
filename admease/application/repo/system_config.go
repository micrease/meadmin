package repo

import (
	models "admease/application/model"
	"admease/system/datasource"
)

type SystemConfig struct {
	Repository[models.SettingConfig]
}

func NewSystemConfig() *SystemConfig {
	repo := &SystemConfig{}
	repo.SetDB(datasource.GetDB())
	return repo
}
