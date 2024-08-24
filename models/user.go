package models

import (
	"GBolg/conf/errmsg"
	"GBolg/dao"
	"GBolg/utils/logrus_logger"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `grom:"type:varchar(20);not null" json:"username"`
	Password string `grom:"type:varchar(20);not null" json:"password"`
	//Role     int    `grom:"type:int" json:"role"`
}

func (User) TableName() string {
	return "users"
}

func CheckUser(username string) (code int) {
	var user User
	dao.DB.Select("id").Where("username = ?", username).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //1001
	}
	return errmsg.SUCCESS
}

func CreateUser(user *User) int {
	err := dao.DB.Create(&user).Error
	if err != nil {
		logrus_logger.LogRus.Errorf("create user error: %v", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetUserList(pageSize int, pageNum int) (users []User) {
	err := dao.DB.Limit(pageNum).Offset(pageSize * pageNum).Order("id desc").Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logrus_logger.LogRus.Errorf("get user list error: %v", err)
		return nil
	}
	return users
}

//func GetUserById(id int) (user User) {
//
//}
