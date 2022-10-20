package middleware

import (
	"github.com/gin-gonic/gin"
	"go-web-template/app/common/e"
	"go-web-template/configs"
	"go-web-template/util"
	"time"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = 200
		token := c.GetHeader(configs.AppConfig.JWT.Header)
		if token == "" {
			code = 404
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.UNAUTHORIZED
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.FORBIDDEN
			}
		}
		if code != e.SUCCESS {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

// JWTAdmin token验证中间件
func JWTAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		token := c.GetHeader("Authorization")
		if token == "" {
			code = e.BAD_REQUEST
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.UNAUTHORIZED
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.FORBIDDEN
			}
		}
		if code != e.SUCCESS {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
