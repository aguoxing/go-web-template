package systemctl

import (
	"github.com/gin-gonic/gin"
	"go-web-template/app/common/result"
	"go-web-template/app/framework"
	"go-web-template/app/model/system/request"
	"go-web-template/app/model/system/response"
	"go-web-template/app/service/syssrv"
	"go-web-template/global"
)

type SysLoginApi struct{}

// Login 用户登录
func (s *SysLoginApi) Login(ctx *gin.Context) {
	var loginBody request.LoginBody
	_ = ctx.ShouldBindJSON(&loginBody)

	//token, err := service.Srv.SysLoginService.Login(ctx, &loginBody)
	//t := framework.SysLoginService{}
	token, err := framework.SysLoginSrv.Login(ctx, &loginBody)
	if err != nil {
		global.Logger.Error(err)
		result.FailWithMessage("登录失败", ctx)
	} else {
		result.OkWithMessage(token, ctx)
	}
}

// GetUserInfo 获取当前登录用户信息
func (s *SysLoginApi) GetUserInfo(ctx *gin.Context) {
	//t := framework.TokenService{}
	loginUser, err := framework.TokenSrv.GetLoginUser(ctx)
	if err != nil {

	}
	//p := framework.SysPermissions{}
	// todo 角色 set
	//p.GetRolePermission(loginUser.SysUserResp)
	// todo 权限 set
	//p.GetMenuPermission(loginUser.SysUserResp)
	userInfo := &response.UserInfo{
		User:        loginUser.SysUserResp.SysUser,
		Roles:       loginUser.SysUserResp.Roles,
		Permissions: loginUser.Permissions,
	}
	result.OkWithData(userInfo, ctx)
}

// GetRouters 获取前端路由信息 菜单
func (s *SysLoginApi) GetRouters(ctx *gin.Context) {
	// todo 菜单树
	//t := framework.TokenService{}
	loginUser, err := framework.TokenSrv.GetLoginUser(ctx)
	if err != nil {
		global.Logger.Error(err)
	}
	//m := syssrv.SysMenuService{}
	menus, err := syssrv.SysMenuSrv.SelectMenuTreeByUserId(ctx, loginUser.SysUserResp.SysUser)
	if err != nil {
		result.Fail(ctx)
	} else {
		result.OkWithData(syssrv.SysMenuSrv.GetBuildMenus(menus), ctx)
	}
}
