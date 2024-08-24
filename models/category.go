package models

import (
	"GBolg/conf/errmsg"
	"GBolg/dao"
	"GBolg/utils/logrus_logger"
	"gorm.io/gorm"
)

// Category 分类
type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

func (Category) TableName() string {
	return "categorys"
}

func CheckCategory(name string) (code int) {
	var category Category
	dao.DB.Select("id").Where("name = ?", name).First(&category)
	if category.ID > 0 {
		return errmsg.ERROR_CATEGORY_USED //2001
	}
	return errmsg.SUCCESS
}

func CreateCategory(category *Category) int {
	err := dao.DB.Create(&category).Error
	if err != nil {
		logrus_logger.LogRus.Errorf("create category error: %v", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetCategoryList(pageSize int, pageNum int) (categorys []Category) {
	err := dao.DB.Limit(pageNum).Offset(pageSize * pageNum).Order("id asc").Find(&categorys).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logrus_logger.LogRus.Errorf("get category list error: %v", err)
		return nil
	}
	return categorys
}

func UpdateCategory(id int, category *Category) int {
	var maps = make(map[string]interface{})
	maps["name"] = category.Name
	err := dao.DB.Model(&Category{}).Where("id=?", id).Updates(maps).Error
	if err != nil {
		logrus_logger.LogRus.Errorf("update category error: %v", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func DeleteCategory(id int) int {
	err := dao.DB.Debug().Where("id =?", id).Delete(&Category{}).Error
	if err != nil {
		logrus_logger.LogRus.Errorf("delete category error: %v", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
