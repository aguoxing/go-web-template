package service

import (
	"go-web-template/app/service/system"
)

type SrvGroup struct {
	//TokenService       framework.TokenService
	//SysLoginService    framework.SysLoginService
	//SysPasswordService framework.SysPasswordService
	//SysConfigService system.SysConfigService
	//FrameSrv framework.FrameSrv
	SysSrv system.SysSrv
}

var Srv = new(SrvGroup)
