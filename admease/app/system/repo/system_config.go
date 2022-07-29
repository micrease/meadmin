package repo

import (
	"admease/app/system/model"
	"admease/system/datasource"
)

type SystemConfig struct {
	Repository[model.SettingConfig]
}

func NewSystemConfig() *SystemConfig {
	repo := &SystemConfig{}
	repo.SetDB(datasource.GetDB())
	return repo
}
