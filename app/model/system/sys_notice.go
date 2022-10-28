// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package system

import (
	"time"
)

const TableNameSysNotice = "sys_notice"

// SysNotice mapped from table <sys_notice>
type SysNotice struct {
	NoticeID      int64     `gorm:"column:notice_id;type:int;primaryKey;autoIncrement:true" json:"noticeId"` // 公告ID
	NoticeTitle   string    `gorm:"column:notice_title;type:varchar(50);not null" json:"noticeTitle"`        // 公告标题
	NoticeType    string    `gorm:"column:notice_type;type:char(1);not null" json:"noticeType"`              // 公告类型（1通知 2公告）
	NoticeContent []byte    `gorm:"column:notice_content;type:longblob" json:"noticeContent"`                // 公告内容
	Status        string    `gorm:"column:status;type:char(1);default:0" json:"status"`                      // 公告状态（0正常 1关闭）
	CreateBy      string    `gorm:"column:create_by;type:varchar(64)" json:"createBy"`                       // 创建者
	CreateTime    time.Time `gorm:"column:create_time;type:int unsigned;autoCreateTime" json:"createTime"`   // 创建时间
	UpdateBy      string    `gorm:"column:update_by;type:varchar(64)" json:"updateBy"`                       // 更新者
	UpdateTime    time.Time `gorm:"column:update_time;type:int unsigned;autoUpdateTime" json:"updateTime"`   // 更新时间
	Remark        string    `gorm:"column:remark;type:varchar(255)" json:"remark"`                           // 备注
}

// TableName SysNotice's table name
func (*SysNotice) TableName() string {
	return TableNameSysNotice
}