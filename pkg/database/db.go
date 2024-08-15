package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB // 使用 GORM 代替 sql.DB,这个DB就可以去操作数据库了

func Connect() {
	var err error

	// 使用 GORM 连接到数据库
	dsn := "root:12345@tcp(localhost:3306)/webchat?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	log.Println("连接到数据库 webchat")

	// 自动迁移你的模型
	err = DB.AutoMigrate(&User{}) // 假设您有一个 User 结构体
	if err != nil {
		log.Fatalf("自动迁移失败: %v", err)
	} else {
		log.Println("数据库迁移成功")
	}

}
