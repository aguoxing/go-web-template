package systemctl

import (
	"github.com/gin-gonic/gin"
	"go-web-template/app/common/result"
	"go-web-template/app/model/system"
	"go-web-template/app/model/system/request"
	"go-web-template/app/service/syssrv"
	"strconv"
)

type SysDictDataApi struct{}

func (*SysDictDataApi) GetDictDataList(ctx *gin.Context) {
	var params request.SysDictData
	params.OpenPage = true
	_ = ctx.ShouldBindJSON(&params)
	data, err := syssrv.SysDictDataSrv.SelectDictDataList(ctx, params)
	if err != nil {
		result.FailWithMessage(err.Error(), ctx)
	} else {
		result.OkWithData(data, ctx)
	}
}

func (*SysDictDataApi) GetDictDataListByDictType(ctx *gin.Context) {
	var params request.SysDictData
	params.OpenPage = false
	params.DictType = ctx.Param("dictType")
	data, err := syssrv.SysDictDataSrv.SelectDictDataList(ctx, params)
	if err != nil {
		result.FailWithMessage(err.Error(), ctx)
	} else {
		result.OkWithData(data, ctx)
	}
}

func (*SysDictDataApi) GetDictData(ctx *gin.Context) {
	dictCode, _ := strconv.Atoi(ctx.Param("dictCode"))
	data, err := syssrv.SysDictDataSrv.SelectDictDataById(ctx, int64(dictCode))
	if err != nil {
		result.FailWithMessage(err.Error(), ctx)
	} else {
		result.OkWithData(data, ctx)
	}
}

func (*SysDictDataApi) AddDictData(ctx *gin.Context) {
	var params system.SysDictData
	_ = ctx.ShouldBindJSON(&params)
	err := syssrv.SysDictDataSrv.AddDictData(ctx, &params)
	if err != nil {
		result.FailWithMessage(err.Error(), ctx)
	} else {
		result.Ok(ctx)
	}
}

func (*SysDictDataApi) UpdateDictData(ctx *gin.Context) {
	var params system.SysDictData
	_ = ctx.ShouldBindJSON(&params)
	err := syssrv.SysDictDataSrv.UpdateDictData(ctx, &params)
	if err != nil {
		result.FailWithMessage(err.Error(), ctx)
	} else {
		result.Ok(ctx)
	}
}

func (*SysDictDataApi) DeleteDictData(ctx *gin.Context) {
	var params request.SysDictData
	_ = ctx.ShouldBindJSON(&params)
	err := syssrv.SysDictDataSrv.DeleteDictDataByIds(ctx, params.Ids)
	if err != nil {
		result.FailWithMessage(err.Error(), ctx)
	} else {
		result.Ok(ctx)
	}
}
