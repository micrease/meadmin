package repo

import (
	"meadmin/app/system/model"
)

type SystemConfig struct {
	Repository[model.SettingConfig]
}

func NewSystemConfig() *SystemConfig {
	repo := &SystemConfig{}
	repo.initialize()
	return repo
}
