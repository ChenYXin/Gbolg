package models

import (
	"GBolg/conf/errmsg"
	"GBolg/dao"
	"GBolg/utils/logrus_logger"
	"errors"
	"gorm.io/gorm"
	"time"
)

type Article struct {
	ID          uint           `gorm:"primarykey"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `gorm:"index" json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Cid         int8           `gorm:"type:int" json:"cid"`
	Title       string         `gorm:"type:varchar(100);not null" json:"title"`
	Description string         `gorm:"type:varchar(200)" json:"description"`
	Content     string         `gorm:"type:longtext" json:"content"`
	Img         string         `gorm:"type:varchar(100)" json:"img"`
}

func (Article) TableName() string {
	return "articles"
}

func CreateArticle(article *Article) int {
	err := dao.DB.Create(&article).Error
	if err != nil {
		logrus_logger.LogRus.Errorf("create article error: %v", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetArticleList(pageSize int, pageNum int) (articles []*Article, total int64, code int) {
	//err := dao.DB.Table("articles").
	//	Select("articles.id,categorys.name as category,articles.title ,articles.content,articles.img ,articles.description").
	//	Joins("left join categorys on articles.cid = categorys.id").
	//	Limit(pageNum).Offset(pageSize * pageNum).Order("id asc").Find(&response).Error
	//
	//dao.DB.Count(&total)
	//if len(response) == 0 {
	//	return nil, errmsg.ErrorArticleListNotFound
	//}
	//if err != nil {
	//	logrus_logger.LogRus.Errorf("get article list error: %v", err)
	//	return nil, errmsg.ERROR
	//}
	//return response, errmsg.SUCCESS

	db := dao.DB.Model(&Article{}) //.Where("id >= ?", 0)
	db.Count(&total)

	err := db.Order("updated_at desc").Offset(pageSize * pageNum).Limit(pageNum).Find(&articles).Error
	if total == 0 {
		return nil, 0, errmsg.ErrorArticleListNotFound
	}
	if err != nil {
		logrus_logger.LogRus.Errorf("get article list error: %v", err)
		return nil, 0, errmsg.ERROR
	}

	return articles, total, errmsg.SUCCESS
}

func GetArticleCategoryList(cid int, pageSize int, pageNum int) (articles []Article, total int64, code int) {
	//err := dao.DB.Table("articles").
	//	Select("articles.id,categorys.name as category,articles.title ,articles.content,articles.img ,articles.description").
	//	Joins("left join categorys on articles.cid = categorys.id").
	//	Where("articles.cid = ?", cid).
	//	Limit(pageNum).Offset(pageSize * pageNum).Order("id asc").Find(&results).Error
	db := dao.DB.Model(&Article{}).Where("cid = ?", cid)
	db.Count(&total)

	err := db.Order("id asc").Offset(pageSize * pageNum).Limit(pageNum).Find(&articles).Error

	//if len(results) == 0 {
	//	return nil, errmsg.ErrorArticleCategoryListNotFound
	//}
	//if err != nil {
	//	logrus_logger.LogRus.Errorf("get article list error: %v", err)
	//	return nil, errmsg.ERROR
	//}
	if len(articles) == 0 {
		return nil, 0, errmsg.ErrorArticleCategoryListNotFound
	}
	if err != nil {
		logrus_logger.LogRus.Errorf("get article category list error: %v", err)
		return nil, 0, errmsg.ERROR
	}

	return articles, total, errmsg.SUCCESS
}

func GetArticleById(id uint) (article *Article, code int) {
	//err := dao.DB.Table("articles").
	//	Select("articles.id,categorys.name as category,articles.title ,articles.content,articles.img ,articles.description").
	//	Joins("left join categorys on articles.cid = categorys.id").
	//	Where("articles.id = ?", id).
	//	First(&response).Error
	err := dao.DB.Where("id = ? ", id).First(&article).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errmsg.ErrorArticleInfoNotFound
	}
	if err != nil {
		logrus_logger.LogRus.Errorf("get article info error: %v", err)
		return nil, errmsg.ERROR
	}

	code = errmsg.SUCCESS
	return
}

func UpdateArticle(id uint, article *Article) int {
	var maps = make(map[string]interface{})
	maps["title"] = article.Title
	maps["cid"] = article.Cid
	maps["description"] = article.Description
	maps["content"] = article.Content
	maps["img"] = article.Img
	err := dao.DB.Model(&Article{}).Where("id=?", id).Updates(maps).Error
	if err != nil {
		logrus_logger.LogRus.Errorf("update article error: %v", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func DeleteArticle(id uint) int {
	err := dao.DB.Where("id =?", id).Delete(&Article{}).Error
	if err != nil {
		logrus_logger.LogRus.Errorf("delete article error: %v", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
