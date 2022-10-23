package middleware

import (
	"github.com/gin-gonic/gin"
	"go-web-template/app/common/e"
	"go-web-template/configs"
	"go-web-template/util"
	"strings"
	"time"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code int
		var data interface{}
		code = 200
		token := getToken(ctx)
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
			ctx.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   data,
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func getToken(ctx *gin.Context) string {
	token := ctx.GetHeader(configs.AppConfig.JWT.Header)
	t := strings.Replace(token, "Bearer ", "", 1)
	return t
}
