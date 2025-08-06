package handlers

import (
	"errors"
	"log"
	"my-music-go/internal/models"
	"my-music-go/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArtistHandler struct {
	artistService *services.ArtistService
}

func NewArtistHandler(artistService *services.ArtistService) *ArtistHandler {
	return &ArtistHandler{artistService: artistService}
}

// GetArtistDetail 处理获取艺术家详情的请求
// @Summary      获取艺术家详情及歌曲列表
// @Description  根据艺术家ID获取其详细信息，并返回该艺术家的歌曲分页列表。歌曲列表支持搜索和排序。
// @Tags         艺术家 (Artist)
// @Produce      json
// @Security     BearerAuth
// @Param        id        path      int     true   "艺术家 ID"
// @Param        page      query     int     false  "歌曲列表的页码" default(1)
// @Param        pageSize  query     int     false  "每页的歌曲数量" default(10)
// @Param        search    query     string  false  "在歌曲中搜索的关键词 (匹配歌曲名或专辑名)"
// @Param        sortBy    query     string  false  "歌曲列表的排序字段" Enums(oldest, latest, play_count, like_count)
// @Success      200       {object}  models.ArtistDetailResponseDTO
// @Failure      400       {object}  map[string]string "{"error": "无效的ID或查询参数"}"
// @Failure      401       {object}  map[string]string "{"error": "需要认证"}"
// @Failure      404       {object}  map[string]string "{"error": "艺术家未找到"}"
// @Failure      500       {object}  map[string]string "{"error": "获取艺术家详情失败"}"
// @Router       /artists/{id} [get]
func (h *ArtistHandler) GetArtistDetail(c *gin.Context) {
	// 1. 解析路径参数 artistID
	artistIDStr := c.Param("id")
	artistID, err := strconv.ParseInt(artistIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的艺术家ID"})
		return
	}

	// 2. 绑定歌曲列表的查询参数 (page, pageSize, etc.)
	var songParams models.ListSongsRequestDTO
	if err := c.ShouldBindQuery(&songParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的查询参数"})
		return
	}

	// 3. 调用 Service
	response, err := h.artistService.GetArtistDetail(artistID, &songParams)
	if err != nil {
		// 假设你在 services/errors.go 中定义了 ErrArtistNotFound
		if errors.Is(err, services.ErrArtistNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "艺术家未找到"})
		} else {
			log.Printf("获取艺术家详情失败: %+v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "获取艺术家详情失败，请稍后重试"})
		}
		return
	}

	// 4. 成功响应
	c.JSON(http.StatusOK, response)
}
