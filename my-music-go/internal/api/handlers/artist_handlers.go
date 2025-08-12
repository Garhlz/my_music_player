package handlers

import (
	"errors"
	"log"
	"my-music-go/internal/models"
	"my-music-go/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ArtistHandler 保持不变
type ArtistHandler struct {
	artistService *services.ArtistService
	songService   *services.SongService // 增加了对 SongService 的依赖
}

// NewArtistHandler 构造函数需要更新
func NewArtistHandler(artistService *services.ArtistService, songService *services.SongService) *ArtistHandler {
	return &ArtistHandler{
		artistService: artistService,
		songService:   songService, // 注入 SongService
	}
}

// GetArtistInfo 获取单个艺术家的详细信息
// @Summary      获取艺术家详情
// @Description  根据艺术家ID获取其详细信息。
// @Tags         艺术家 (Artist)
// @Produce      json
// @Security     BearerAuth
// @Param        id        path      int     true   "艺术家 ID"
// @Success      200       {object}  models.Artist
// @Failure      400       {object}  map[string]string "{"error": "无效的艺术家ID"}"
// @Failure      401       {object}  map[string]string "{"error": "需要认证"}"
// @Failure      404       {object}  map[string]string "{"error": "艺术家未找到"}"
// @Failure      500       {object}  map[string]string "{"error": "获取艺术家失败"}"
// @Router       /artists/{id} [get]
func (h *ArtistHandler) GetArtistInfo(c *gin.Context) {
	artistID, err := GetIDFromParam(c, "id")
	if err != nil {
		return
	}
	artist, err := h.artistService.GetArtist(artistID)
	if err != nil {
		if errors.Is(err, services.ErrArtistNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "艺术家未找到"})
		} else {
			log.Printf("获取艺术家失败: %+v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "获取艺术家失败"})
		}
		return
	}

	c.JSON(http.StatusOK, artist)
}

// ListSongsByArtist 获取指定艺术家的歌曲分页列表
// @Summary      获取艺术家的歌曲列表
// @Description  根据艺术家ID获取其歌曲的分页列表，支持搜索和排序。
// @Tags         艺术家 (Artist)
// @Produce      json
// @Security     BearerAuth
// @Param        id        path      int     true   "艺术家 ID"
// @Param        page      query     int     false  "歌曲列表的页码" default(1)
// @Param        pageSize  query     int     false  "每页的歌曲数量" default(10)
// @Param        search    query     string  false  "在歌曲中搜索的关键词 (匹配歌曲名或专辑名)"
// @Param        sortBy    query     string  false  "歌曲列表的排序字段" Enums(oldest, latest, play_count, like_count)
// @Success      200       {object}  models.PaginatedResponseDTO{List=[]models.SongDetailDTO}
// @Failure      400       {object}  map[string]string "{"error": "无效的ID或查询参数"}"
// @Failure      401       {object}  map[string]string "{"error": "需要认证"}"
// @Failure      500       {object}  map[string]string "{"error": "获取歌曲列表失败"}"
// @Router       /artists/{id}/songs [get]
func (h *ArtistHandler) ListSongsByArtist(c *gin.Context) {
	// 1. 解析路径参数 artistID
	artistID, err := GetIDFromParam(c, "id")
	if err != nil {
		return
	}

	// 2. 绑定歌曲列表的查询参数 (page, pageSize, etc.)
	var songParams models.ListSongsRequestDTO
	if err := c.ShouldBindQuery(&songParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的查询参数"})
		return
	}

	// 3. 将从路径中获取的 artistID 设置到查询参数中
	songParams.ArtistID = &artistID

	// 4. 调用 SongService 的方法
	response, err := h.songService.ListSongs(&songParams)
	if err != nil {
		log.Printf("获取艺术家歌曲列表失败: %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取歌曲列表失败"})
		return
	}

	c.JSON(http.StatusOK, response)
}
