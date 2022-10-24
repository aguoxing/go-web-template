package middleware

import (
	"github.com/gin-gonic/gin"
	"go-web-template/app/common/e"
	"go-web-template/app/common/result"
	"go-web-template/app/framework"
	"go-web-template/configs"
	"go-web-template/util"
	"strings"
	"time"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code = 200
		token := getToken(ctx)
		if token == "" {
			code = 404
		} else {
			claims, err := util.ParseToken(token)
			loginUser, _ := framework.TokenSrv.GetLoginUser(ctx)
			framework.TokenSrv.VerifyToken(loginUser)
			if err != nil {
				code = e.UNAUTHORIZED
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.FORBIDDEN
			}
		}
		if code != e.SUCCESS {
			result.FailWithDetailed(result.Response{Code: code, Msg: e.GetMsg(code)}, e.GetMsg(code), ctx)
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
