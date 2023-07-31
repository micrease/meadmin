package service

import (
	"github.com/jinzhu/copier"
	"github.com/spf13/cast"
	"meadmin/app/system/model"
	"meadmin/app/system/repo"
	"meadmin/app/system/vo"
	"meadmin/library/context/api"
)

type SystemMenu struct {
	repo *repo.SystemMenu
}

func NewSystemMenu() *SystemMenu {
	service := &SystemMenu{}
	service.repo = repo.NewSystemMenu()
	return service
}

func (this *SystemMenu) GetMenuList(ctx *api.Context) ([]vo.MenuTree, error) {
	var menuList []model.SystemMenu
	var err error
	if ctx.JwtClaimData.IsSuperAdmin {
		menuList, err = this.repo.Order("sort desc").Where("status=?", model.StatusEnable).List()
		if err != nil {
			return nil, err
		}
	} else {
		userId := ctx.JwtClaimData.UserId
		roleIds, err := repo.NewSystemUserRole().Where("user_id", userId).Values("role_id")
		menuIds := this.GetMenuIdsByRoleIds(roleIds)
		menuList, err = this.repo.Order("sort desc").
			Where("status=?", model.StatusEnable).
			Where("id in(?)", menuIds).
			List()
		if err != nil {
			return nil, err
		}
	}

	var menuTreeList []vo.MenuTree
	for _, menu := range menuList {
		var menuTree vo.MenuTree
		err = copier.Copy(&menuTree, &menu)
		if err != nil {
			return nil, err
		}
		menuTreeList = append(menuTreeList, menuTree)
	}
	return this.ToMenuTree(menuTreeList, 0), nil
}

func (this *SystemMenu) ToMenuTree(data []vo.MenuTree, parentId uint64) []vo.MenuTree {
	var tree []vo.MenuTree
	if len(data) == 0 {
		return tree
	}
	for _, value := range data {
		if value.ParentId == parentId {
			child := this.ToMenuTree(data, value.ID)
			if len(child) > 0 {
				value.Children = child
			} else {
				value.Children = []vo.MenuTree{}
			}
			tree = append(tree, value)
		}
	}
	return tree
}

func (this *SystemMenu) GetSuperAdminRouters() ([]vo.RouterTree, error) {
	list, err := this.repo.Order("sort desc").
		Where("status", model.StatusEnable).
		List()
	if err != nil {
		return nil, err
	}
	return this.SysMenuToRouterTree(list), nil
}

func (this *SystemMenu) GetRoutersByIds(menuIds []any) ([]vo.RouterTree, error) {
	list, err := this.repo.Order("sort desc").
		Where("status", model.StatusEnable).
		WhereIn("id", menuIds).
		List()
	if err != nil {
		return nil, err
	}
	return this.SysMenuToRouterTree(list), nil
}

// getMenuCode
func (this *SystemMenu) GetMenuCode(menuIds []any) []any {
	if len(menuIds) == 0 {
		return []any{}
	}
	ids, _ := this.repo.WhereIn("id", menuIds).Values("code")
	return ids
}

func (this *SystemMenu) SysMenuToRouterTree(menuList []model.SystemMenu) []vo.RouterTree {
	var routerTree []vo.RouterTree
	if len(menuList) == 0 {
		return routerTree
	}
	routers := []vo.RouterTree{}
	for _, menu := range menuList {
		router := this.SetRouter(menu)
		routers = append(routers, router)
	}
	return this.ToTree(routers, 0)
}

func (this *SystemMenu) GetMenuIdsByRoleIds(roleIds []any) []any {
	if len(roleIds) == 0 {
		return []any{}
	}
	roleRepo := repo.NewSystemRoleMenu()
	menuIds, _ := roleRepo.WhereIn("role_id", roleIds).Values("menu_id")
	ids, _ := this.repo.WhereIn("id", menuIds).Values("id")
	return ids
}

func (this *SystemMenu) ToTree(data []vo.RouterTree, parentId uint) []vo.RouterTree {
	var tree []vo.RouterTree
	if len(data) == 0 {
		return tree
	}

	for _, value := range data {
		if value.ParentId == parentId {
			child := this.ToTree(data, value.ID)
			if len(child) > 0 {
				value.Children = child
			} else {
				value.Children = []vo.RouterTree{}
			}
			tree = append(tree, value)
		}
	}
	return tree
}

func (this *SystemMenu) SetRouter(menu model.SystemMenu) vo.RouterTree {
	route := "/" + menu.Route
	if menu.Type == "L" {
		route = menu.Route
	}
	router := vo.RouterTree{
		ID:        cast.ToUint(menu.ID),
		ParentId:  cast.ToUint(menu.ParentId),
		Name:      menu.Code,
		Component: menu.Component,
		Path:      route,
		Redirect:  menu.Redirect,
		Meta: vo.RouterMeta{
			Type:             menu.Type,
			Icon:             menu.Icon,
			Title:            menu.Name,
			Hidden:           menu.IsHidden == 0,
			HiddenBreadcrumb: false,
		},
	}
	return router
}
