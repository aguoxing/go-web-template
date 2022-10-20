package system

import (
	"context"
	"go-web-template/app/model/system"
	"go-web-template/configs"
	"gorm.io/gorm"
)

type SysMenuDao struct {
	*gorm.DB
}

func NewSysMenuDao(ctx context.Context) *SysMenuDao {
	return &SysMenuDao{configs.GetDB(ctx)}
}

// SelectMenuTreeAll 查询所有目录M、菜单C
func (dao *SysMenuDao) SelectMenuTreeAll() (menus []*system.SysMenu, err error) {
	err = dao.DB.Model(&system.SysMenu{}).Where("menu_type in (M,C)?").Where("status = ?", 0).Order("parent_id,order_num").Error
	return
}

// SelectMenuTreeByUserId 根据用户Id查询所有目录M、菜单C 涉及sys_role_menu、sys_user_role、sys_role、sys_user
func (dao *SysMenuDao) SelectMenuTreeByUserId() (menus []*system.SysMenu, err error) {
	err = dao.DB.Model(&system.SysMenu{}).Where("user_id=?", "1").Error
	return
}
