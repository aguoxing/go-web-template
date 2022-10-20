package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("鉴权中间件")

		// TODO...

		// 前置中间件
		c.Next()
	}
}
