package systemctl

import (
	"github.com/gin-gonic/gin"
	"go-web-template/app/common/result"
	"go-web-template/app/model/system"
	"go-web-template/app/model/system/request"
	"go-web-template/app/service/syssrv"
	"strconv"
)

type SysConfigApi struct{}

// GetConfigList 获取参数列表
func (s *SysConfigApi) GetConfigList(ctx *gin.Context) {
	var params request.SysConfig
	params.OpenPage = true
	_ = ctx.ShouldBindJSON(&params)
	data, err := syssrv.SysConfigSrv.SelectSysConfigList(ctx.Request.Context(), &params)
	if err != nil {
		result.FailWithMessage(err.Error(), ctx)
	} else {
		result.OkWithData(data, ctx)
	}
}

// GetConfigById 根据参数编号获取详细信息
func (s *SysConfigApi) GetConfigById(ctx *gin.Context) {
	configId, _ := strconv.Atoi(ctx.Param("configId"))
	data, err := syssrv.SysConfigSrv.SelectSysConfigById(ctx.Request.Context(), int64(configId))
	if err != nil {
		result.FailWithMessage(err.Error(), ctx)
	} else {
		result.OkWithData(data, ctx)
	}
}

// GetConfigKey 根据参数键名查询参数值
func (s *SysConfigApi) GetConfigKey(ctx *gin.Context) {
	var config request.SysConfig
	_ = ctx.ShouldBindJSON(&config)
	data, err := syssrv.SysConfigSrv.SelectSysConfigByKey(ctx.Request.Context(), config.ConfigKey)
	if err != nil {
		result.FailWithMessage(err.Error(), ctx)
	} else {
		result.OkWithData(data, ctx)
	}
}

// AddConfig 添加配置
func (s *SysConfigApi) AddConfig(ctx *gin.Context) {
	var config system.SysConfig
	_ = ctx.ShouldBindJSON(&config)
	err := syssrv.SysConfigSrv.AddSysConfig(ctx.Request.Context(), &config)
	if err != nil {
		result.FailWithMessage(err.Error(), ctx)
	} else {
		result.Ok(ctx)
	}
}

// EditConfig 修改配置
func (s *SysConfigApi) EditConfig(ctx *gin.Context) {
	var config system.SysConfig
	_ = ctx.ShouldBindJSON(&config)
	err := syssrv.SysConfigSrv.UpdateSysConfig(ctx.Request.Context(), &config)
	if err != nil {
		result.FailWithMessage(err.Error(), ctx)
	} else {
		result.Ok(ctx)
	}
}

// RemoveConfig 删除配置
func (s *SysConfigApi) RemoveConfig(ctx *gin.Context) {
	var config request.SysConfig
	_ = ctx.ShouldBindJSON(&config)
	err := syssrv.SysConfigSrv.DeleteSysConfigByIds(ctx.Request.Context(), config.Ids)
	if err != nil {
		result.FailWithMessage(err.Error(), ctx)
	} else {
		result.Ok(ctx)
	}
}

// RefreshCache 刷新配置缓存
func (s *SysConfigApi) RefreshCache(ctx *gin.Context) {
	syssrv.SysConfigSrv.ResetConfigCache(ctx)
}
