package handlers

import (
	"errors"
	"log"
	"my-music-go/internal/models"
	"my-music-go/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SongHandler struct {
	songService *services.SongService
}

func NewSongHandler(songService *services.SongService) *SongHandler {
	return &SongHandler{songService: songService}
}

// GetSongDetail 获取单曲详情
// @Summary      获取单曲详情
// @Description  根据歌曲ID获取单曲的详细信息，包括艺术家和专辑数据。
// @Tags         歌曲 (Song)
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "歌曲 ID"
// @Success      200  {object}  models.SongDetailDTO
// @Failure      400  {object}  map[string]string "{"error": "无效的歌曲ID格式"}"
// @Failure      401  {object}  map[string]string "{"error": "需要认证"}"
// @Failure      404  {object}  map[string]string "{"error": "歌曲未找到"}"
// @Failure      500  {object}  map[string]string "{"error": "获取歌曲详情失败，请稍后重试"}"
// @Router       /songs/{id} [get]
func (h *SongHandler) GetSongDetail(c *gin.Context) {
	songID, err := GetIDFromParam(c, "id")
	if err != nil {
		return
	}

	songDetail, err := h.songService.GetSongDetail(songID)
	if err != nil {
		if errors.Is(err, services.ErrSongNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "歌曲未找到"})
		} else {
			log.Printf("获取歌曲详情失败: %+v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "获取歌曲详情失败，请稍后重试"})
		}
		return
	}

	c.JSON(http.StatusOK, songDetail)
}

// ListSongs 获取歌曲列表（分页、搜索、排序）
// @Summary      获取歌曲列表
// @Description  获取公开的歌曲列表，支持分页、按歌曲名或艺术家名搜索、以及排序。
// @Tags         歌曲 (Song)
// @Produce      json
// @Security     BearerAuth
// @Param        page      query     int     false  "页码" default(1)
// @Param        pageSize  query     int     false  "每页数量" default(10)
// @Param        search    query     string  false  "搜索关键词 (匹配歌曲名或艺术家)"
// @Param        sortBy    query     string  false  "排序字段 (oldest, latest, play_count, like_count)" Enums(oldest,latest, play_count, like_count)
// @Success      200       {object}  models.PaginatedResponseDTO{List=[]models.SongDetailDTO}
// @Failure      400       {object}  map[string]string "{"error": "无效的查询参数"}"
// @Failure      401       {object}  map[string]string "{"error": "需要认证"}"
// @Failure      500       {object}  map[string]string "{"error": "获取歌曲列表失败"}"
// @Router       /songs [get]
func (h *SongHandler) ListSongs(c *gin.Context) {
	var params models.ListSongsRequestDTO
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的查询参数"})
		return
	}

	response, err := h.songService.ListSongs(&params)
	if err != nil {
		log.Printf("获取歌曲列表失败: %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取歌曲列表失败"})
		return
	}

	c.JSON(http.StatusOK, response)
}
