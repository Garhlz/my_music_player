// internal/database/db.go
package database

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx" // 1. 导入 sqlx 包
	"github.com/joho/godotenv"
)

// Init 函数用于初始化数据库连接
func Init() *sqlx.DB { // 2. 返回类型从 *sql.DB 改为 *sqlx.DB
	// 加载 .env 文件
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("无法加载 .env 文件: %v", err)
	}

	// 构建数据库连接字符串 (DSN) - 这部分不变
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"), // 建议在 .env 中也加入 DB_PORT
		os.Getenv("DB_NAME"),
	)

	// 3. 使用 sqlx.Connect 替代 sql.Open 和 db.Ping()
	// sqlx.Connect 是一个便捷函数，它会帮你完成驱动连接和 Ping() 检查两步操作
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		// 如果连接失败，使用 log.Fatalf 会打印错误并终止程序
		log.Fatalf("无法连接到数据库: %v", err)
	}

	// 你可以在这里设置连接池参数
	db.SetMaxOpenConns(25) // 最大打开连接数
	db.SetMaxIdleConns(25) // 最大空闲连接数

	fmt.Println("Database connected successfully using sqlx!")
	return db
}