// cmd/server/main.go
package main

import (
	"log"
	"my-music-go/internal/api"
	"my-music-go/internal/api/handlers"
	"my-music-go/internal/database" // 现在我们只用它来 Init
	"my-music-go/internal/repository"
	services "my-music-go/internal/services"
	// "net/http"
	// "github.com/gin-gonic/gin"
)

func main() {
	// todo log
	// logger := log.New(os.Stdout, "MUSIC_API: ", log.LstdFlags)
	// 初始化数据库连接，得到 *sqlx.DB 实例
	db := database.Init()
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("关闭数据库连接失败: %v", err)
		}
	}()

	// --- 依赖注入流程 ---
	// 使用手动依赖注入, 而不是控制反转的思路
	// 将 db 实例注入到 NewUserRepository 中，创建 repo
	userRepo := repository.NewUserRepository(db)

	// 将 repo 实例注入到 NewUserService 中，创建 service
	userService := services.NewUserService(userRepo)

	// 将 service 实例注入到 NewUserHandler 中，创建 handler
	userHandler := handlers.NewUserHandler(userService)

	router := api.SetupRouter(userHandler)

	log.Println("服务器启动于 :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
