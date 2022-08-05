package service

import (
	"fmt"
	"github.com/spf13/cast"
	"meadmin/app/system/dto"
	"meadmin/app/system/model"
	"meadmin/app/system/repo"
	"meadmin/library/context/api"
	"meadmin/library/files"
	"meadmin/system/config"
	"mime/multipart"
	"strings"
	"time"
)

type SystemUploadfile struct {
	repo *repo.SystemUploadfile
}

func NewSystemUploadFile() SystemUploadfile {
	service := SystemUploadfile{}
	service.repo = repo.NewSystemUploadfile()
	return service
}

func (this SystemUploadfile) GetPageList(ctx *api.Context) (*dto.SystemPageListResp[model.SystemUploadfile], error) {
	pageList, err := this.repo.Paginate(1, 10)
	if err != nil {
		return nil, err
	}
	resp := ToSystemPage[model.SystemUploadfile](pageList)
	return &resp, nil
}

func (this SystemUploadfile) Uploadfile(ctx *api.Context, file *multipart.FileHeader) (model.SystemUploadfile, error) {

	model := this.repo.NewModel()
	model.MimeType = file.Header.Get("Content-Type")
	model.OriginName = file.Filename

	suffix := ""
	if len(file.Filename) > 0 {
		arr := strings.Split(file.Filename, ".")
		if len(arr) > 1 {
			suffix = arr[len(arr)-1]
		}
	}
	model.Suffix = suffix

	ts := time.Now().UnixMicro()
	name := fmt.Sprintf("%d.%s", ts, suffix)
	model.ObjectName = name
	model.SizeByte = cast.ToUint64(file.Size)

	storagePath := "/runtime/upload/img/"
	size := cast.ToFloat64(file.Size / 1000)
	unit := "KB"
	if size >= 1000 {
		size = size / 1000
		unit = "MB"
	}

	sizeInfo := fmt.Sprintf("%0.2f%s", size, unit)
	model.StoragePath = storagePath
	model.SizeInfo = sizeInfo
	model.StorageMode = "1"
	model.Url = "/static/img/" + model.ObjectName
	model.CreatedBy = ctx.JwtClaimData.UserId
	model.UpdatedBy = ctx.JwtClaimData.UserId
	model.CreatedAt = time.Now()
	model.UpdatedAt = time.Now()

	conf := config.GetConfig()
	err := files.MkdirIfNotExist(conf.ProjectPath + storagePath)
	if err != nil {
		return model, err
	}
	path := conf.ProjectPath + storagePath + model.ObjectName
	err = ctx.GinCtx.SaveUploadedFile(file, path)
	if err != nil {
		return model, err
	}

	err = this.repo.Save(&model).Error
	return model, err
}
