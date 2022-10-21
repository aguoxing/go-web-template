package response

import (
	"encoding/json"
	"go-web-template/app/model/system"
	"time"
)

// SysUserResp 已登录用户的聚合信息
type SysUserResp struct {
	SysUser *system.SysUser   `json:"sysUser"` // 用户信息
	SysDept *system.SysDept   `json:"sysDept"` // 用户所在部门信息
	Roles   []*system.SysRole `json:"roles"`   // 用户角色信息集合
	RoleIds []int64           `json:"roleIds"` // 角色id集合
	RoleId  int64             `json:"roleId"`  // 角色id
	PostIds []int64           `json:"postIds"` // 岗位id集合
}

// LoginUser 已登录用户的聚合信息 缓存到redis
type LoginUser struct {
	UserID        int64        `json:"userId"`
	DeptID        int64        `json:"deptId"`
	UserKey       string       `json:"userKey"`
	LoginTime     time.Time    `json:"loginTime"`
	ExpireTime    time.Time    `json:"expireTime"`
	IpAddr        string       `json:"ipAddr"`
	LoginLocation string       `json:"loginLocation"`
	Browser       string       `json:"browser"`
	Os            string       `json:"os"`
	Permissions   []string     `json:"permissions"`
	SysUserResp   *SysUserResp `json:"sysUserResp"`
}

// UserInfo 返回前端的
type UserInfo struct {
	User        *system.SysUser   `json:"user"`
	Roles       []*system.SysRole `json:"roles"`
	Permissions []string          `json:"permissions"`
}

func (m *SysUserResp) MarshalBinary() (data []byte, err error) {
	return json.Marshal(m)
}

func (m *SysUserResp) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)

}
