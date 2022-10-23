package syssrv

import (
	"context"
	"go-web-template/app/common/page"
	"go-web-template/app/dao/sysdao"
	"go-web-template/app/model/system"
	"go-web-template/app/model/system/request"
	"go-web-template/global"
)

type SysDictTypeService struct{}

var SysDictTypeSrv = new(SysDictTypeService)

func (*SysDictTypeService) Init() {

}

func (*SysDictTypeService) SelectDictTypeList(ctx context.Context, dictType request.SysDictType) (*page.Pagination, error) {
	sysDictTypeDao := sysdao.NewSysDictTypeDao(ctx)
	data, err := sysDictTypeDao.SelectList(dictType)
	if err != nil {
		global.Logger.Error(err)
		return data, err
	}
	return data, err
}

func (*SysDictTypeService) SelectDictTypeAll() {

}

func (*SysDictTypeService) SelectDictDataByType() {

}

func (*SysDictTypeService) SelectDictTypeById(ctx context.Context, dictId int64) (*system.SysDictType, error) {
	sysDictTypeDao := sysdao.NewSysDictTypeDao(ctx)
	data, err := sysDictTypeDao.SelectById(dictId)
	if err != nil {
		global.Logger.Error(err)
		return nil, err
	}
	return data, err
}

func (*SysDictTypeService) SelectDictTypeByType(ctx context.Context, dictType request.SysDictType) (*page.Pagination, error) {
	sysDictTypeDao := sysdao.NewSysDictTypeDao(ctx)
	data, err := sysDictTypeDao.SelectList(dictType)
	if err != nil {
		global.Logger.Error(err)
		return nil, err
	}
	return data, err
}

func (*SysDictTypeService) DeleteDictTypeByIds(ctx context.Context, ids []int64) error {
	sysDictTypeDao := sysdao.NewSysDictTypeDao(ctx)
	err := sysDictTypeDao.DeleteByIds(ids)
	if err != nil {
		global.Logger.Error(err)
		return err
	}
	return nil
}

func (s *SysDictTypeService) AddDictType(ctx context.Context, dictType *system.SysDictType) error {
	sysDictTypeDao := sysdao.NewSysDictTypeDao(ctx)
	err := sysDictTypeDao.Insert(dictType)
	if err != nil {
		global.Logger.Error(err)
		return err
	}
	return nil
}

func (*SysDictTypeService) UpdateDictType(ctx context.Context, dictType *system.SysDictType) error {
	sysDictTypeDao := sysdao.NewSysDictTypeDao(ctx)
	err := sysDictTypeDao.UpdateById(dictType)
	if err != nil {
		global.Logger.Error(err)
		return err
	}
	return nil
}

func (*SysDictTypeService) CheckDictTypeUnique() {

}

func loadingDictCache() {

}

func clearDictCache() {

}

func resetDictCache() {

}
