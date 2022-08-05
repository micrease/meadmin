package repo

import (
	"meadmin/app/system/model"
	"meadmin/system/datasource"
)

type SystemConfig struct {
	Repository[model.SettingConfig]
}

func NewSystemConfig() *SystemConfig {
	repo := &SystemConfig{}
	repo.SetDB(datasource.GetDB())
	return repo
}
