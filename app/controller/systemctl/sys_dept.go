package systemctl

import (
	"github.com/gin-gonic/gin"
	"go-web-template/app/common/result"
	"go-web-template/app/model/system"
	"go-web-template/app/model/system/request"
	"go-web-template/app/service/syssrv"
	"strconv"
)

type SysDeptApi struct{}

func (*SysDeptApi) GetSysDeptList(ctx *gin.Context) {
	var params request.SysDept
	params.OpenPage = true
	_ = ctx.ShouldBindJSON(&params)
	data, err := syssrv.SysDeptSrv.GetSysDeptList(ctx, &params)
	if err != nil {
		result.FailWithMessage(err.Error(), ctx)
	} else {
		result.OkWithData(data, ctx)
	}
}

func (*SysDeptApi) GetSysDeptTreeList(ctx *gin.Context) {
	var params request.SysDept
	//params.OpenPage = true
	_ = ctx.ShouldBindJSON(&params)
	data, err := syssrv.SysDeptSrv.GetDeptTreeList(ctx, &params)
	if err != nil {
		result.FailWithMessage(err.Error(), ctx)
	} else {
		result.OkWithData(data, ctx)
	}
}

func (*SysDeptApi) GetSysDept(ctx *gin.Context) {
	deptId, _ := strconv.Atoi(ctx.Param("deptId"))
	data, err := syssrv.SysDeptSrv.GetSysDeptById(ctx, int64(deptId))
	if err != nil {
		result.FailWithMessage(err.Error(), ctx)
	} else {
		result.OkWithData(data, ctx)
	}
}

func (*SysDeptApi) AddSysDept(ctx *gin.Context) {
	var params system.SysDept
	_ = ctx.ShouldBindJSON(&params)
	err := syssrv.SysDeptSrv.AddSysDept(ctx, &params)
	if err != nil {
		result.FailWithMessage(err.Error(), ctx)
	} else {
		result.Ok(ctx)
	}
}

func (*SysDeptApi) UpdateSysDept(ctx *gin.Context) {
	var params system.SysDept
	_ = ctx.ShouldBindJSON(&params)
	err := syssrv.SysDeptSrv.UpdateDeptById(ctx, &params)
	if err != nil {
		result.FailWithMessage(err.Error(), ctx)
	} else {
		result.Ok(ctx)
	}
}

func (*SysDeptApi) DeleteSysDept(ctx *gin.Context) {
	var params request.SysDept
	_ = ctx.ShouldBindJSON(&params)
	err := syssrv.SysDeptSrv.DeleteSysDeptByIds(ctx, params.Ids)
	if err != nil {
		result.FailWithMessage(err.Error(), ctx)
	} else {
		result.Ok(ctx)
	}
}
