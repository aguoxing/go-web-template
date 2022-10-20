package sysdao

import (
	"context"
	"go-web-template/app/common/e"
	"go-web-template/app/common/page"
	"go-web-template/app/model/system"
	"go-web-template/app/model/system/request"
	"go-web-template/configs"
	"gorm.io/gorm"
)

type SysUserDao struct {
	*gorm.DB
}

func NewSysUserDao(ctx context.Context) *SysUserDao {
	return &SysUserDao{configs.GetDB(ctx)}
}

// SelectById 通过id查询用户信息
func (dao *SysUserDao) SelectById(id int64) (user *system.SysUser, err error) {
	err = dao.DB.Model(&system.SysUser{}).Where("user_id=?", id).First(&user).Error
	return
}

// SelectUserByUserName 根据用户名（账号）查询用户信息
func (dao *SysUserDao) SelectUserByUserName(username string) (sysUser *system.SysUser, err error) {
	err = dao.DB.Model(&system.SysUser{}).Where("user_name=?", username).First(&sysUser).Error
	return
}

// Insert 新增用户
func (dao *SysUserDao) Insert(user *system.SysUser) error {
	return dao.DB.Model(&system.SysUser{}).Create(&user).Error
}

// DeleteById 根据 id 删除用户
func (dao *SysUserDao) DeleteById(id int64) error {
	return dao.DB.Where("user_id=?", id).Delete(&system.SysUser{}).Error
}

// UpdateById 通过 id 修改用户
func (dao *SysUserDao) UpdateById(user *system.SysUser) error {
	return dao.DB.Model(&system.SysUser{}).Where("user_id=?", user.UserID).Updates(user).Error
}

// SelectList 分页查询
func (dao *SysUserDao) SelectList(user *request.SysUser) (*page.Pagination, error) {
	var userList []*system.SysUser
	p := page.Pagination{
		PageNum:  user.PageNum,
		PageSize: user.PageSize,
	}
	err := dao.DB.Scopes(page.SelectPage(userList, &p, dao.DB)).Find(&userList).Error
	if err != nil {
		p.Code = e.ERROR
		p.Msg = err.Error()
		return nil, err
	}
	p.Rows = userList
	return &p, err
}
