package system

import (
	"github.com/gin-gonic/gin"
	"go-web-template/app/dao"
	"go-web-template/app/framework"
	"go-web-template/app/model/system"
	"go-web-template/app/model/system/response"
)

type SysMenuService struct{}

func (m *SysMenuService) selectMenuList(ctx *gin.Context) {

}

func (m *SysMenuService) selectTreeByPage(ctx *gin.Context) {

}

func (m *SysMenuService) selectMenuPermsByUserId(ctx *gin.Context) {

}

func (m *SysMenuService) selectMenuPermsByRoleId(ctx *gin.Context) {

}

// SelectMenuTreeByUserId 根据用户ID查询菜单
func (m *SysMenuService) SelectMenuTreeByUserId(ctx *gin.Context, userId int64) (menus []*system.SysMenu, err error) {
	sysMenuDao := dao.NewSysMenuDao(ctx)
	t := framework.TokenService{}
	loginUser, err := t.GetLoginUser(ctx)
	if err != nil {

	}
	if loginUser.SysUserResp.SysUser.IsAdmin(loginUser.UserID) {
		menus, err = sysMenuDao.SelectMenuTreeAll()
	} else {

	}
	return buildTree(menus), nil
}

func (m *SysMenuService) SelectMenuListByRoleId(ctx *gin.Context) {

}

func (m *SysMenuService) GetBuildMenus(menus []*system.SysMenu) []*response.RouterVo {
	return buildMenus(menus)
}

// 构建菜单树
func buildTree(menus []*system.SysMenu) []*system.SysMenu {
	var mapTemp map[int64]*system.SysMenu
	for _, menu := range menus {
		mapTemp[menu.MenuID] = menu
	}

	var resList []*system.SysMenu
	var childList []*system.SysMenu

	for _, menu := range menus {

		parent, ok := mapTemp[menu.ParentID]

		if ok && parent != nil {
			if parent.Children != nil {
				childList = parent.Children
				if childList != nil {
					childList = []*system.SysMenu{}
				}
				childList = append(childList, menu)
				parent.Children = childList
			} else {
				parent.Children = append(parent.Children, menu)
			}
		}

		if parent == nil || menu.ParentID == 0 {
			resList = append(resList, menu)
		}
	}
	return resList
}

// BuildMenus 构建前端路由所需要的菜单
func buildMenus(menus []*system.SysMenu) []*response.RouterVo {
	var routers []*response.RouterVo

	for _, menu := range menus {
		meta := response.MetaVo{
			Title:   menu.MenuName,
			Icon:    menu.Icon,
			NoCache: menu.IsCache == 1,
			Link:    menu.Path,
		}
		router := &response.RouterVo{
			Hidden:    menu.Visible == "1",
			Name:      getRouteName(menu),
			Path:      getRouterPath(menu),
			Component: getComponent(menu),
			Query:     menu.Query,
			Meta:      &meta,
		}

		childMenus := menu.Children
		if len(childMenus) > 0 && menu.MenuType == "M" {
			router.AlwaysShow = true
			router.Redirect = "noRedirect"
			router.Children = buildMenus(childMenus)
		} else if isMenuFrame(menu) {
			var m *response.MetaVo = nil
			router.Meta = m

			var childrenList []*response.RouterVo
			var child *response.RouterVo
			routerPath := innerLinkReplaceEach(menu.Path)
			child.Path = innerLinkReplaceEach(routerPath)
			child.Component = "InnerLink"
			child.Name = routerPath
			child.Query = menu.Query
			child.Meta = &response.MetaVo{
				Title: menu.MenuName,
				Icon:  menu.Icon,
				Link:  menu.Path,
			}
			childrenList = append(childrenList, child)
			router.Children = childrenList
		} else if menu.ParentID == 0 && isInnerLink(menu) {
			router.Meta = &response.MetaVo{
				Title: menu.MenuName,
				Icon:  menu.Icon,
			}
			router.Path = menu.Path

			var childrenList []*response.RouterVo
			var child *response.RouterVo
			child.Path = menu.Path
			child.Component = menu.Component
			child.Name = menu.Path
			child.Query = menu.Query
			child.Meta = &response.MetaVo{
				Title:   menu.MenuName,
				Icon:    menu.Icon,
				Link:    menu.Path,
				NoCache: menu.IsCache == 1,
			}
			childrenList = append(childrenList, child)
			router.Children = childrenList
		}
		routers = append(routers, router)
	}

	return routers
}

// 获取路由名称
func getRouteName(sysMenu *system.SysMenu) string {
	return ""
}

// 获取路由地址
func getRouterPath(sysMenu *system.SysMenu) string {
	return ""
}

// 获取组件信息
func getComponent(sysMenu *system.SysMenu) string {
	return ""
}

// 是否为菜单内部跳转
func isMenuFrame(sysMenu *system.SysMenu) bool {
	return false
}

// 是否为内链组件
func isInnerLink(sysMenu *system.SysMenu) bool {
	return false
}

// 是否为parent_view组件
func isParentView(sysMenu *system.SysMenu) bool {
	return false
}

// 内链域名特殊字符替换
func innerLinkReplaceEach(str string) string {
	return ""
}
