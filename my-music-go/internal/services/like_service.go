package services

import (
	"fmt"
	"my-music-go/internal/models"
	"my-music-go/internal/repository"
	"sync"
)

type LikeService struct {
	likeRepo repository.ILikeRepository
	songRepo repository.ISongRepository // 依赖 songRepo 来检查歌曲是否存在
}

func NewLikeService(likeRepo repository.ILikeRepository, songRepo repository.ISongRepository) *LikeService {
	return &LikeService{likeRepo: likeRepo, songRepo: songRepo}
}

// ListLikedSongs 使用并发来提升性能
func (s *LikeService) ListLikedSongs(userID int64, params *models.ListLikedSongsRequestDTO) (*models.PaginatedResponseDTO, error) {
	// 设置分页默认值
	if params.Page <= 0 {
		params.Page = 1
	}
	if params.PageSize <= 0 {
		params.PageSize = 10
	}

	var wg sync.WaitGroup
	var songs []models.SongDetailDTO
	var total int
	var errSongs, errTotal error

	wg.Add(2) // 我们有两个并发任务

	// 1. 获取歌曲列表
	go func() {
		defer wg.Done()
		songs, errSongs = s.likeRepo.List(userID, params)
	}()

	// 2. 获取总数
	go func() {
		defer wg.Done()
		total, errTotal = s.likeRepo.Count(userID, params)
	}()

	wg.Wait() // 等待两个任务都完成

	if errSongs != nil {
		return nil, fmt.Errorf("failed to list liked songs: %w", errSongs)
	}
	if errTotal != nil {
		return nil, fmt.Errorf("failed to count liked songs: %w", errTotal)
	}

	return &models.PaginatedResponseDTO{
		Total: total,
		List:  songs,
	}, nil
}

func (s *LikeService) LikeSong(userID, songID int64) error {
	// 1. 检查歌曲是否存在
	song, err := s.songRepo.FindDetailByID(songID)
	if err != nil {
		return fmt.Errorf("database error on find song: %w", err)
	}
	if song == nil {
		return ErrSongNotFound
	}

	// 2. 检查是否已经喜欢过
	liked, err := s.likeRepo.IsLiked(userID, songID)
	if err != nil {
		return fmt.Errorf("database error on check if liked: %w", err)
	}
	if liked {
		return ErrSongAlreadyLiked // 或者也可以直接返回成功，取决于产品逻辑
	}

	// 3. 创建喜欢记录
	return s.likeRepo.Create(userID, songID)
}

func (s *LikeService) UnlikeSong(userID, songID int64) error {
	rowsAffected, err := s.likeRepo.Delete(userID, songID)
	if err != nil {
		return fmt.Errorf("database error on delete like: %w", err)
	}
	if rowsAffected == 0 {
		return ErrLikeNotFound // 告诉 Handler，没有找到对应的点赞记录
	}
	return nil
}
