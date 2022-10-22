package middleware

import (
	"github.com/gin-gonic/gin"
	"go-web-template/app/common/result"
	"go-web-template/app/framework"
	"go-web-template/global"
)

// HasPerm 验证用户是否具备某权限
func HasPerm(perms string) gin.HandlerFunc {
	return func(c *gin.Context) {

		loginUser, _ := framework.TokenSrv.GetLoginUser(c)

		if loginUser == nil || len(loginUser.Permissions) < 0 {
			global.Logger.Warn("没有权限")
			c.Abort()
			result.Forbidden(c)
			return
		}
		if hasPermissions(loginUser.Permissions, perms) {
			c.Next()
		} else {
			global.Logger.Warn("没有权限")
			c.Abort()
			result.Forbidden(c)
			return
		}
	}
}

// HasRole 判断用户是否拥有某个角色
func HasRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if role == "" {
			return
		}

		loginUser, _ := framework.TokenSrv.GetLoginUser(c)

		if loginUser == nil || len(loginUser.SysUserResp.Roles) < 0 {
			global.Logger.Warn("没有权限")
			c.Abort()
			result.Forbidden(c)
			return
		}

		for _, sysRole := range loginUser.SysUserResp.Roles {
			roleKey := sysRole.RoleKey
			if roleKey == "admin" || role == roleKey {
				c.Next()
				return
			} else {
				global.Logger.Warn("没有权限")
				c.Abort()
				result.Forbidden(c)
				return
			}
		}
	}
}

func hasPermissions(permissions []string, perm string) bool {
	if perm == "*:*:*" {
		return true
	}
	for _, permission := range permissions {
		if perm == permission {
			return true
		}
	}
	return false
}
