package handlers

import (
	"errors"
	"log"
	"my-music-go/internal/models"
	"my-music-go/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	commentService *services.CommentService
}

func NewCommentHandler(commentService *services.CommentService) *CommentHandler {
	return &CommentHandler{commentService: commentService}
}

// ListSongComments 获取歌曲的根评论列表
// @Summary 获取歌曲评论
// @Description 获取一首歌的根评论（非回复）的分页列表。
// @Tags  评论 (Comment)
// @Produce json
// @Param id path int true "歌曲 ID"
// @Param page query int false "页码" default(1)
// @Param pageSize query int false "每页数量" default(10)
// @Success 200 {object} models.PaginatedResponseDTO{List=[]models.CommentDTO}
// @Router /songs/{id}/comments [get]
func (h *CommentHandler) ListSongComments(c *gin.Context) {
	songID, err := GetIDFromParam(c, "id")
	if err != nil {
		return
	}

	var params models.ListCommentsRequestDTO
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的查询参数"})
		return
	}

	// 尝试从 context 获取 userID，如果不存在（未登录），则为0
	// 这样 Service 层就可以根据 userID 是否 > 0 来判断是否要查询“我是否已点赞”
	var currentUserID int64
	if useridAny, ok := c.Get("userID"); ok {
		currentUserID = useridAny.(int64)
	}

	response, err := h.commentService.ListSongComments(songID, currentUserID, &params)
	if err != nil {
		log.Printf("获取评论列表失败: %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取评论列表失败"})
		return
	}
	c.JSON(http.StatusOK, response)
}

// ListCommentReplies 获取评论的回复列表
// @Summary 获取评论回复
// @Description 根据一条父评论的ID，获取它的所有回复列表（不分页）。
// @Tags  评论 (Comment)
// @Produce json
// @Param commentId path int true "父评论 ID"
// @Success 200 {object} []models.CommentDTO
// @Router /comments/{commentId}/replies [get]
func (h *CommentHandler) ListCommentReplies(c *gin.Context) {
	parentID, err := GetIDFromParam(c, "commentId")
	if err != nil {
		return
	}

	var currentUserID int64
	if useridAny, ok := c.Get("userID"); ok {
		currentUserID = useridAny.(int64)
	}

	response, err := h.commentService.ListCommentReplies(parentID, currentUserID)
	if err != nil {
		log.Printf("获取评论回复列表失败: %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取回复列表失败"})
		return
	}
	c.JSON(http.StatusOK, response)
}

// CreateComment 发表新评论或回复
// @Summary 发表评论
// @Description 为一首歌发表新评论，或回复一条已有的评论。
// @Tags  评论 (Comment)
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "歌曲 ID"
// @Param createRequest body models.CreateCommentRequestDTO true "评论内容"
// @Success 201 {object} models.CommentDTO
// @Router /songs/{id}/comments [post]
func (h *CommentHandler) CreateComment(c *gin.Context) {
	authUserID := c.MustGet("userID").(int64)
	songID, err := GetIDFromParam(c, "id")
	if err != nil {
		return
	}

	var req models.CreateCommentRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数: " + err.Error()})
		return
	}

	comment, err := h.commentService.CreateComment(authUserID, songID, &req)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrSongNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": "评论的歌曲不存在"})
		case errors.Is(err, services.ErrCommentNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": "回复的评论不存在"})
		case errors.Is(err, services.ErrBadRequest):
			c.JSON(http.StatusBadRequest, gin.H{"error": "不能回复其他歌曲下的评论"})
		default:
			log.Printf("发表评论失败: %+v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "发表评论失败"})
		}
		return
	}
	c.JSON(http.StatusCreated, comment)
}

// UpdateComment 更新评论
// @Summary 更新我的评论
// @Description 更新当前认证用户发表的一条评论。
// @Tags 评论 (Comment)
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param commentId path int true "评论 ID"
// @Param updateRequest body models.UpdateCommentRequestDTO true "更新的评论内容"
// @Success 200 {object} map[string]string "{"message": "更新成功"}"
// @Router /comments/{commentId} [put]
func (h *CommentHandler) UpdateComment(c *gin.Context) {
	authUserID := c.MustGet("userID").(int64)
	commentID, err := GetIDFromParam(c, "commentId")
	if err != nil {
		return
	}

	var req models.UpdateCommentRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数: " + err.Error()})
		return
	}

	err = h.commentService.UpdateComment(authUserID, commentID, &req)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrCommentNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": "评论不存在"})
		case errors.Is(err, services.ErrForbidden):
			c.JSON(http.StatusForbidden, gin.H{"error": "无权修改此评论"})
		default:
			log.Printf("更新评论失败: %+v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "更新失败"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "更新成功"})
}

// DeleteComment 删除评论
// @Summary 删除我的评论
// @Description 删除当前认证用户发表的一条评论（及其所有回复）。
// @Tags  评论 (Comment)
// @Produce json
// @Security BearerAuth
// @Param commentId path int true "评论 ID"
// @Success 200 {object} map[string]string "{"message": "删除成功"}"
// @Router /comments/{commentId} [delete]
func (h *CommentHandler) DeleteComment(c *gin.Context) {
	authUserID := c.MustGet("userID").(int64)
	commentID, err := GetIDFromParam(c, "commentId")
	if err != nil {
		return
	}

	err = h.commentService.DeleteComment(authUserID, commentID)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrCommentNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": "评论不存在"})
		case errors.Is(err, services.ErrForbidden):
			c.JSON(http.StatusForbidden, gin.H{"error": "无权删除此评论"})
		default:
			log.Printf("删除评论失败: %+v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "删除失败"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

// ToggleLikeComment 点赞/取消点赞评论
// @Summary 点赞或取消点赞评论
// @Description 对一条评论进行点赞。如果已经点赞，则取消点赞。
// @Tags  评论 (Comment)
// @Produce json
// @Security BearerAuth
// @Param commentId path int true "评论 ID"
// @Success 200 {object} models.LikeCommentResponseDTO
// @Router /comments/{commentId}/like [post]
func (h *CommentHandler) ToggleLikeComment(c *gin.Context) {
	authUserID := c.MustGet("userID").(int64)
	commentID, err := GetIDFromParam(c, "commentId")
	if err != nil {
		return
	}

	response, err := h.commentService.ToggleLikeComment(authUserID, commentID)
	if err != nil {
		if errors.Is(err, services.ErrCommentNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "评论不存在"})
		} else {
			log.Printf("点赞评论失败: %+v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "操作失败"})
		}
		return
	}
	c.JSON(http.StatusOK, response)
}
