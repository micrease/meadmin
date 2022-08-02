package service

import (
	dto2 "admease/app/system/dto"
	"admease/app/system/model"
	repo2 "admease/app/system/repo"
	"admease/library/context/api"
	"admease/library/context/result"
	"admease/library/validate"
	"admease/system/config"
	"admease/system/consts"
	"admease/system/middleware"
	"crypto/md5"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/spf13/cast"
	"time"
)

type SystemUser struct {
	repo *repo2.SystemUser
}

func NewSystemUser() SystemUser {
	service := SystemUser{}
	service.repo = repo2.NewSystemUser()
	return service
}

func (u SystemUser) Login(req dto2.SystemLoginReq) (dto2.SystemLoginResp, *result.ErrorMessage) {
	model, err := u.repo.Where("username=?", req.Username).First()
	resp := dto2.SystemLoginResp{}
	if err != nil {
		return resp, result.Error(err)
	}

	if model.ID == 0 {
		return resp, result.Message("帐号不存在")
	}

	if model.Status == "1" {
		return resp, result.Message("帐号已停用")
	}

	password := md5.Sum([]byte(req.Password))
	passwordStr := fmt.Sprintf("%x", password)
	if passwordStr != model.Password {
		return resp, result.Message("密码不正确" + passwordStr)
	}

	conf := config.GetConfig()
	jwtSecret := conf.JwtSecret
	jwtClaimData := consts.JwtClaimData{
		UserId:       model.ID,
		Username:     model.Username,
		IsSuperAdmin: model.ID == consts.SuperAdmin,
	}

	token, err := middleware.GenerateToken(jwtSecret, jwtClaimData)
	if err != nil {
		return resp, result.Error(err, "生成Token失败")
	}
	resp.Token = token
	return resp, nil
}

func (u SystemUser) GetIInfo(userId uint64, IsSuperAdmin bool) (dto2.SystemInfoResp, *result.ErrorMessage) {
	resp := dto2.SystemInfoResp{}

	user, err := u.repo.GetById(cast.ToUint(userId))
	if err != nil {
		result.Error(err)
	}

	resp.User = dto2.UserDTO{}
	err = copier.Copy(&resp.User, &user)
	if err != nil {
		result.Error(err)
	}

	userRoleRepo := repo2.NewSystemUserRole()
	builder := userRoleRepo.NewQueryBuilder().Where("user_id=?", userId)
	roleIds := repo2.NewSystemUserRole().Values(builder, "role_id")
	var roles []string
	if len(roleIds) > 0 {
		roleRepo := repo2.NewSystemRole()
		builder = roleRepo.NewQueryBuilder().Where("id in(?)", roleIds)
		codes := roleRepo.Values(builder, "code")
		for _, val := range codes {
			roles = append(roles, cast.ToString(val))
		}
	}

	resp.Roles = roles
	if IsSuperAdmin {
		resp.Codes = []string{"*"}
		resp.Routers, err = NewSystemMenu().GetSuperAdminRouters()
		if err != nil {
			result.Error(err)
		}
	} else {
		menuService := NewSystemMenu()
		menuIds := menuService.GetMenuIdsByRoleIds(roleIds)
		resp.Routers, err = menuService.GetRoutersByIds(menuIds)
		menuCodes := menuService.GetMenuCode(menuIds)
		var codes []string
		for _, val := range menuCodes {
			codes = append(codes, cast.ToString(val))
		}
		resp.Codes = codes
	}
	return resp, nil
}

//登出
func (u SystemUser) Logout(ctx *api.Context) error {

	return nil
}

func (u SystemUser) PageList(ctx *api.Context) (*dto2.SystemPageListResp[model.SystemUser], *result.ErrorMessage) {
	var req dto2.SystemUserListReq
	validate.BindWithPanic(ctx, &req)
	builder := u.repo.NewQueryBuilder()
	withBuilder := u.repo.QueryWithBuilder(builder)
	if req.UserName != "" {
		withBuilder.Where("username=?", req.UserName)
	}
	if req.NickName != "" {
		withBuilder.Where("nickname=?", req.NickName)
	}
	if req.Phone != "" {
		withBuilder.Where("phone=?", req.Phone)
	}
	if req.Email != "" {
		withBuilder.Where("email=?", req.Email)
	}
	if req.Status != "" {
		withBuilder.Where("status=?", req.Status)
	}

	if req.MinDate != "" {
		location, err := time.Parse("2006-01-02", req.MinDate)
		if err == nil {
			withBuilder.Where("created_at>?", location)
		}
	}
	if req.MaxDate != "" {
		location, err := time.Parse("2006-01-02", req.MaxDate)
		if err == nil {
			withBuilder.Where("created_at<?", location)
		}
	}

	pageList, err := withBuilder.Paginate(req.Page, req.PageSize)
	if err != nil {
		return nil, result.Error(err)
	}
	page := ToSystemPage[model.SystemUser](pageList)
	return &page, nil
}

func (u SystemUser) ReadInfoById(ctx *api.Context, userId uint64) (dto2.SystemUserInfoResp, *result.ErrorMessage) {
	resp := dto2.SystemUserInfoResp{}
	user, err := u.repo.GetById(cast.ToUint(userId))
	if err != nil {
		return resp, result.Error(err)
	}
	err = copier.Copy(&resp, &user)

	userRoleRepo := repo2.NewSystemUserRole()
	builder := userRoleRepo.NewQueryBuilder().Where("user_id=?", userId)
	roleIds := repo2.NewSystemUserRole().Values(builder, "role_id")

	roles, err := NewSystemRole().GetRoutersByIds(roleIds)
	if err != nil {
		return resp, result.Error(err)
	}
	resp.RoleList = roles
	resp.PostList = make([]int, 0)
	return resp, nil
}

// ChangeStatus 设置用户状态
func (u SystemUser) ChangeStatus(userId uint64, status string) *result.ErrorMessage {
	builder := u.repo.NewQueryBuilder().Where("id=?", userId)
	err := builder.UpdateColumn("status", status).Error
	if err != nil {
		return result.Error(err)
	}
	return nil
}
