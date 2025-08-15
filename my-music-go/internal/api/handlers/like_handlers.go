package handlers

import (
	"errors"
	"log"
	"my-music-go/internal/models"
	"my-music-go/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LikeHandler struct {
	likeService *services.LikeService
}

func NewLikeHandler(likeService *services.LikeService) *LikeHandler {
	return &LikeHandler{likeService: likeService}
}

// ListMyLikedSongs 获取当前登录用户的喜欢列表
// @Summary      获取我喜欢的歌曲列表
// @Description  获取当前认证用户（通过JWT Token识别）的喜欢歌曲分页列表，支持搜索和排序。
// @Tags         喜欢 (Like)
// @Produce      json
// @Security     BearerAuth
// @Param        page      query     int     false  "页码" default(1)
// @Param        pageSize  query     int     false  "每页数量" default(10)
// @Param        search    query     string  false  "搜索关键词 (匹配歌曲名或艺术家名)"
// @Param        sortBy    query     string  false  "排序字段 (latest, oldest, name, duration)" Enums(latest, oldest, name, duration)
// @Success      200       {object}  models.PaginatedResponseDTO{List=[]models.SongDetailDTO}
// @Failure      400       {object}  map[string]string "{"error": "无效的查询参数"}"
// @Failure      401       {object}  map[string]string "{"error": "需要认证"}"
// @Failure      500       {object}  map[string]string "{"error": "获取列表失败"}"
// @Router       /me/liked-songs [get]
func (h *LikeHandler) ListMyLikedSongs(c *gin.Context) {
	authUserID := c.MustGet("userID").(int64)

	var params models.ListLikedSongsRequestDTO
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的查询参数"})
		return
	}

	response, err := h.likeService.ListLikedSongs(authUserID, &params)
	if err != nil {
		log.Printf("获取用户喜欢列表失败: %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取列表失败"})
		return
	}

	c.JSON(http.StatusOK, response)
}

// ListUserLikedSongs 获取指定用户的喜欢列表
// @Summary      获取指定用户的喜欢列表
// @Description  获取指定用户的公开喜欢歌曲分页列表，支持搜索和排序。
// @Tags         喜欢 (Like)
// @Produce      json
// @Security     BearerAuth
// @Param        id        path      int     true   "用户 ID"
// @Param        page      query     int     false  "页码" default(1)
// @Param        pageSize  query     int     false  "每页数量" default(10)
// @Param        search    query     string  false  "搜索关键词 (匹配歌曲名或艺术家名)"
// @Param        sortBy    query     string  false  "排序字段 (latest, oldest, name, duration)" Enums(latest, oldest, name, duration)
// @Success      200       {object}  models.PaginatedResponseDTO{List=[]models.SongDetailDTO}
// @Failure      400       {object}  map[string]string "{"error": "无效的ID或查询参数"}"
// @Failure      401       {object}  map[string]string "{"error": "需要认证"}"
// @Failure      500       {object}  map[string]string "{"error": "获取喜欢的歌曲列表失败"}"
// @Router       /users/{id}/liked-songs [get]
func (h *LikeHandler) ListUserLikedSongs(c *gin.Context) {
	userID, err := GetIDFromParam(c, "id")
	if err != nil {
		// GetIDFromParam 内部已经处理了响应
		return
	}

	var params models.ListLikedSongsRequestDTO
	if err = c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的查询参数"})
		return
	}

	response, err := h.likeService.ListLikedSongs(userID, &params)
	if err != nil {
		log.Printf("获取喜欢的歌曲列表失败: %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取喜欢的歌曲列表失败"})
		return
	}

	c.JSON(http.StatusOK, response)
}

// LikeSong 添加喜欢歌曲
// @Summary      添加歌曲到我喜欢列表
// @Description  将一首指定的歌曲添加到当前认证用户的“我喜欢”列表中。
// @Tags         喜欢 (Like)
// @Produce      json
// @Security     BearerAuth
// @Param        songId   path      int  true  "歌曲 ID"
// @Success      201      {object}  map[string]string "{"message": "添加成功"}"
// @Failure      400      {object}  map[string]string "{"error": "无效的歌曲ID格式"}"
// @Failure      401      {object}  map[string]string "{"error": "需要认证"}"
// @Failure      404      {object}  map[string]string "{"error": "歌曲不存在"}"
// @Failure      409      {object}  map[string]string "{"error": "已经喜欢过这首歌了"}"
// @Failure      500      {object}  map[string]string "{"error": "操作失败，请稍后重试"}"
// @Router       /me/liked-songs/{songId} [post]
func (h *LikeHandler) LikeSong(c *gin.Context) {
	songID, err := GetIDFromParam(c, "songId")
	if err != nil {
		return
	}

	authUserID := c.MustGet("userID").(int64)

	err = h.likeService.LikeSong(authUserID, songID)
	if err != nil {
		if errors.Is(err, services.ErrSongNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "歌曲不存在"})
		} else if errors.Is(err, services.ErrSongAlreadyLiked) {
			c.JSON(http.StatusConflict, gin.H{"error": "已经喜欢过这首歌了"})
		} else {
			log.Printf("添加喜欢失败: %+v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "操作失败，请稍后重试"})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "添加成功"})
}

// UnlikeSong 取消喜欢歌曲
// @Summary      从我喜欢列表移除歌曲
// @Description  将一首指定的歌曲从当前认证用户的“我喜欢”列表中移除。
// @Tags         喜欢 (Like)
// @Produce      json
// @Security     BearerAuth
// @Param        songId   path      int  true  "歌曲 ID"
// @Success      200      {object}  map[string]string "{"message": "取消成功"}"
// @Failure      400      {object}  map[string]string "{"error": "无效的歌曲ID格式"}"
// @Failure      401      {object}  map[string]string "{"error": "需要认证"}"
// @Failure      404      {object}  map[string]string "{"error": "您尚未喜欢这首歌"}"
// @Failure      500      {object}  map[string]string "{"error": "操作失败，请稍后重试"}"
// @Router       /me/liked-songs/{songId} [delete]
func (h *LikeHandler) UnlikeSong(c *gin.Context) {
	songID, err := GetIDFromParam(c, "songId")
	if err != nil {
		return
	}
	authUserID := c.MustGet("userID").(int64)

	err = h.likeService.UnlikeSong(authUserID, songID)
	if err != nil {
		if errors.Is(err, services.ErrLikeNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "您尚未喜欢这首歌"})
		} else {
			log.Printf("取消喜欢失败: %+v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "操作失败，请稍后重试"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "取消成功"})
}

// GetMyLikedSongStatus 检查当前认证用户是否喜欢某首歌曲
// @Summary      检查歌曲喜欢状态
// @Description  根据歌曲ID，检查当前通过认证的用户是否已经喜欢了该歌曲。
// @Tags         喜欢 (Like)
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        songId   path      int  true  "歌曲 ID"
// @Success      200      {object}  models.LikeStatusResponse "成功返回喜欢状态"
// @Failure      401      {object}  map[string]string "{"error": "需要认证（未提供Token或Token无效）"}"
// @Failure      404      {object}  map[string]string "{"error": "歌曲不存在"}"
// @Failure      500      {object}  map[string]string "{"error": "服务器内部错误"}"
// @Router       /me/liked-songs/{songId}/status [get]
func (h *LikeHandler) GetMyLikedSongStatus(c *gin.Context) {
	authUserID := c.MustGet("userID").(int64)
	songID, err := GetIDFromParam(c, "songId")
	if err != nil {
		return
	}

	isLiked, err := h.likeService.GetLikedSongStatus(authUserID, songID)
	if err != nil {
		if errors.Is(err, services.ErrSongNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "歌曲不存在"})
		} else {
			log.Printf("获取当前用户喜欢状态失败: %+v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "操作失败，请稍后重试"})
		}
		return // 如果出错, 要停止执行...
	}

	c.JSON(http.StatusOK, gin.H{
		"isLiked": *isLiked,
	})
}
