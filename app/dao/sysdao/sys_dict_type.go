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

type SysDictTypeDao struct {
	*gorm.DB
}

func NewSysDictTypeDao(ctx context.Context) *SysDictTypeDao {
	return &SysDictTypeDao{configs.GetDB(ctx)}
}

// SelectList 根据条件分页查询字典类型
func (dao *SysDictTypeDao) SelectList(dictType request.SysDictType) (p *page.Pagination, err error) {
	var DictTypeList []*system.SysDictType
	p = new(page.Pagination)

	if dictType.DictName != "" {
		dao.DB = dao.DB.Where("dict_name = ?", dictType.DictName)
	}
	if dictType.Status != "" {
		dao.DB = dao.DB.Where("status = ?", dictType.Status)
	}
	if dictType.DictType != "" {
		dao.DB = dao.DB.Where("dict_type = ?", dictType.DictType)
	}

	if dictType.OpenPage {
		p.PageNum = dictType.PageNum
		p.PageSize = dictType.PageSize
		err = dao.DB.Scopes(page.SelectPage(DictTypeList, p, dao.DB)).Find(&DictTypeList).Error
	} else {
		err = dao.DB.Find(&DictTypeList).Error
	}
	p.Rows = DictTypeList
	if err != nil {
		p.Code = e.ERROR
		p.Msg = err.Error()
		return p, err
	}
	return p, err
}

func (dao *SysDictTypeDao) SelectAll() (list []system.SysDictType, err error) {
	err = dao.DB.Find(&list).Error
	return
}

func (dao *SysDictTypeDao) SelectById(dictId int64) (dictType *system.SysDictType, err error) {
	err = dao.DB.Where("dict_id = ?", dictId).Find(&dictType).Error
	if err != nil {
		return nil, err
	}
	return
}

func (dao *SysDictTypeDao) Insert(dictType *system.SysDictType) error {
	return dao.DB.Create(dictType).Error
}

func (dao *SysDictTypeDao) UpdateById(dictType *system.SysDictType) error {
	return dao.DB.Save(dictType).Error
}

func (dao *SysDictTypeDao) DeleteByIds(ids []int64) error {
	return dao.DB.Where("dict_id in (?)", ids).Delete(&system.SysDictType{}).Error
}

func (dao *SysDictTypeDao) CheckDictTypeUnique(dictType string) (count int64, err error) {
	err = dao.DB.Model(&system.SysDictType{}).Where("dict_type=?", dictType).Count(&count).Error
	return
}
