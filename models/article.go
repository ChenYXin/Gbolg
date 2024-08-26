package models

import (
	"GBolg/conf/errmsg"
	"GBolg/dao"
	"GBolg/utils/logrus_logger"
	"errors"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Cid         int8   `gorm:"type:int" json:"cid"`
	Title       string `gorm:"type:varchar(100);not null" json:"title"`
	Description string `gorm:"type:varchar(200)" json:"description"`
	Content     string `gorm:"type:longtext" json:"content"`
	Img         string `gorm:"type:varchar(100)" json:"img"`
}

type articleResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Img         string `json:"img"`
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

func GetArticleList(pageSize int, pageNum int) (response []articleResponse, code int) {
	err := dao.DB.Table("articles").
		Select("articles.id,categorys.name as category,articles.title ,articles.content,articles.img ,articles.description").
		Joins("left join categorys on articles.cid = categorys.id").
		Limit(pageNum).Offset(pageSize * pageNum).Order("id asc").Find(&response).Error

	if len(response) == 0 {
		return nil, errmsg.ErrorArticleListNotFound
	}
	if err != nil {
		logrus_logger.LogRus.Errorf("get article list error: %v", err)
		return nil, errmsg.ERROR
	}
	return response, errmsg.SUCCESS
}

func GetArticleCategoryList(cid int, pageSize int, pageNum int) (results []articleResponse, code int) {
	err := dao.DB.Table("articles").
		Select("articles.id,categorys.name as category,articles.title ,articles.content,articles.img ,articles.description").
		Joins("left join categorys on articles.cid = categorys.id").
		Where("articles.cid = ?", cid).
		Limit(pageNum).Offset(pageSize * pageNum).Order("id asc").Find(&results).Error

	if len(results) == 0 {
		return nil, errmsg.ErrorArticleCategoryListNotFound
	}
	if err != nil {
		logrus_logger.LogRus.Errorf("get article list error: %v", err)
		return nil, errmsg.ERROR
	}
	return results, errmsg.SUCCESS
}

func GetArticleById(id uint) (response *articleResponse, code int) {
	err := dao.DB.Table("articles").
		Select("articles.id,categorys.name as category,articles.title ,articles.content,articles.img ,articles.description").
		Joins("left join categorys on articles.cid = categorys.id").
		Where("articles.id = ?", id).
		First(&response).Error
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
	err := dao.DB.Debug().Where("id =?", id).Delete(&Article{}).Error
	if err != nil {
		logrus_logger.LogRus.Errorf("delete article error: %v", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
