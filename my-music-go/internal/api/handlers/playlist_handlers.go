package handlers

import (
	"errors"
	"log"
	"my-music-go/internal/models"
	"my-music-go/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PlaylistHandler struct {
	playlistService *services.PlaylistService
}

func NewPlaylistHandler(playlistService *services.PlaylistService) *PlaylistHandler {
	return &PlaylistHandler{playlistService: playlistService}
}

// ListMyPlaylists 获取我自己的歌单列表
// @Summary 获取我的歌单列表
// @Tags  歌单 (Playlist)
// @Produce json
// @Security BearerAuth
// @Param page query int false "页码"
// @Param pageSize query int false "每页数量"
// @Success 200 {object} models.PaginatedResponseDTO{List=[]models.PlaylistInfoDTO}
// @Router /me/playlists [get]
func (h *PlaylistHandler) ListMyPlaylists(c *gin.Context) {
	authUserID := c.MustGet("userID").(int64)
	var params models.ListPlaylistsRequestDTO
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的查询参数"})
		return
	}
	response, err := h.playlistService.ListUserPlaylists(authUserID, &params)
	if err != nil {
		log.Printf("获取我的歌单列表失败: %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取列表失败"})
		return
	}
	c.JSON(http.StatusOK, response)
}

// ListUserPlaylists 获取指定用户的歌单列表
// @Summary 获取指定用户的歌单列表
// @Tags  歌单 (Playlist)
// @Produce json
// @Security BearerAuth
// @Param id path int true "用户 ID"
// @Param page query int false "页码"
// @Param pageSize query int false "每页数量"
// @Success 200 {object} models.PaginatedResponseDTO{List=[]models.PlaylistInfoDTO}
// @Router /users/{id}/playlists [get]
func (h *PlaylistHandler) ListUserPlaylists(c *gin.Context) {
	userID, err := GetIDFromParam(c, "id")
	if err != nil {
		return
	}
	var params models.ListPlaylistsRequestDTO
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的查询参数"})
		return
	}
	response, err := h.playlistService.ListUserPlaylists(userID, &params)
	if err != nil {
		log.Printf("获取用户歌单列表失败: %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取列表失败"})
		return
	}
	c.JSON(http.StatusOK, response)
}

// CreatePlaylist 创建一个新歌单
// @Summary 创建我的新歌单
// @Tags  歌单 (Playlist)
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param createRequest body models.CreatePlaylistRequestDTO true "创建歌单请求"
// @Success 201 {object} map[string]interface{} "{"id": 1, "message": "创建成功"}"
// @Router /me/playlists [post]
func (h *PlaylistHandler) CreatePlaylist(c *gin.Context) {
	authUserID := c.MustGet("userID").(int64)
	var req models.CreatePlaylistRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数: " + err.Error()})
		return
	}
	playlistID, err := h.playlistService.CreatePlaylist(authUserID, &req)
	if err != nil {
		log.Printf("创建歌单失败: %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建失败"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": playlistID, "message": "创建成功"})
}

// GetPlaylistDetail 获取歌单详情
// @Summary 获取歌单详情
// @Tags  歌单 (Playlist)
// @Produce json
// @Security BearerAuth
// @Param playlistId path int true "歌单 ID"
// @Param page query int false "歌曲列表页码"
// @Param pageSize query int false "歌曲列表每页数量"
// @Success 200 {object} models.PlaylistDetailDTO
// @Router /playlists/{playlistId} [get]
func (h *PlaylistHandler) GetPlaylistDetail(c *gin.Context) {
	playlistID, err := GetIDFromParam(c, "playlistId")
	if err != nil {
		return
	}
	var songParams models.ListSongsRequestDTO
	if err := c.ShouldBindQuery(&songParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的查询参数"})
		return
	}
	response, err := h.playlistService.GetPlaylistDetail(playlistID, &songParams)
	if err != nil {
		if errors.Is(err, services.ErrPlaylistNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "歌单未找到"})
		} else {
			log.Printf("获取歌单详情失败: %+v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "获取详情失败"})
		}
		return
	}
	c.JSON(http.StatusOK, response)
}

// UpdatePlaylist 更新歌单信息
// @Summary 更新我的歌单信息
// @Tags  歌单 (Playlist)
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param playlistId path int true "歌单 ID"
// @Param updateRequest body models.UpdatePlaylistRequestDTO true "更新歌单请求"
// @Success 200 {object} map[string]string "{"message": "更新成功"}"
// @Router /me/playlists/{playlistId} [put]
func (h *PlaylistHandler) UpdatePlaylist(c *gin.Context) {
	authUserID := c.MustGet("userID").(int64)
	playlistID, err := GetIDFromParam(c, "playlistId")
	if err != nil {
		return
	}
	var req models.UpdatePlaylistRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数: " + err.Error()})
		return
	}
	err = h.playlistService.UpdatePlaylist(authUserID, playlistID, &req)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrPlaylistNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": "歌单未找到"})
		case errors.Is(err, services.ErrForbidden):
			c.JSON(http.StatusForbidden, gin.H{"error": "无权修改此歌单"})
		default:
			log.Printf("更新歌单失败: %+v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "更新失败"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "更新成功"})
}

// DeletePlaylist 删除歌单
// @Summary 删除我的歌单
// @Tags  歌单 (Playlist)
// @Produce json
// @Security BearerAuth
// @Param playlistId path int true "歌单 ID"
// @Success 200 {object} map[string]string "{"message": "删除成功"}"
// @Router /me/playlists/{playlistId} [delete]
func (h *PlaylistHandler) DeletePlaylist(c *gin.Context) {
	authUserID := c.MustGet("userID").(int64)
	playlistID, err := GetIDFromParam(c, "playlistId")
	if err != nil {
		return
	}
	err = h.playlistService.DeletePlaylist(authUserID, playlistID)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrPlaylistNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": "歌单未找到"})
		case errors.Is(err, services.ErrForbidden):
			c.JSON(http.StatusForbidden, gin.H{"error": "无权删除此歌单"})
		default:
			log.Printf("删除歌单失败: %+v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "删除失败"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

// AddSongToPlaylist 向歌单中添加歌曲
// @Summary 向我的歌单添加歌曲
// @Tags  歌单 (Playlist)
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param playlistId path int true "歌单 ID"
// @Param addSongRequest body models.AddSongToPlaylistRequestDTO true "添加歌曲请求"
// @Success 201 {object} map[string]string "{"message": "添加成功"}"
// @Router /me/playlists/{playlistId}/songs [post]
func (h *PlaylistHandler) AddSongToPlaylist(c *gin.Context) {
	authUserID := c.MustGet("userID").(int64)
	playlistID, err := GetIDFromParam(c, "playlistId")
	if err != nil {
		return
	}
	var req models.AddSongToPlaylistRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数: " + err.Error()})
		return
	}
	err = h.playlistService.AddSongToPlaylist(authUserID, playlistID, req.SongID)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrPlaylistNotFound), errors.Is(err, services.ErrSongNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case errors.Is(err, services.ErrForbidden):
			c.JSON(http.StatusForbidden, gin.H{"error": "无权操作此歌单"})
		case errors.Is(err, services.ErrSongAlreadyInPlaylist):
			c.JSON(http.StatusConflict, gin.H{"error": "歌曲已在歌单中"})
		default:
			log.Printf("添加歌曲到歌单失败: %+v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "添加失败"})
		}
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "添加成功"})
}

// RemoveSongFromPlaylist 从歌单中移除歌曲
// @Summary 从我的歌单移除歌曲
// @Tags  歌单 (Playlist)
// @Produce json
// @Security BearerAuth
// @Param playlistId path int true "歌单 ID"
// @Param songId path int true "歌曲 ID"
// @Success 200 {object} map[string]string "{"message": "移除成功"}"
// @Router /me/playlists/{playlistId}/songs/{songId} [delete]
func (h *PlaylistHandler) RemoveSongFromPlaylist(c *gin.Context) {
	authUserID := c.MustGet("userID").(int64)
	playlistID, err := GetIDFromParam(c, "playlistId")
	if err != nil {
		return
	}
	songID, err := GetIDFromParam(c, "songId")
	if err != nil {
		return
	}
	err = h.playlistService.RemoveSongFromPlaylist(authUserID, playlistID, songID)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrPlaylistNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": "歌单未找到"})
		case errors.Is(err, services.ErrForbidden):
			c.JSON(http.StatusForbidden, gin.H{"error": "无权操作此歌单"})
		case errors.Is(err, services.ErrSongNotInPlaylist):
			c.JSON(http.StatusNotFound, gin.H{"error": "歌曲不在此歌单中"})
		default:
			log.Printf("从歌单移除歌曲失败: %+v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "移除失败"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "移除成功"})
}
