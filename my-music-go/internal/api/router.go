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
	likeHandler *handlers.LikeHandler,
	playlistHandler *handlers.PlaylistHandler,
	commentHandler *handlers.CommentHandler,
) *gin.Engine {

	router := gin.Default()

	// --- 基础路由 ---
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authMiddleware := middleware.NewAuthMiddleware(cfg)

	// API v1 路由组
	apiV1 := router.Group("/api/v1")
	{
		// --- 公共路由组 (Public Routes) ---
		// 只有登录和注册是真正公开的，放在这里
		authRoutes := apiV1.Group("/auth")
		{
			authRoutes.POST("/register", userHandler.Register)
			authRoutes.POST("/login", userHandler.Login)
		}

		// --- 认证保护路由组 (Protected Routes) ---
		// 所有其他业务接口都放在这个由中间件保护的分组下
		protected := apiV1.Group("/")
		protected.Use(authMiddleware) // 使用闭包传递回的函数
		{
			{
				// “我”的专属路由，处理对当前登录用户自身资源的操作
				meRoutes := protected.Group("/me")
				{
					meRoutes.PUT("/profile", userHandler.UpdateMyProfile)

					likedSongsRoutes := meRoutes.Group("/liked-songs")
					{
						likedSongsRoutes.GET("", likeHandler.ListMyLikedSongs)
						likedSongsRoutes.POST("/:songId", likeHandler.LikeSong)
						likedSongsRoutes.DELETE("/:songId", likeHandler.UnlikeSong)
					}

					playlistRoutes := meRoutes.Group("/playlists")
					{
						playlistRoutes.GET("", playlistHandler.ListMyPlaylists)
						playlistRoutes.POST("", playlistHandler.CreatePlaylist)
						playlistRoutes.PUT("/:playlistId", playlistHandler.UpdatePlaylist)
						playlistRoutes.DELETE("/:playlistId", playlistHandler.DeletePlaylist)
						playlistRoutes.POST("/:playlistId/songs", playlistHandler.AddSongToPlaylist)
						playlistRoutes.DELETE("/:playlistId/songs/:songId", playlistHandler.RemoveSongFromPlaylist)
					}
				}

				// 对其他用户或通用资源的查询与操作
				userRoutes := protected.Group("/users")
				{
					userRoutes.GET("/:id", userHandler.GetUserProfile)
					userRoutes.GET("/:id/name", userHandler.GetUsernameAndName)
					userRoutes.GET("/:id/liked-songs", likeHandler.ListUserLikedSongs)
					userRoutes.GET("/:id/playlists", playlistHandler.ListUserPlaylists)
				}

				songRoutes := protected.Group("/songs")
				{
					songRoutes.GET("", songHandler.ListSongs)
					songRoutes.GET("/:id", songHandler.GetSongDetail)
					songRoutes.GET("/:id/comments", commentHandler.ListSongComments)
					songRoutes.POST("/:id/comments", commentHandler.CreateComment)
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

				playlistsRoutes := protected.Group("/playlists")
				{
					playlistsRoutes.GET("/playlists/:playlistId", playlistHandler.GetPlaylistDetail)
				}

				commentsRoutes := protected.Group("/comments")
				{
					commentsRoutes.GET("/:commentId/replies", commentHandler.ListCommentReplies)
					commentsRoutes.PUT("/:commentId", commentHandler.UpdateComment)
					commentsRoutes.DELETE("/:commentId", commentHandler.DeleteComment)
					commentsRoutes.POST("/:commentId/like", commentHandler.ToggleLikeComment)
				}
			}
		}
	}
	return router
}
