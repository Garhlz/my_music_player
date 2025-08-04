// internal/api/router.go
package api

import (
	"my-music-go/internal/api/handlers" // 假设未来会有中间件
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/swaggo/gin-swagger"
)

// SetupRouter 负责初始化和注册所有路由
func SetupRouter(
	userHandler *handlers.UserHandler,
	// 未来会有 songHandler, artistHandler 等等，都在这里加
) *gin.Engine {

	// 1. 创建 Gin 引擎
	router := gin.Default()

	// 2. 在这里注册所有路由
	// 健康检查接口
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Swagger 文档路由
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 使用路由分组来更好地组织 API
	// 比如，所有 V1 版本的 API 都在 /api/v1 路径下
	apiV1 := router.Group("/mmp-go/api/dev")
	{
		// 用户认证相关的路由组
		authRoutes := apiV1.Group("/auth")
		{
			authRoutes.POST("/register", userHandler.Register)
			authRoutes.POST("/login", userHandler.Login) // 登录接口
		}

		// 用户信息相关的路由组 (假设需要 JWT 认证)
		// userRoutes := apiV1.Group("/users")
		// userRoutes.Use(middleware.AuthMiddleware()) // 对整个组应用认证中间件
		// {
		// 	// userRoutes.GET("/:id", userHandler.GetProfile)
		// }

		// // 歌曲相关的路由组
		// songRoutes := apiV1.Group("/songs")
		// {
		// 	// songRoutes.GET("/", songHandler.ListSongs)
		// 	// songRoutes.GET("/:id", songHandler.GetSong)
		// }
	}

	// 3. 返回配置好的 router 实例
	return router
}
