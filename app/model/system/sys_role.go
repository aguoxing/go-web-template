// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package system

import (
	"time"
)

const TableNameSysRole = "sys_role"

// SysRole mapped from table <sys_role>
type SysRole struct {
	RoleID            int64     `gorm:"column:role_id;type:bigint;primaryKey;autoIncrement:true" json:"roleId"`        // 角色ID
	RoleName          string    `gorm:"column:role_name;type:varchar(30);not null" json:"roleName"`                    // 角色名称
	RoleKey           string    `gorm:"column:role_key;type:varchar(100);not null" json:"roleKey"`                     // 角色权限字符串
	RoleSort          int64     `gorm:"column:role_sort;type:int;not null" json:"roleSort"`                            // 显示顺序
	DataScope         string    `gorm:"column:data_scope;type:char(1);default:1" json:"dataScope"`                     // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	MenuCheckStrictly bool      `gorm:"column:menu_check_strictly;type:tinyint(1);default:1" json:"menuCheckStrictly"` // 菜单树选择项是否关联显示
	DeptCheckStrictly bool      `gorm:"column:dept_check_strictly;type:tinyint(1);default:1" json:"deptCheckStrictly"` // 部门树选择项是否关联显示
	Status            string    `gorm:"column:status;type:char(1);not null" json:"status"`                             // 角色状态（0正常 1停用）
	DelFlag           string    `gorm:"column:del_flag;type:char(1);default:0" json:"delFlag"`                         // 删除标志（0代表存在 2代表删除）
	CreateBy          string    `gorm:"column:create_by;type:varchar(64)" json:"createBy"`                             // 创建者
	CreateTime        time.Time `gorm:"column:create_time;type:int unsigned;autoCreateTime" json:"createTime"`         // 创建时间
	UpdateBy          string    `gorm:"column:update_by;type:varchar(64)" json:"updateBy"`                             // 更新者
	UpdateTime        time.Time `gorm:"column:update_time;type:int unsigned;autoUpdateTime" json:"updateTime"`         // 更新时间
	Remark            string    `gorm:"column:remark;type:varchar(500)" json:"remark"`                                 // 备注
}

// TableName SysRole's table name
func (*SysRole) TableName() string {
	return TableNameSysRole
}
