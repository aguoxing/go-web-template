package router

import (
	"github.com/gin-gonic/gin"
	"go-web-template/app/controller/common"
	"go-web-template/app/controller/systemctl"
	"go-web-template/configs"
	"go-web-template/middleware"
)

func InitRouter() {
	r := gin.New()
	// logrus 日志
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	captcha := common.CaptchaHandler{}
	r.GET("/captchaImage", captcha.GetCaptcha)
	r.POST("/verify", captcha.VerifyCaptcha)

	sysLoginApi := systemctl.SysLoginApi{}
	r.POST("/login", sysLoginApi.Login)
	r.POST("/logout", sysLoginApi.Logout)
	r.GET("/getInfo", sysLoginApi.GetUserInfo)
	r.GET("/getRouters", sysLoginApi.GetRouters)

	/* 系统模块 */
	systemRoutes := r.Group("system")
	// jwt 认证
	systemRoutes.Use(middleware.JWT())
	// auth 鉴权
	//systemRoutes.Use(middleware.Auth())

	// 配置管理
	configRoutes := systemRoutes.Group("config")
	configApi := systemctl.SysConfigApi{}
	{
		configRoutes.POST("/list", configApi.GetConfigList)
		configRoutes.GET("/:configId", configApi.GetConfigById)
		configRoutes.GET("/configKey/:configKey", configApi.GetConfigKey)
		configRoutes.POST("", configApi.AddConfig)
		configRoutes.PUT("", configApi.EditConfig)
		configRoutes.DELETE("", configApi.RemoveConfig)
		configRoutes.DELETE("/refreshCache", configApi.RefreshCache)
	}

	// 字典管理
	dictRoutes := systemRoutes.Group("dict")
	dictDataApi := systemctl.SysDictDataApi{}
	dictTypeApi := systemctl.SysDictTypeApi{}
	{
		dictRoutes.POST("/data/list", dictDataApi.GetDictDataList)
		dictRoutes.GET("/data/type/:dictType", dictDataApi.GetDictDataListByDictType)
		dictRoutes.GET("/data/:dictCode", dictDataApi.GetDictData)
		dictRoutes.POST("/data", dictDataApi.AddDictData)
		dictRoutes.PUT("/data", dictDataApi.UpdateDictData)
		dictRoutes.DELETE("/data", dictDataApi.DeleteDictData)

		dictRoutes.POST("/type/list", dictTypeApi.GetDictTypeList)
		dictRoutes.GET("/type/:dictId", dictTypeApi.GetDictType)
		dictRoutes.POST("/type", dictTypeApi.AddDictType)
		dictRoutes.PUT("/type", dictTypeApi.UpdateDictType)
		dictRoutes.DELETE("/type", dictTypeApi.DeleteDictType)
		dictRoutes.DELETE("/refreshCache", dictTypeApi.RefreshCache)
		dictRoutes.GET("/optionSelect", dictTypeApi.OptionSelect)
	}

	// 用户管理
	userRoutes := systemRoutes.Group("user")
	userApi := systemctl.SysUserApi{}
	{
		userRoutes.GET("/:userId", userApi.GetUserInfo)
	}

	_ = r.Run(configs.AppConfig.Server.Port)
}

func NewRouter() *gin.Engine {
	r := gin.New()
	return r
}
