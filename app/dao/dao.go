package dao

import (
	"context"
	"go-web-template/app/dao/system"
	"go-web-template/configs"
)

func NewUserDao(ctx context.Context) *system.SysUserDao {
	return &system.SysUserDao{DB: configs.GetDB(ctx)}
}

func NewSysMenuDao(ctx context.Context) *system.SysMenuDao {
	return &system.SysMenuDao{DB: configs.GetDB(ctx)}
}

func NewSysConfigDao(ctx context.Context) *system.SysConfigDao {
	return &system.SysConfigDao{DB: configs.GetDB(ctx)}
}
