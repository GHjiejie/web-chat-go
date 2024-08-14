package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"

	// 导入文件驱动
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var DB *sql.DB //这是一个全局变量，这一行代码的作用是声明一个全局变量 DB，它是一个指向 sql.DB 类型的指针。
// 这里的 sql.DB 是 Go 语言标准库中用于数据库操作的结构体。

func Connect() {
	var err error
	// 使用不带数据库名的连接字符串进行初始连接
	db, err := sql.Open("mysql", "root:12345@tcp(localhost:3306)/") // 这里的 sql.Open() 函数用于创建一个新的数据库连接，第一个参数是驱动名，第二个参数是连接字符串。
	if err != nil {
		panic(err)
	} else {
		log.Println("数据库连接成功，接下来检查数据库是否存在")
	}

	// 检查与 MySQL 服务器的连接（类比于终端ping操作，只不过这里是测试数据库连接的情况）
	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		log.Println("连接到 MySQL 服务器")
	}

	// 创建数据库（如果不存在）执行SQL语句
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS webchat")
	if err != nil {
		log.Fatalf("Error creating database: %v", err)
	}

	// 关闭初始连接
	if err = db.Close(); err != nil {
		panic(err)
	}

	// 现在使用指定的数据库名重新打开连接
	DB, err = sql.Open("mysql", "root:12345@tcp(localhost:3306)/webchat")
	if err != nil {
		panic(err)
	}

	// 验证与数据库的连接
	if err = DB.Ping(); err != nil {
		panic(err)
	} else {
		log.Println("连接到数据库 webchat")
	}

	// 进行数据库迁移()
	runMigrations()
}

func runMigrations() {
	// 创建一个msl驱动
	driver, err := mysql.WithInstance(DB, &mysql.Config{})
	if err != nil {
		log.Fatalf("创建数据库迁移实例失败: %v", err)
	}
	// 创建数据库迁移实例
	m, err := migrate.NewWithDatabaseInstance(
		"file://../../migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatalf("创建数据库迁移实例失败: %v", err)
	}

	// 执行数据库迁移
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("执行数据库迁移失败: %v", err)
	} else {
		log.Println("数据库迁移成功")
	}

	// 向数据库中插入一些测试数据
	// _, err = DB.Exec("INSERT INTO users (username, password) VALUES ('jie1', '12345')")
	// if err != nil {
	// 	log.Fatalf("插入测试数据失败: %v", err)
	// } else {
	// 	log.Println("插入测试数据成功")
	// }
}
