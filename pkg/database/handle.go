package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // MySQL 驱动
)

// User 代表用户实体

// RegisterUser 将新用户插入数据库
func RegisterUser(db *sql.DB, user User) error {
	query := "INSERT INTO users (username, password) VALUES (?, ?)"
	_, err := db.Exec(query, user.Username, user.Password)
	return err
}
