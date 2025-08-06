package api

import (
	"my-music-go/internal/api/handlers"
	"my-music-go/internal/api/middleware"
	"my-music-go/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter 接收所有需要的 handler
func SetupRouter(
	cfg config.Config,
	userHandler *handlers.UserHandler,
	songHandler *handlers.SongHandler,
	artistHandler *handlers.ArtistHandler,
	albumHandler *handlers.AlbumHandler,
) *gin.Engine {

	router := gin.Default()

	// --- 基础路由, 位于所有分组之外 ---
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authMiddleware := middleware.NewAuthMiddleware(cfg)

	// API v1 路由组
	apiV1 := router.Group("/api/v1")
	{
		// --- 公共路由组 (Public Routes) ---
		// 只有登录和注册是真正公开的
		authRoutes := apiV1.Group("/auth")
		{
			authRoutes.POST("/register", userHandler.Register)
			authRoutes.POST("/login", userHandler.Login)
		}

		// --- 认证保护路由组 (Protected Routes) ---
		// 这个分组下的所有接口都需要通过 JWT 中间件认证
		protected := apiV1.Group("/")
		protected.Use(authMiddleware) // 使用闭包传递回的函数
		{
			userRoutes := protected.Group("/users")
			{
				userRoutes.GET("/:id", userHandler.GetUserProfile)
				userRoutes.GET("/:id/name", userHandler.GetUsernameAndName)
				userRoutes.PUT("/:id", userHandler.UpdateUserProfile)
			}

			songRoutes := protected.Group("/songs")
			{
				songRoutes.GET("", songHandler.ListSongs)
				songRoutes.GET("/:id", songHandler.GetSongDetail)
			}

			artistRoutes := protected.Group("/artists")
			{
				artistRoutes.GET("/:id", artistHandler.GetArtistDetail)
			}

			albumRoutes := protected.Group("/albums")
			{
				albumRoutes.GET("", albumHandler.ListAlbums)
				albumRoutes.GET("/:id", albumHandler.GetAlbumDetail)
			}
		}
	}

	return router
}
