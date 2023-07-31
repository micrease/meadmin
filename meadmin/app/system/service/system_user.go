package service

import (
	"crypto/md5"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/spf13/cast"
	"meadmin/app/system/model"
	"meadmin/app/system/repo"
	"meadmin/app/system/vo"
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

func NewSystemUser(ctx *api.Context) *SystemUser {
	service := &SystemUser{}
	service.Init(ctx)
	service.repo = repo.NewSystemUser()
	return service
}

func (this *SystemUser) Login(req vo.SystemLoginReq) (vo.SystemLoginResp, *result.Error) {
	model, err := this.repo.Where("username=?", req.Username).First()
	resp := vo.SystemLoginResp{}
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

func (this *SystemUser) GetInfo() (vo.SystemInfoResp, *result.Error) {
	resp := vo.SystemInfoResp{}

	userId := this.ctx.JwtClaimData.UserId
	IsSuperAdmin := this.ctx.JwtClaimData.IsSuperAdmin

	user, err := this.repo.GetById(cast.ToUint(userId))
	if err != nil {
		return resp, this.Error(err)
	}

	resp.User = vo.User{}
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
func (this *SystemUser) Logout() error {
	return nil
}

func (this *SystemUser) PageList(ctx *api.Context) (*vo.SystemPageListResp[model.SystemUser], *result.Error) {
	var req vo.SystemUserListReq
	validate.BindWithPanic(ctx, &req)

	query := this.repo.NewQuery()
	if len(req.UserName) > 0 {
		query.Where("username", req.UserName)
	}

	if len(req.NickName) > 0 {
		query.Where("nickname", req.NickName)
	}

	if len(req.Phone) > 0 {
		query.Where("phone", req.Phone)
	}

	if len(req.Email) > 0 {
		query.Where("email", req.Email)
	}

	if len(req.Status) > 0 {
		query.Where("status", req.Status)
	}

	if len(req.MinDate) > 0 {
		location, err := time.Parse("2006-01-02", req.MinDate)
		if err == nil {
			query.Where("created_at", ">", location)
		}
	}

	if len(req.MaxDate) > 0 {
		location, err := time.Parse("2006-01-02", req.MaxDate)
		if err == nil {
			query.Where("created_at", "<", location)
		}
	}

	query.IsNull("deleted_at")
	if len(req.OrderBy) > 0 {
		OrderType := "DESC"
		if req.OrderType == "ascending" {
			OrderType = "ASC"
		}
		query.Order(req.OrderBy + " " + OrderType)
	}

	pageList, err := query.Paginate(req.PageNo, req.PageSize)
	if err != nil {
		return nil, this.Error(err)
	}
	page := ToSystemPage[model.SystemUser](pageList)
	return &page, nil
}

func (this *SystemUser) ReadInfoById(ctx *api.Context, userId uint64) (vo.SystemUserInfoResp, *result.Error) {
	resp := vo.SystemUserInfoResp{}
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
func (this *SystemUser) ChangeStatus(userId uint64, status string) *result.Error {
	builder := this.repo.NewQueryBuilder().Where("id=?", userId)
	err := builder.UpdateColumn("status", status).Error
	if err != nil {
		return this.Error(err)
	}
	return nil
}

func (this *SystemUser) Save(ctx *api.Context, req vo.SystemUserSaveReq) *result.Error {

	userModel := model.SystemUser{}
	var err error
	if req.ID > 0 {
		userModel, err := this.repo.GetById(cast.ToUint(req.ID))
		if err != nil {
			return this.Error(err)
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
	err = this.repo.Save(&userModel).Error
	if err != nil {
		return this.Error(err)
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
			return this.Error(err)
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
			return this.Error(err)
		}
	}

	return nil
}

func (this *SystemUser) Delete(ids []string) *result.Error {
	err := this.repo.NewQueryBuilder().Delete(&model.SystemUser{}, ids).Error
	if err != nil {
		return this.Error(err)
	}
	return nil
}

func (this *SystemUser) ResetPassword(id uint64) *result.Error {
	password := md5.Sum([]byte("123456"))
	builder := this.repo.NewQueryBuilder().Where("id=?", id)
	err := builder.UpdateColumn("password", fmt.Sprintf("%x", password)).Error
	if err != nil {
		return this.Error(err)
	}
	return nil
}

func (this *SystemUser) SetHomePage(req vo.SystemUserSetHomePageReq) *result.Error {
	builder := this.repo.NewQueryBuilder().Where("id=?", req.ID)
	err := builder.UpdateColumn("dashboard", req.Dashboard).Error
	if err != nil {
		return this.Error(err)
	}
	return nil
}
