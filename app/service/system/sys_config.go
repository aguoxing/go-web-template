package system

import (
	"context"
	"go-web-template/app/dao"
	"go-web-template/app/model/system"
)

type SysConfigService struct{}

func (s *SysConfigService) SelectConfigById(ctx context.Context, configId int64) (*system.SysConfig, error) {
	configDao := dao.NewSysConfigDao(ctx)
	data, err := configDao.SelectConfigById(configId)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *SysConfigService) SelectConfigByKey(ctx context.Context, configKey string) (string, error) {
	return "", nil
}

func (s *SysConfigService) SelectCaptchaEnabled(ctx context.Context) (bool, error) {
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
	configDao := dao.NewSysConfigDao(ctx)
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
