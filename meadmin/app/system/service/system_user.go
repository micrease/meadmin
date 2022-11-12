package service

import (
	"crypto/md5"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/spf13/cast"
	"meadmin/app/system/dto"
	"meadmin/app/system/model"
	"meadmin/app/system/repo"
	"meadmin/library/context/api"
	"meadmin/library/context/result"
	"meadmin/library/validate"
	"meadmin/system/config"
	"meadmin/system/consts"
	"meadmin/system/middleware"
	"time"
)

type SystemUser struct {
	Service
	repo *repo.SystemUser
}

func NewSystemUser(ctx *api.Context) SystemUser {
	service := SystemUser{}
	service.Init(ctx)
	service.repo = repo.NewSystemUser()
	return service
}

func (this SystemUser) Login(req dto.SystemLoginReq) (dto.SystemLoginResp, *result.Error) {
	model, err := this.repo.Where("username=?", req.Username).First()
	resp := dto.SystemLoginResp{}
	if err != nil {
		return resp, this.Error(err)
	}

	if model.ID == 0 {
		return resp, this.Message("帐号不存在")
	}

	if model.Status == "1" {
		return resp, this.Message("帐号已停用")
	}

	password := md5.Sum([]byte(req.Password))
	passwordStr := fmt.Sprintf("%x", password)
	if passwordStr != model.Password {
		return resp, this.Message("密码不正确" + passwordStr)
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
		return resp, this.Error(err, "生成Token失败")
	}
	resp.Token = token
	return resp, nil
}

func (this SystemUser) GetInfo() (dto.SystemInfoResp, *result.Error) {
	resp := dto.SystemInfoResp{}

	userId := this.ctx.JwtClaimData.UserId
	IsSuperAdmin := this.ctx.JwtClaimData.IsSuperAdmin

	user, err := this.repo.GetById(cast.ToUint(userId))
	if err != nil {
		return resp, this.Error(err)
	}

	resp.User = dto.User{}
	err = copier.Copy(&resp.User, &user)
	if err != nil {
		return resp, this.Error(err)
	}

	roleIds, err := repo.NewSystemUserRole().Where("user_id", userId).Values("role_id")
	if err != nil {
		return resp, this.Error(err)
	}
	var roles []string
	if len(roleIds) > 0 {
		roleRepo := repo.NewSystemRole()
		codes, _ := roleRepo.Where("id in(?)", roleIds).Values("code")
		for _, val := range codes {
			roles = append(roles, cast.ToString(val))
		}
	}
	resp.Roles = roles
	if IsSuperAdmin {
		resp.Codes = []string{"*"}
		resp.Routers, err = NewSystemMenu().GetSuperAdminRouters()
		if err != nil {
			return resp, this.Error(err)
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

// 登出
func (this SystemUser) Logout() error {

	return nil
}

func (this SystemUser) PageList(ctx *api.Context) (*dto.SystemPageListResp[model.SystemUser], *result.Error) {
	var req dto.SystemUserListReq
	validate.BindWithPanic(ctx, &req)
	if req.UserName != "" {
		this.repo.Where("username", req.UserName)
	}
	if req.NickName != "" {
		this.repo.Where("nickname", req.NickName)
	}
	if req.Phone != "" {
		this.repo.Where("phone", req.Phone)
	}
	if req.Email != "" {
		this.repo.Where("email", req.Email)
	}
	if req.Status != "" {
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
	this.repo.IsNull("deleted_at")
	if req.OrderBy != "" {
		OrderType := "DESC"
		if req.OrderType == "ascending" {
			OrderType = "ASC"
		}
		this.repo.Order(req.OrderBy + " " + OrderType)
	}
	pageList, err := this.repo.Paginate(req.PageNo, req.PageSize)
	if err != nil {
		return nil, this.Error(err)
	}
	page := ToSystemPage[model.SystemUser](pageList)
	return &page, nil
}

func (this SystemUser) ReadInfoById(ctx *api.Context, userId uint64) (dto.SystemUserInfoResp, *result.Error) {
	resp := dto.SystemUserInfoResp{}
	user, err := this.repo.GetById(cast.ToUint(userId))
	if err != nil {
		return resp, this.Error(err)
	}
	err = copier.Copy(&resp, &user)
	roleIds, _ := repo.NewSystemUserRole().Where("user_id=?", userId).Values("role_id")
	roles, err := NewSystemRole().GetRoutersByIds(roleIds)
	if err != nil {
		return resp, this.Error(err)
	}
	resp.RoleList = roles
	resp.PostList = make([]int, 0)
	return resp, nil
}

// ChangeStatus 设置用户状态
func (u SystemUser) ChangeStatus(userId uint64, status string) *result.Error {
	builder := u.repo.NewQueryBuilder().Where("id=?", userId)
	err := builder.UpdateColumn("status", status).Error
	if err != nil {
		return u.Error(err)
	}
	return nil
}

func (u SystemUser) Save(ctx *api.Context, req dto.SystemUserSaveReq) *result.Error {
	userModel := u.repo.NewModel()
	var err error
	if req.ID > 0 {
		userModel, err = u.repo.GetById(cast.ToUint(req.ID))
		if err != nil {
			return u.Error(err)
		}
		userModel.UpdatedBy = ctx.JwtClaimData.UserId
	} else {
		userModel.CreatedBy = ctx.JwtClaimData.UserId
	}
	userModel.ID = req.ID
	userModel.Username = req.UserName
	password := md5.Sum([]byte(req.Password))
	userModel.Password = fmt.Sprintf("%x", password)
	userModel.UserType = "100"
	userModel.Nickname = req.NickName
	userModel.Phone = req.Phone
	userModel.Email = req.Email
	userModel.Avatar = req.Avatar
	userModel.DeptId = req.DeptID

	userModel.Status = req.Status
	userModel.Remark = req.Remark
	userModel.LoginTime = time.Now()
	err = u.repo.Save(&userModel).Error
	if err != nil {
		return u.Error(err)
	}

	userId := userModel.GetID()

	// 保存用户-角色关系
	roleRepo := repo.NewSystemUserRole()
	for _, id := range req.RoleIds {
		roleModel := roleRepo.NewModel()
		roleModel.UserId = cast.ToUint64(userId)
		roleModel.RoleId = id
		err = roleRepo.NewQueryBuilder().Save(&roleModel).Error
		if err != nil {
			return u.Error(err)
		}
	}

	// 保存用户-岗位关系
	postRepo := repo.NewSystemUserPost()
	for _, id := range req.PostIds {
		postModel := postRepo.NewModel()
		postModel.UserId = cast.ToUint64(userId)
		postModel.PostId = id
		err = postRepo.NewQueryBuilder().Save(&postModel).Error
		if err != nil {
			return u.Error(err)
		}
	}

	return nil
}

func (u SystemUser) Delete(ids []string) *result.Error {
	err := u.repo.NewQueryBuilder().Delete(&model.SystemUser{}, ids).Error
	if err != nil {
		return u.Error(err)
	}
	return nil
}

func (u SystemUser) ResetPassword(id uint64) *result.Error {
	password := md5.Sum([]byte("123456"))
	builder := u.repo.NewQueryBuilder().Where("id=?", id)
	err := builder.UpdateColumn("password", fmt.Sprintf("%x", password)).Error

	if err != nil {
		return u.Error(err)
	}
	return nil
}

func (u SystemUser) SetHomePage(req dto.SystemUserSetHomePageReq) *result.Error {
	builder := u.repo.NewQueryBuilder().Where("id=?", req.ID)
	err := builder.UpdateColumn("dashboard", req.Dashboard).Error
	if err != nil {
		return u.Error(err)
	}
	return nil
}
