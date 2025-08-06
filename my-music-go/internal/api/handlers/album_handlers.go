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

type AlbumHandler struct {
	albumService *services.AlbumService
}

func NewAlbumHandler(albumService *services.AlbumService) *AlbumHandler {
	return &AlbumHandler{albumService: albumService}
}

// ListAlbums 获取专辑列表
// @Summary      获取专辑列表
// @Description  获取专辑的分页列表，支持按专辑名或艺术家名进行搜索，并可按指定字段排序。
// @Tags         专辑 (Album)
// @Produce      json
// @Security     BearerAuth
// @Param        page         query     int     false  "页码" default(1)
// @Param        pageSize     query     int     false  "每页数量" default(10)
// @Param        search       query     string  false  "搜索关键词 (匹配专辑名或艺术家名)"
// @Param        sortBy       query     string  false  "排序字段 (latest, release_date)" Enums(latest, release_date)
// @Success      200          {object}  models.PaginatedResponseDTO{List=[]models.AlbumListDTO}
// @Failure      400          {object}  map[string]string "{"error": "无效的查询参数"}"
// @Failure      401          {object}  map[string]string "{"error": "需要认证"}"
// @Failure      500          {object}  map[string]string "{"error": "获取专辑列表失败"}"
// @Router       /albums [get]
func (h *AlbumHandler) ListAlbums(c *gin.Context) {
	var params models.ListAlbumsRequestDTO
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的查询参数"})
		return
	}

	response, err := h.albumService.ListAlbums(&params)
	if err != nil {
		log.Printf("获取专辑列表失败: %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取专辑列表失败"})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetAlbumDetail 获取专辑详情
// @Summary      获取专辑详情及歌曲列表
// @Description  根据专辑ID获取其详细信息（包括艺术家），并返回该专辑包含的歌曲分页列表。
// @Tags         专辑 (Album)
// @Produce      json
// @Security     BearerAuth
// @Param        id        path      int     true   "专辑 ID"
// @Param        page      query     int     false  "歌曲列表的页码" default(1)
// @Param        pageSize  query     int     false  "每页的歌曲数量" default(10)
// @Param        search    query     string  false  "在歌曲中搜索的关键词 (匹配歌曲名或艺术家名)"
// @Param        sortBy    query     string  false  "歌曲列表的排序字段 (latest, play_count, like_count)" Enums(latest, play_count, like_count)
// @Success      200       {object}  models.AlbumDetailDTO
// @Failure      400       {object}  map[string]string "{"error": "无效的ID或查询参数"}"
// @Failure      401       {object}  map[string]string "{"error": "需要认证"}"
// @Failure      404       {object}  map[string]string "{"error": "专辑未找到"}"
// @Failure      500       {object}  map[string]string "{"error": "获取专辑详情失败"}"
// @Router       /albums/{id} [get]
func (h *AlbumHandler) GetAlbumDetail(c *gin.Context) {
	albumIDStr := c.Param("id")
	albumID, err := strconv.ParseInt(albumIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的专辑ID"})
		return
	}

	var songParams models.ListSongsRequestDTO
	if err := c.ShouldBindQuery(&songParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的查询参数"})
		return
	}

	response, err := h.albumService.GetAlbumDetail(albumID, &songParams)
	if err != nil {
		// 定义了 ErrAlbumNotFound
		if errors.Is(err, services.ErrAlbumNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "专辑未找到"})
		} else {
			log.Printf("获取专辑详情失败: %+v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "获取专辑详情失败，请稍后重试"})
		}
		return
	}

	// 4. 成功响应
	c.JSON(http.StatusOK, response)
}
