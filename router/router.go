package router

import (
	"github.com/gin-gonic/gin"
	"go-web-template/app/controller/common"
	"go-web-template/app/controller/system"
	"go-web-template/configs"
	"go-web-template/middleware"
)

func InitRouter() {
	r := gin.New()
	// logrus 日志
	r.Use(middleware.Logger())

	//r.GET("/", index.Index)
	//r.GET("/login", index.Login)
	//r.POST("/checklogin", index.CheckLogin)
	//r.GET("/captchaImage", index.CaptchaImage)
	//r.GET("/500", errorc.Error)
	//r.GET("/404", errorc.NotFound)
	//r.GET("/403", "")
	captcha := common.CaptchaHandler{}
	r.GET("/captcha", captcha.GetCaptcha)
	r.POST("/verify", captcha.VerifyCaptcha)

	sysLoginApi := system.SysLoginApi{}
	r.POST("/login", sysLoginApi.Login)

	/* 系统模块 */
	systemRoutes := r.Group("system")
	// jwt 认证
	//systemRoutes.Use(middleware.JWT())
	// auth 鉴权
	systemRoutes.Use(middleware.Auth())

	// 配置管理
	configRoutes := systemRoutes.Group("config")
	configApi := system.SysConfigApi{}
	{
		configRoutes.GET("/:configId", configApi.GetConfigById)
		configRoutes.POST("/list", configApi.GetConfigList)
		configRoutes.POST("/add", configApi.AddConfig)
		configRoutes.PUT("/update", configApi.EditConfig)
		configRoutes.DELETE("/delete", configApi.RemoveConfig)
	}

	// 用户管理
	userRoutes := systemRoutes.Group("user")
	userApi := system.SysUserApi{}
	{
		userRoutes.GET("/:userId", userApi.GetUserInfo)
	}

	_ = r.Run(configs.AppConfig.Server.Port)
}

func NewRouter() *gin.Engine {
	r := gin.New()
	return r
}
