package system

import (
	"context"
	"go-web-template/app/common/e"
	"go-web-template/app/common/page"
	"go-web-template/app/model/system"
	"go-web-template/app/model/system/request"
	"go-web-template/configs"
	"gorm.io/gorm"
)

type SysConfigDao struct {
	*gorm.DB
}

func NewSysConfigDao(ctx context.Context) *SysConfigDao {
	return &SysConfigDao{configs.GetDB(ctx)}
}

func (dao *SysConfigDao) SelectConfig(config *system.SysConfig) (sysConfig *system.SysConfig, err error) {
	err = dao.DB.Model(&system.SysConfig{}).Where(&config).First(&system.SysConfig{}).Error
	return
}

func (dao *SysConfigDao) SelectConfigById(id int64) (config *system.SysConfig, err error) {
	err = dao.DB.Model(&system.SysConfig{}).Where("config_id=?", id).First(&config).Error
	return
}

func (dao *SysConfigDao) SelectConfigList(config *request.SysConfig) (*page.Pagination, error) {
	var configList []*system.SysConfig
	p := page.Pagination{
		PageNum:  config.PageNum,
		PageSize: config.PageSize,
	}
	err := dao.DB.Scopes(page.SelectPage(configList, &p, dao.DB)).Find(&configList).Error
	if err != nil {
		p.Code = e.ERROR
		p.Msg = err.Error()
		return nil, err
	}
	p.Rows = configList
	return &p, err
}

func (dao *SysConfigDao) InsertConfig(config *system.SysConfig) error {
	return dao.DB.Model(&system.SysConfig{}).Create(&config).Error
}

func (dao *SysConfigDao) DeleteConfigById(id int64) error {
	return dao.DB.Where("config_id=?", id).Delete(&system.SysConfig{}).Error
}

func (dao *SysConfigDao) DeleteConfigByIds(ids []int64) error {
	//return dao.DB.Where("config_id IN ?", ids).Delete(&system.SysConfig{}).Error
	return dao.DB.Delete(&system.SysConfig{}, "config_id IN ?", ids).Error
}

func (dao *SysConfigDao) UpdateConfigById(config *system.SysConfig) error {
	return dao.DB.Model(&system.SysConfig{}).Where("config_id=?", config.ConfigID).Updates(&config).Error
}

func (dao *SysConfigDao) CheckConfigKeyUnique(config *system.SysConfig) (count int64, err error) {
	err = dao.DB.Model(&system.SysConfig{}).Where("config_key=?", config.ConfigKey).Count(&count).Error
	return
}
