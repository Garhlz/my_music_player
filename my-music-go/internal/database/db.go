package database

import (
	"fmt"
	"log"
	"my-music-go/internal/config" // 引入 config 包

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Init 函数接收一个 config.Config 实例作为参数
func Init(cfg config.Config) *sqlx.DB {
	// 1. 不再需要 godotenv.Load() 和 os.Getenv()

	// 2. 从传入的 cfg 结构体中获取数据库连接信息
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	// 使用 sqlx.Connect 连接数据库
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalf("无法连接到数据库: %v", err)
	}

	// 设置连接池参数
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)

	fmt.Println("Database connected successfully using sqlx!")
	return db
}