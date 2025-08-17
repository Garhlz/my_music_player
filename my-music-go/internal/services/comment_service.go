package services

import (
	"fmt"
	"my-music-go/internal/models"
	"my-music-go/internal/repository"
	"sync"
	"time"
)

type CommentService struct {
	commentRepo repository.ICommentRepository
	songRepo    repository.ISongRepository
}

func NewCommentService(commentRepo repository.ICommentRepository, songRepo repository.ISongRepository) *CommentService {
	return &CommentService{commentRepo: commentRepo, songRepo: songRepo}
}

// checkCommentOwner 权限检查辅助函数
func (s *CommentService) checkCommentOwner(userID, commentID int64) (*models.CommentDTO, error) {
	comment, err := s.commentRepo.FindDTOByID(commentID, &userID)
	if err != nil {
		return nil, fmt.Errorf("db error on find comment: %w", err)
	}
	if comment == nil {
		return nil, ErrCommentNotFound
	}
	if comment.UserID != userID {
		return nil, ErrForbidden
	}
	return comment, nil
}

// ListSongComments 获取歌曲的根评论列表 (并发)
func (s *CommentService) ListSongComments(songID, currentUserID int64, params *models.ListCommentsRequestDTO) (*models.PaginatedResponseDTO, error) {
	// 和获取歌曲列表一样，并发地获取列表和总数，以提升性能。
	if params.Page <= 0 {
		params.Page = 1
	}
	if params.PageSize <= 0 {
		params.PageSize = 10
	}

	var wg sync.WaitGroup
	var comments []models.CommentDTO
	var total int
	var errComments, errTotal error

	var currentUserIDPtr *int64
	if currentUserID > 0 {
		currentUserIDPtr = &currentUserID
	}
	wg.Add(2)
	go func() {
		defer wg.Done()
		comments, errComments = s.commentRepo.ListBySongID(songID, currentUserIDPtr, params)
	}()
	go func() {
		defer wg.Done()
		total, errTotal = s.commentRepo.CountBySongID(songID)
	}()
	wg.Wait()

	if errComments != nil {
		return nil, fmt.Errorf("failed to list comments: %w", errComments)
	}
	if errTotal != nil {
		return nil, fmt.Errorf("failed to count comments: %w", errTotal)
	}

	return &models.PaginatedResponseDTO{Total: total, List: comments}, nil
}

// ListCommentReplies 获取评论的回复列表
func (s *CommentService) ListCommentReplies(parentID, currentUserID int64) ([]models.CommentDTO, error) {
	var currentUserIDPtr *int64
	if currentUserID > 0 {
		currentUserIDPtr = &currentUserID
	}
	return s.commentRepo.ListByParentID(parentID, currentUserIDPtr)
}

// CreateComment 发表评论
func (s *CommentService) CreateComment(userID, songID int64, req *models.CreateCommentRequestDTO) (*models.CommentDTO, error) {
	song, err := s.songRepo.FindDetailByID(songID)
	if err != nil {
		return nil, fmt.Errorf("db error on find song: %w", err)
	}
	if song == nil {
		return nil, ErrSongNotFound
	}

	var currentUserIDPtr *int64
	if userID > 0 {
		currentUserIDPtr = &userID
	}
	if req.ParentID != nil {
		parentComment, err := s.commentRepo.FindDTOByID(*req.ParentID, currentUserIDPtr)
		if err != nil {
			return nil, fmt.Errorf("db error on find parent comment: %w", err)
		}
		if parentComment == nil {
			return nil, ErrCommentNotFound
		}
		if parentComment.SongID != songID {
			return nil, ErrBadRequest
		}
	}

	isRoot := req.ParentID == nil || *req.ParentID == 0
	comment := &models.Comment{
		UserID:        userID,
		SongID:        songID,
		Text:          req.Text,
		ParentID:      req.ParentID,
		ReplyToUserID: req.ReplyToUserID,
		IsRoot:        isRoot,
	}

	newID, err := s.commentRepo.Create(comment)
	if err != nil {
		return nil, fmt.Errorf("failed to create comment: %w", err)
	}

	return s.commentRepo.FindDTOByID(newID, currentUserIDPtr)
}

// UpdateComment 更新评论
func (s *CommentService) UpdateComment(userID, commentID int64, req *models.UpdateCommentRequestDTO) error {
	comment, err := s.checkCommentOwner(userID, commentID)
	if err != nil {
		return err
	}

	comment.Text = req.Text
	comment.UpdatedAt = time.Now()

	return s.commentRepo.Update(&comment.Comment)
}

// DeleteComment 删除评论
func (s *CommentService) DeleteComment(userID, commentID int64) error {
	if _, err := s.checkCommentOwner(userID, commentID); err != nil {
		return err
	}
	return s.commentRepo.Delete(commentID)
}

// ToggleLikeComment 点赞/取消点赞评论
func (s *CommentService) ToggleLikeComment(userID, commentID int64) (*models.LikeCommentResponseDTO, error) {
	// 检查评论是否存在
	var currentUserIDPtr *int64
	if userID > 0 {
		currentUserIDPtr = &userID
	}
	comment, err := s.commentRepo.FindDTOByID(commentID, currentUserIDPtr)
	if err != nil {
		return nil, fmt.Errorf("db error on find comment: %w", err)
	}
	if comment == nil {
		return nil, ErrCommentNotFound
	}

	isLiked, err := s.commentRepo.ToggleLike(userID, commentID)
	if err != nil {
		return nil, fmt.Errorf("db error on toggle like: %w", err)
	}

	likeCount, err := s.commentRepo.GetLikeCount(commentID)
	if err != nil {
		return nil, fmt.Errorf("db error on get like count: %w", err)
	}

	return &models.LikeCommentResponseDTO{
		LikeCount:   likeCount,
		IsLikedByMe: isLiked,
	}, nil
}

func (s *CommentService) GetLikeCount(commentID int64) (*int, error) {
	// 检查评论是否存在
	var currentUserIDPtr *int64
	comment, err := s.commentRepo.FindDTOByID(commentID, currentUserIDPtr)
	if err != nil {
		return nil, fmt.Errorf("db error on find comment: %w", err)
	}
	if comment == nil {
		return nil, ErrCommentNotFound
	}

	likeCount, err := s.commentRepo.GetLikeCount(commentID)
	if err != nil {
		return nil, fmt.Errorf("db error on get like count: %w", err)
	}

	return &likeCount, err
}
