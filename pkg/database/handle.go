package database

import (
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL 驱动
	"gorm.io/gorm"
)

// User 代表用户实体

// RegisterUser 将新用户插入数据库
func RegisterUser(db *gorm.DB, user User) error {
	log.Println("输出插入的用户信息", user)
	return db.Create(&user).Error
}

// 查询用户是否存在，返回值为 User 和 error，因为可能查询不到用户，所以返回值有两个
func GetUserByUsername(db *gorm.DB, username string) (User, error) {
	var user User
	err := db.Where("username = ?", username).First(&user).Error
	return user, err
}
