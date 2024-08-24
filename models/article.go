package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}

func (Article) TableName() string {
	return "articles"
}
