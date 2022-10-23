package systemctl

import (
	"github.com/gin-gonic/gin"
	"go-web-template/app/common/result"
	"go-web-template/app/model/system"
	"go-web-template/app/model/system/request"
	"go-web-template/app/service/syssrv"
	"go-web-template/global"
	"strconv"
)

type SysDictTypeApi struct{}

func (*SysDictTypeApi) GetDictTypeList(ctx *gin.Context) {
	var params request.SysDictType
	params.OpenPage = true
	_ = ctx.ShouldBindJSON(&params)
	data, err := syssrv.SysDictTypeSrv.SelectDictTypeList(ctx, params)
	if err != nil {
		result.Fail(ctx)
	} else {
		result.OkWithData(data, ctx)
	}
}

func (*SysDictTypeApi) GetDictType(ctx *gin.Context) {
	dictCode, _ := strconv.Atoi(ctx.Param("dictId"))
	data, err := syssrv.SysDictTypeSrv.SelectDictTypeById(ctx, int64(dictCode))
	if err != nil {
		result.Fail(ctx)
	} else {
		result.OkWithData(data, ctx)
	}
}

func (*SysDictTypeApi) AddDictType(ctx *gin.Context) {
	var params system.SysDictType
	_ = ctx.ShouldBindJSON(&params)
	err := syssrv.SysDictTypeSrv.AddDictType(ctx, &params)
	if err != nil {
		result.Fail(ctx)
	} else {
		result.Ok(ctx)
	}
}

func (*SysDictTypeApi) UpdateDictType(ctx *gin.Context) {
	var params system.SysDictType
	_ = ctx.ShouldBindJSON(&params)
	err := syssrv.SysDictTypeSrv.UpdateDictType(ctx, &params)
	if err != nil {
		result.Fail(ctx)
	} else {
		result.Ok(ctx)
	}
}

func (*SysDictTypeApi) DeleteDictType(ctx *gin.Context) {
	var params request.SysDictType
	_ = ctx.ShouldBindJSON(&params)
	err := syssrv.SysDictTypeSrv.DeleteDictTypeByIds(ctx, params.Ids)
	if err != nil {
		global.Logger.Error(err)
		result.Fail(ctx)
	} else {
		result.Ok(ctx)
	}
}

func (*SysDictTypeApi) RefreshCache(ctx *gin.Context) {

}

func (*SysDictTypeApi) OptionSelect(ctx *gin.Context) {

}
