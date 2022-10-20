package systemctl

import (
	"github.com/gin-gonic/gin"
	"go-web-template/app/common/result"
	"go-web-template/app/model/system"
	"go-web-template/app/model/system/request"
	"go-web-template/app/service/syssrv"
	"go-web-template/global"
	"strconv"
)

type SysConfigApi struct{}

// GetConfigList 获取参数列表
func (s *SysConfigApi) GetConfigList(ctx *gin.Context) {
	//data, err := service.Srv.SysConfigService.SelectConfigList(ctx)
	data, err := syssrv.SysConfigSrv.SelectConfigList(ctx)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	result.OkWithData(data, ctx)
}

// GetConfigById 根据参数编号获取详细信息
func (s *SysConfigApi) GetConfigById(ctx *gin.Context) {
	configId, _ := strconv.Atoi(ctx.Param("configId"))
	//data, err := service.Srv.SysConfigService.SelectConfigById(ctx, int64(configId))
	data, err := syssrv.SysConfigSrv.SelectConfigById(ctx, int64(configId))
	if err != nil {
		global.Logger.Error(err)
		return
	}
	result.OkWithData(data, ctx)
}

// GetConfigKey 根据参数键名查询参数值
func (s *SysConfigApi) GetConfigKey(ctx *gin.Context) {
	var config request.SysConfig
	_ = ctx.ShouldBindJSON(&config)

	//data, err := service.Srv.SysConfigService.SelectConfigByKey(ctx, config.ConfigKey)
	data, err := syssrv.SysConfigSrv.SelectConfigByKey(ctx, config.ConfigKey)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	result.OkWithData(data, ctx)
}

// AddConfig 添加配置
func (s *SysConfigApi) AddConfig(ctx *gin.Context) {
	var config system.SysConfig
	_ = ctx.ShouldBindJSON(&config)

	//_, err := service.Srv.SysConfigService.InsertConfig(ctx, &config)
	_, err := syssrv.SysConfigSrv.InsertConfig(ctx, &config)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	result.Ok(ctx)
}

// EditConfig 修改配置
func (s *SysConfigApi) EditConfig(ctx *gin.Context) {
	var config system.SysConfig
	_ = ctx.ShouldBindJSON(&config)

	//_, err := service.Srv.SysConfigService.UpdateConfig(ctx, &config)
	_, err := syssrv.SysConfigSrv.UpdateConfig(ctx, &config)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	result.Ok(ctx)
}

// RemoveConfig 删除配置
func (s *SysConfigApi) RemoveConfig(ctx *gin.Context) {
	var config request.SysConfig
	_ = ctx.ShouldBindJSON(&config)

	//err := service.Srv.SysConfigService.DeleteConfigByIds(ctx, config.Ids)
	err := syssrv.SysConfigSrv.DeleteConfigByIds(ctx, config.Ids)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	result.Ok(ctx)
}

// RefreshCache 刷新配置缓存
func (s *SysConfigApi) RefreshCache(ctx *gin.Context) {

}
