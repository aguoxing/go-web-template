package syssrv

import (
	"github.com/gin-gonic/gin"
	"go-web-template/app/dao/sysdao"
	"go-web-template/app/model/system"
	"go-web-template/app/model/system/response"
	"strings"
)

type SysMenuService struct{}

var SysMenuSrv = new(SysMenuService)

func (m *SysMenuService) selectMenuList(ctx *gin.Context) {

}

func (m *SysMenuService) selectTreeByPage(ctx *gin.Context) {

}

func (m *SysMenuService) selectMenuPermsByUserId(ctx *gin.Context) {

}

func (m *SysMenuService) selectMenuPermsByRoleId(ctx *gin.Context) {

}

// SelectMenuTreeByUserId 根据用户ID查询菜单
func (m *SysMenuService) SelectMenuTreeByUserId(ctx *gin.Context, sysUser *system.SysUser) ([]*system.SysMenu, error) {
	sysMenuDao := sysdao.NewSysMenuDao(ctx)
	var menus []*system.SysMenu
	if sysUser.IsAdmin(sysUser.UserID) {
		menus, _ = sysMenuDao.SelectMenuTreeAll()
	} else {

	}
	return buildTree(menus), nil
}

func (m *SysMenuService) SelectMenuListByRoleId(ctx *gin.Context) {

}

func (m *SysMenuService) GetBuildMenus(menus []*system.SysMenu) []response.RouterVo {
	return buildMenus(menus)
}

// 构建菜单树
func buildTree(menus []*system.SysMenu) []*system.SysMenu {
	menuMap := make(map[int64]*system.SysMenu)
	for i, menu := range menus {
		menu.ArrIdx = i
		menuMap[menu.MenuID] = menu
	}

	var resList []*system.SysMenu
	var childList []*system.SysMenu

	for _, menu := range menus {

		parent, ok := menuMap[menu.ParentID]

		if ok {
			if len(parent.Children) == 0 {
				childList = menus[parent.ArrIdx].Children
				if childList == nil {
					childList = []*system.SysMenu{}
				}
				childList = append(childList, menu)
				menus[parent.ArrIdx].Children = childList
			} else {
				menus[parent.ArrIdx].Children = append(menus[parent.ArrIdx].Children, menu)
			}
		}

		if menu.ParentID == 0 {
			resList = append(resList, menu)
		}
	}
	return resList
}

// BuildMenus 构建前端路由所需要的菜单
func buildMenus(menus []*system.SysMenu) []response.RouterVo {
	var routers []response.RouterVo

	for _, menu := range menus {
		meta := response.MetaVo{
			Title:   menu.MenuName,
			Icon:    menu.Icon,
			NoCache: menu.IsCache == 1,
			Link:    menu.Path,
		}
		router := response.RouterVo{
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

			var childrenList []response.RouterVo
			var child response.RouterVo
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

			var childrenList []response.RouterVo
			var child response.RouterVo
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
	routerName := strings.ToTitle(sysMenu.Path)
	// 非外链并且是一级目录（类型为目录）
	if isMenuFrame(sysMenu) {
		routerName = ""
	}
	return routerName
}

// 获取路由地址
func getRouterPath(sysMenu *system.SysMenu) string {
	routerPath := sysMenu.Path
	// 内链打开外网方式
	if sysMenu.ParentID != 0 && isInnerLink(sysMenu) {
		routerPath = innerLinkReplaceEach(routerPath)
	}
	// 非外链并且是一级目录（类型为目录）
	if sysMenu.ParentID == 0 && sysMenu.MenuType == "M" && sysMenu.IsFrame == 1 {
		routerPath = "/" + sysMenu.Path
	} else if isMenuFrame(sysMenu) {
		// 非外链并且是一级目录（类型为菜单）
		routerPath = "/"
	}
	return routerPath
}

// 获取组件信息
func getComponent(sysMenu *system.SysMenu) string {
	component := "Layout"
	if sysMenu.Component != "" && isMenuFrame(sysMenu) {
		component = sysMenu.Component
	} else if sysMenu.Component == "" && sysMenu.ParentID != 0 && isInnerLink(sysMenu) {
		component = "InnerLink"
	} else if sysMenu.Component == "" && isParentView(sysMenu) {
		component = "ParentView"
	}
	return component
}

// 是否为菜单内部跳转
func isMenuFrame(sysMenu *system.SysMenu) bool {
	return sysMenu.ParentID == 0 && sysMenu.MenuType == "C" && sysMenu.IsFrame == 1
}

// 是否为内链组件
func isInnerLink(sysMenu *system.SysMenu) bool {
	return sysMenu.IsFrame == 1
}

// 是否为parent_view组件
func isParentView(sysMenu *system.SysMenu) bool {
	return sysMenu.ParentID != 0 && sysMenu.MenuType == "M"
}

// 内链域名特殊字符替换
func innerLinkReplaceEach(str string) string {
	str1 := strings.Replace(str, "http://", "", 1)
	str2 := strings.Replace(str1, "https://", "", 1)
	str3 := strings.Replace(str2, "www.", "", 1)
	str4 := strings.Replace(str3, ".", "/", 1)
	return str4
}
