package models

import (
	"GBolg/conf/errmsg"
	"GBolg/dao"
	"GBolg/utils/logrus_logger"
	"encoding/base64"
	"errors"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
	"time"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time      `gorm:"index" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	UserName  string         `grom:"type:varchar(20);not null" json:"username"`
	Password  string         `grom:"type:varchar(20);not null" json:"-"`
	Role      int            `grom:"type:int" json:"role"`
	Token     string         `gorm:"type:varchar(255);not null" json:"token"`
}

func (User) TableName() string {
	return "users"
}

func CheckUser(username string) (code int) {
	var user User
	dao.DB.Select("id").Where("user_name = ?", username).First(&user)
	if user.ID > 0 {
		return errmsg.ErrorUserIsExist
	}
	return errmsg.SUCCESS
}
func CheckLogin(username string, password string) (user *User, code int) {
	//err := dao.DB.Table("users").
	//	Select("id,user_name,password,role,token").
	//	Where("user_name = ?", username).
	//	First(&response).Error
	err := dao.DB.Where("user_name = ?", username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errmsg.ErrorUserNotExist
	}
	if err != nil {
		return nil, errmsg.ERROR
	}

	if ScryptPw(password) != user.Password {
		return nil, errmsg.ErrorPassword
	}
	if user.Role != 0 {
		return nil, errmsg.ErrorUserIllegalPermissions
	}
	return user, errmsg.SUCCESS

}

func CreateUser(user *User) int {
	//加密
	user.Password = ScryptPw(user.Password)
	err := dao.DB.Create(&user).Error
	if err != nil {
		logrus_logger.LogRus.Errorf("create user error: %v", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetUserList(pageSize int, pageNum int) (users []User, code int) {
	err := dao.DB.Limit(pageNum).Offset(pageSize * pageNum).Order("id desc").Find(&users).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errmsg.ErrorUserListIsEmpty
	}
	if err != nil {
		logrus_logger.LogRus.Errorf("get user list error: %v", err)
		return nil, errmsg.ERROR
	}
	//err := dao.DB.Table("users").
	//	Select("id,user_name").
	//	Limit(pageNum).Offset(pageSize * pageNum).Order("id desc").Find(&users).Error
	//if err != nil {
	//	logrus_logger.LogRus.Errorf("get user list error: %v", err)
	//	return nil, errmsg.ERROR
	//}
	return users, errmsg.SUCCESS
}

func UpdateUser(id uint, user *User) int {
	var maps = make(map[string]interface{})
	maps["user_name"] = user.UserName
	maps["role"] = user.Role
	maps["token"] = user.Token
	err := dao.DB.Model(&User{}).Where("id=?", id).Updates(maps).Error
	if err != nil {
		logrus_logger.LogRus.Errorf("update user error: %v", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func DeleteUser(id uint) int {
	//var user User
	err := dao.DB.Debug().Where("id =?", id).Delete(&User{}).Error
	if err != nil {
		logrus_logger.LogRus.Errorf("delete user error: %v", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// ScryptPw 密码加密
func ScryptPw(password string) string {
	const Keylen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 6, 66, 22, 222, 11}
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, Keylen)
	if err != nil {
		logrus_logger.LogRus.Error("scrypt pw error: %v", err)
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw

}
