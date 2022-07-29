package service

import (
	"giftcard/application/dto"
	models "giftcard/application/model"
	"giftcard/application/repo"
	"giftcard/library/context/api"
	"github.com/jinzhu/copier"
	"github.com/spf13/cast"
)

type SystemMenu struct {
	repo *repo.SystemMenu
}

func NewSystemMenu() SystemMenu {
	service := SystemMenu{}
	service.repo = repo.NewSystemMenu()
	return service
}

func (this SystemMenu) GetMenuList(ctx *api.Context) ([]dto.MenuTreeDTO, error) {
	var menuList []models.SystemMenu
	var err error
	if ctx.JwtClaimData.IsSuperAdmin {
		menuList, err = this.repo.Order("sort desc").Where("status=?", models.StatusEnable).List()
		if err != nil {
			return nil, err
		}
	} else {
		userId := ctx.JwtClaimData.UserId
		userRoleRepo := repo.NewSystemUserRole()
		builder := userRoleRepo.NewQueryBuilder().Where("user_id=?", userId)
		roleIds := repo.NewSystemUserRole().Values(builder, "role_id")
		menuIds := this.GetMenuIdsByRoleIds(roleIds)
		menuList, err = this.repo.Order("sort desc").
			Where("status=?", models.StatusEnable).
			Where("id in(?)", menuIds).
			List()
		if err != nil {
			return nil, err
		}
	}

	var menuTreeList []dto.MenuTreeDTO
	for _, menu := range menuList {
		var menuTree dto.MenuTreeDTO
		err = copier.Copy(&menuTree, &menu)
		if err != nil {
			return nil, err
		}
		menuTreeList = append(menuTreeList, menuTree)
	}
	return this.ToMenuTree(menuTreeList, 0), nil
}

func (this SystemMenu) ToMenuTree(data []dto.MenuTreeDTO, parentId uint64) []dto.MenuTreeDTO {
	var tree []dto.MenuTreeDTO
	if len(data) == 0 {
		return tree
	}
	for _, value := range data {
		if value.ParentId == parentId {
			child := this.ToMenuTree(data, value.ID)
			if len(child) > 0 {
				value.Children = child
			} else {
				value.Children = []dto.MenuTreeDTO{}
			}
			tree = append(tree, value)
		}
	}
	return tree
}

func (this SystemMenu) GetSuperAdminRouters() ([]dto.RouterTreeDTO, error) {
	list, err := this.repo.Order("sort desc").
		Where("status=?", models.StatusEnable).
		List()
	if err != nil {
		return nil, err
	}
	return this.SysMenuToRouterTree(list), nil
}

func (this SystemMenu) GetRoutersByIds(menuIds []any) ([]dto.RouterTreeDTO, error) {
	list, err := this.repo.Order("sort desc").
		Where("status=?", models.StatusEnable).
		Where("id in(?)", menuIds).
		List()
	if err != nil {
		return nil, err
	}
	return this.SysMenuToRouterTree(list), nil
}

//getMenuCode
func (this SystemMenu) GetMenuCode(menuIds []any) []any {
	if len(menuIds) == 0 {
		return []any{}
	}
	builder := this.repo.NewQueryBuilder().Where("id in(?)", menuIds)
	return this.repo.Values(builder, "code")
}

func (this SystemMenu) SysMenuToRouterTree(menuList []models.SystemMenu) []dto.RouterTreeDTO {
	var routerTree []dto.RouterTreeDTO
	if len(menuList) == 0 {
		return routerTree
	}
	routers := []dto.RouterTreeDTO{}
	for _, menu := range menuList {
		router := this.SetRouter(menu)
		routers = append(routers, router)
	}
	return this.ToTree(routers, 0)
}

func (this SystemMenu) GetMenuIdsByRoleIds(roleIds []any) []any {
	if len(roleIds) == 0 {
		return []any{}
	}
	roleRepo := repo.NewSystemRoleMenu()
	builder := roleRepo.NewQueryBuilder().Where("role_id in(?)", roleIds)
	menuIds := roleRepo.Values(builder, "menu_id")

	builder = this.repo.NewQueryBuilder().Where("id in(?)", menuIds)
	return this.repo.Values(builder, "id")
}

func (this SystemMenu) ToTree(data []dto.RouterTreeDTO, parentId uint) []dto.RouterTreeDTO {
	var tree []dto.RouterTreeDTO
	if len(data) == 0 {
		return tree
	}

	for _, value := range data {
		if value.ParentId == parentId {
			child := this.ToTree(data, value.ID)
			if len(child) > 0 {
				value.Children = child
			} else {
				value.Children = []dto.RouterTreeDTO{}
			}
			tree = append(tree, value)
		}
	}
	return tree
}

func (this SystemMenu) SetRouter(menu models.SystemMenu) dto.RouterTreeDTO {
	route := "/" + menu.Route
	if menu.Type == "L" {
		route = menu.Route
	}
	routerDTO := dto.RouterTreeDTO{
		ID:        cast.ToUint(menu.ID),
		ParentId:  cast.ToUint(menu.ParentId),
		Name:      menu.Code,
		Component: menu.Component,
		Path:      route,
		Redirect:  menu.Redirect,
		Meta: dto.RouterMetaDTO{
			Type:             menu.Type,
			Icon:             menu.Icon,
			Title:            menu.Name,
			Hidden:           menu.IsHidden == 0,
			HiddenBreadcrumb: false,
		},
	}
	return routerDTO
}
