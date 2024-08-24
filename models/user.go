package models

import (
	"GBolg/conf/errmsg"
	"GBolg/dao"
	"GBolg/utils/logrus_logger"
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `grom:"type:varchar(20);not null" json:"username"`
	Password string `grom:"type:varchar(20);not null" json:"password"`
	Role     int    `grom:"type:int" json:"role"`
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
	//加密
	user.Password = ScryptPw(user.Password)
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

func UpdateUser(id int, user *User) int {
	var maps = make(map[string]interface{})
	maps["username"] = user.Username
	maps["role"] = user.Role
	err := dao.DB.Model(&User{}).Where("id=?", id).Updates(maps).Error
	if err != nil {
		logrus_logger.LogRus.Errorf("update user error: %v", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func DeleteUser(id int) int {
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
