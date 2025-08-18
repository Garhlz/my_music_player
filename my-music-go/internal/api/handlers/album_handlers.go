package handlers

import (
	"errors"
	"log"
	"my-music-go/internal/models"
	"my-music-go/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AlbumHandler struct {
	albumService *services.AlbumService
	songService  *services.SongService
}

func NewAlbumHandler(albumService *services.AlbumService, songService *services.SongService) *AlbumHandler {
	return &AlbumHandler{albumService: albumService, songService: songService}
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

// GetAlbumInfo 获取单张专辑的详细信息
// @Summary      获取专辑详情
// @Description  根据专辑ID获取其详细信息。
// @Tags         专辑 (Album)
// @Produce      json
// @Security     BearerAuth
// @Param        id        path      int     true   "专辑 ID"
// @Success      200       {object}  models.Album
// @Failure      400       {object}  map[string]string "{"error": "无效的专辑ID"}"
// @Failure      401       {object}  map[string]string "{"error": "需要认证"}"
// @Failure      404       {object}  map[string]string "{"error": "专辑未找到"}"
// @Failure      500       {object}  map[string]string "{"error": "获取专辑失败"}"
// @Router       /albums/{id} [get]
func (h *AlbumHandler) GetAlbumInfo(c *gin.Context) {
	albumID, err := GetIDFromParam(c, "id")
	if err != nil {
		return
	}

	artist, err := h.albumService.GetAlbum(albumID)
	if err != nil {
		if errors.Is(err, services.ErrAlbumNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "专辑未找到"})
		} else {
			log.Printf("获取专辑失败: %+v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "获取专辑失败"})
		}
		return
	}

	c.JSON(http.StatusOK, artist)
}

// ListSongsByAlbum 获取指定专辑的歌曲分页列表
// @Summary      获取专辑的歌曲列表
// @Description  根据专辑ID获取其歌曲的分页列表，支持搜索和排序。
// @Tags         专辑 (Album)
// @Produce      json
// @Security     BearerAuth
// @Param        id        path      int     true   "专辑 ID"
// @Param        page      query     int     false  "歌曲列表的页码" default(1)
// @Param        pageSize  query     int     false  "每页的歌曲数量" default(10)
// @Param        search    query     string  false  "在歌曲中搜索的关键词 (匹配歌曲名或专辑名)"
// @Param        sortBy    query     string  false  "歌曲列表的排序字段" Enums(oldest, latest, play_count, like_count)
// @Success      200       {object}  models.PaginatedResponseDTO{List=[]models.SongDetailDTO}
// @Failure      400       {object}  map[string]string "{"error": "无效的ID或查询参数"}"
// @Failure      401       {object}  map[string]string "{"error": "需要认证"}"
// @Failure      500       {object}  map[string]string "{"error": "获取歌曲列表失败"}"
// @Router       /albums/{id}/songs [get]
func (h *AlbumHandler) ListSongsByAlbum(c *gin.Context) {
	// TODO 其实所有获取艺术家的地方, 都需要考虑有多个艺术家的情况...这里暂且不管
	// 1. 解析路径参数 albumID
	albumID, err := GetIDFromParam(c, "id")
	if err != nil {
		return
	}

	// 2. 绑定歌曲列表的查询参数 (page, pageSize, etc.)
	var songParams models.ListSongsRequestDTO
	if err := c.ShouldBindQuery(&songParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的查询参数"})
		return
	}

	// 3. 将从路径中获取的 albumID 设置到查询参数中
	songParams.AlbumID = &albumID

	// 4. 调用 SongService 的方法
	response, err := h.songService.ListSongs(&songParams)
	if err != nil {
		log.Printf("获取专辑歌曲列表失败: %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取歌曲列表失败"})
		return
	}

	c.JSON(http.StatusOK, response)
}
