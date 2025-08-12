// cmd/server/main.go
package main

import (
	"fmt"
	"log"
	_ "my-music-go/docs" // 导入 swag 生成的 docs 包
	"my-music-go/internal/api"
	"my-music-go/internal/api/handlers"
	"my-music-go/internal/config" // 引入 config 包
	"my-music-go/internal/database"
	"my-music-go/internal/repository"
	services "my-music-go/internal/services"
	"os"
)

// @title        Elaine 的音乐播放器 API
// @version      1.0
// @description  这是一个使用 Go 语言和 Gin 框架开发的音乐播放器后端 API.
// @termsOfService https://github.com/Garhlz

// @contact.name   Elaine
// @contact.url    https://github.com/Garhlz
// @contact.email  your.email@example.com

// @license.name  MIT License
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	fmt.Println("Environment Variables from OS:")
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
	logger := log.New(os.Stdout, "MUSIC_API: ", log.LstdFlags|log.Lshortfile)

	cfg, err := config.LoadConfig(".") // 从当前目录加载 .env
	if err != nil {
		logger.Fatalf("无法加载配置: %v", err)
	}

	db := database.Init(cfg)
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("关闭数据库连接失败: %v", err)
		}
	}()

	var userRepo repository.IUserRepository = repository.NewUserRepository(db)
	var songRepo repository.ISongRepository = repository.NewSongRepository(db)
	var artistRepo repository.IArtistRepository = repository.NewArtistRepository(db) // 新增
	var albumRepo repository.IAlbumRepository = repository.NewAlbumRepository(db)
	var likeRepo repository.ILikeRepository = repository.NewLikeRepository(db)
	var playlistRepo repository.IPlaylistRepository = repository.NewPlaylistRepository(db)
	var commentRepo repository.ICommentRepository = repository.NewCommentRepository(db)

	// 将 repo 和 config 实例注入到 Service 中
	userService := services.NewUserService(userRepo, cfg)
	songService := services.NewSongService(songRepo)
	artistService := services.NewArtistService(artistRepo)
	albumService := services.NewAlbumService(albumRepo)
	playlistService := services.NewPlaylistService(playlistRepo, songService)
	likeService := services.NewLikeService(likeRepo, songRepo)
	commentService := services.NewCommentService(commentRepo, songRepo)

	// 将 service 实例注入到 Handler 中
	userHandler := handlers.NewUserHandler(userService)
	songHandler := handlers.NewSongHandler(songService)
	artistHandler := handlers.NewArtistHandler(artistService, songService)
	albumHandler := handlers.NewAlbumHandler(albumService, songService)
	playlistHandler := handlers.NewPlaylistHandler(playlistService)
	likeHandler := handlers.NewLikeHandler(likeService)
	commentHandler := handlers.NewCommentHandler(commentService)

	// 将所有 handler 注入到 router 设置函数中
	router := api.SetupRouter(cfg, userHandler, songHandler, artistHandler, albumHandler, likeHandler, playlistHandler, commentHandler)

	logger.Printf("服务器启动于 %s", cfg.ServerAddress)
	if err := router.Run(cfg.ServerAddress); err != nil {
		logger.Fatalf("服务器启动失败: %v", err)
	}
}
