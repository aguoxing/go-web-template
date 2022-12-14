package systemctl

import (
	"github.com/gin-gonic/gin"
	"go-web-template/app/common/result"
	"go-web-template/app/model/system"
	"go-web-template/app/model/system/request"
	"go-web-template/app/service/syssrv"
	"strconv"
)

type SysUserApi struct{}

func (*SysUserApi) GetSysUserList(ctx *gin.Context) {
	var params request.SysUser
	_ = ctx.ShouldBindJSON(&params)
	data, err := syssrv.SysUserSrv.GetSysUserList(ctx, &params)
	if err != nil {
		result.FailWithMessage(err.Error(), ctx)
	} else {
		result.OkWithData(data, ctx)
	}
}

func (*SysUserApi) GetSysUser(ctx *gin.Context) {
	roleId, _ := strconv.Atoi(ctx.Param("roleId"))
	data, err := syssrv.SysUserSrv.GetSysUserById(ctx, int64(roleId))
	if err != nil {
		result.FailWithMessage(err.Error(), ctx)
	} else {
		result.OkWithData(data, ctx)
	}
}

func (*SysUserApi) AddSysUser(ctx *gin.Context) {
	var params system.SysUser
	_ = ctx.ShouldBindJSON(&params)
	err := syssrv.SysUserSrv.AddSysUser(ctx, &params)
	if err != nil {
		result.FailWithMessage(err.Error(), ctx)
	} else {
		result.Ok(ctx)
	}
}

func (*SysUserApi) UpdateSysUser(ctx *gin.Context) {
	var params system.SysUser
	_ = ctx.ShouldBindJSON(&params)
	err := syssrv.SysUserSrv.UpdateUserById(ctx, &params)
	if err != nil {
		result.FailWithMessage(err.Error(), ctx)
	} else {
		result.Ok(ctx)
	}
}

func (*SysUserApi) DeleteSysUser(ctx *gin.Context) {
	var params request.SysUser
	_ = ctx.ShouldBindJSON(&params)
	err := syssrv.SysUserSrv.DeleteSysUserByIds(ctx, params.Ids)
	if err != nil {
		result.FailWithMessage(err.Error(), ctx)
	} else {
		result.Ok(ctx)
	}
}

func (*SysUserApi) ResetPwd(ctx *gin.Context) {

}

func (*SysUserApi) ChangeStatus(ctx *gin.Context) {

}
