package request

import (
	"time"
)

type SysDictType struct {
	OpenPage   bool      `json:"openPage"`                 // 开启分页
	PageNum    int       `json:"PageNum" form:"PageNum"`   // 页码
	PageSize   int       `json:"pageSize" form:"pageSize"` // 每页大小
	Ids        []int64   `json:"ids"`                      // dictIds
	DictID     int64     `json:"dictId"`                   // 字典主键
	DictName   string    `json:"dictName"`                 // 字典名称
	DictType   string    `json:"dictType"`                 // 字典类型
	Status     string    `json:"status"`                   // 状态（0正常 1停用）
	CreateBy   string    `json:"createBy"`                 // 创建者
	CreateTime time.Time `json:"createTime"`               // 创建时间
	UpdateBy   string    `json:"updateBy"`                 // 更新者
	UpdateTime time.Time `json:"updateTime"`               // 更新时间
	Remark     string    `json:"remark"`                   // 备注
}
