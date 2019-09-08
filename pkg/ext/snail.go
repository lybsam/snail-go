package ext

import (
	"fmt"
	"snail/pkg/model"
	"snail/pkg/setting"
)

func FindsExt(v interface{}, page int, sql string, args ...interface{}) bool {
	e := model.So().Where(fmt.Sprintf("%s AND deleted_on = ?", sql), args, 0).
		Order("-created_on").
		Limit(setting.Conf.PageSize).
		Offset(setting.Conf.PageSize * page).
		Find(v).Error
	return IsErrorExt(e)
}

func FindExt(v interface{}, page int) bool {
	e := model.So().Where("deleted_on = ?", 0).
		Order("-created_on").
		Limit(setting.Conf.PageSize).
		Offset(setting.Conf.PageSize * page).
		Find(v).Error
	return IsErrorExt(e)
}

func FirstExt(v interface{}, sql string, args ...interface{}) bool {
	e := model.So().Where(fmt.Sprintf("%s AND deleted_on = ?", sql), args, 0).First(v).Error
	return IsErrorExt(e)
}

func CreateExt(v interface{}) bool {
	e := model.So().Create(v).Error
	return IsErrorExt(e)
}

func DelExt(v interface{}, sql string, args ...interface{}) bool {
	e := model.So().Where(fmt.Sprintf("%s AND deleted_on = ?", sql), args, 0).Delete(v).Error
	return IsErrorExt(e)
}

func IsErrorExt(err error) bool {
	if err != nil {
		return false
	}
	return true
}
