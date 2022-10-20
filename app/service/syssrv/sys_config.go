package syssrv

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-web-template/app/dao/sysdao"
	"go-web-template/app/model/system"
)

type SysConfigService struct{}

var SysConfigSrv = new(SysConfigService)

func (s *SysConfigService) SelectConfigById(ctx context.Context, configId int64) (*system.SysConfig, error) {
	configDao := sysdao.NewSysConfigDao(ctx)
	data, err := configDao.SelectConfigById(configId)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *SysConfigService) SelectConfigByKey(ctx context.Context, configKey string) (string, error) {
	return "", nil
}

func (s *SysConfigService) SelectCaptchaEnabled(ctx *gin.Context) (bool, error) {
	return false, nil
}

func (s *SysConfigService) SelectConfigList(ctx context.Context) (bool, error) {
	return true, nil
}

func (s *SysConfigService) InsertConfig(ctx context.Context, config *system.SysConfig) (int, error) {
	return 200, nil
}

func (s *SysConfigService) UpdateConfig(ctx context.Context, config *system.SysConfig) (int, error) {
	return 200, nil
}

func (s *SysConfigService) DeleteConfigByIds(ctx context.Context, ids []int64) error {
	configDao := sysdao.NewSysConfigDao(ctx)
	return configDao.DeleteConfigByIds(ids)
}

func checkConfigKeyUnique(ctx *context.Context, config *system.SysConfig) (string, error) {
	return "", nil
}

func loadingConfigCache(ctx *context.Context, ids []*system.SysConfig) error {
	return nil
}

func clearConfigCache(ctx *context.Context, ids []*system.SysConfig) error {
	return nil
}
