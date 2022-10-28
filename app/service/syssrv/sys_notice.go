package syssrv

import (
	"context"
	"go-web-template/app/common/page"
	"go-web-template/app/dao/sysdao"
	"go-web-template/app/model/system"
	"go-web-template/app/model/system/request"
	"go-web-template/global"
)

type SysNoticeService struct{}

var SysNoticeSrv = new(SysNoticeService)

func (*SysNoticeService) GetSysNoticeList(ctx context.Context, sysNotice *request.SysNotice) (*page.Pagination, error) {
	sysNoticeDao := sysdao.NewSysNoticeDao(ctx)
	data, err := sysNoticeDao.SelectList(sysNotice)
	if err != nil {
		global.Logger.Error(err)
		return data, err
	}
	return data, err
}

func (*SysNoticeService) GetSysNoticeById(ctx context.Context, roleId int64) (*system.SysNotice, error) {
	sysNoticeDao := sysdao.NewSysNoticeDao(ctx)
	data, err := sysNoticeDao.SelectById(roleId)
	if err != nil {
		global.Logger.Error(err)
		return data, err
	}
	return data, err
}

func (*SysNoticeService) AddSysNotice(ctx context.Context, sysNotice *system.SysNotice) error {
	sysNoticeDao := sysdao.NewSysNoticeDao(ctx)
	err := sysNoticeDao.Insert(sysNotice)
	if err != nil {
		global.Logger.Error(err)
		return err
	}
	return nil
}

func (*SysNoticeService) UpdateNoticeById(ctx context.Context, sysNotice *system.SysNotice) error {
	sysNoticeDao := sysdao.NewSysNoticeDao(ctx)
	err := sysNoticeDao.UpdateById(sysNotice)
	if err != nil {
		global.Logger.Error(err)
		return err
	}
	return nil
}

func (*SysNoticeService) DeleteSysNoticeByIds(ctx context.Context, ids []int64) error {
	sysNoticeDao := sysdao.NewSysNoticeDao(ctx)
	err := sysNoticeDao.DeleteByIds(ids)
	if err != nil {
		global.Logger.Error(err)
		return err
	}
	return nil
}
