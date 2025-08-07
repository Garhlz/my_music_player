package services

import (
	"fmt"
	"my-music-go/internal/models"
	"my-music-go/internal/repository"
	"sync"
)

type SongService struct {
	songRepo repository.ISongRepository
}

func NewSongService(songRepo repository.ISongRepository) *SongService {
	return &SongService{songRepo: songRepo}
}

func (s *SongService) GetSongDetail(songID int64) (*models.SongDetailDTO, error) {
	song, err := s.songRepo.FindDetailByID(songID)
	if err != nil {
		return nil, fmt.Errorf("database error on find song detail: %w", err)
	}
	if song == nil {
		return nil, ErrSongNotFound
	}

	return song, nil
}

// ListSongs 使用并发来提升性能, 经典的分页查询...
func (s *SongService) ListSongs(params *models.ListSongsRequestDTO) (*models.PaginatedResponseDTO, error) {
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

	// 任务1：获取歌曲列表
	go func() {
		defer wg.Done()
		songs, errSongs = s.songRepo.List(params)
	}()

	// 任务2：获取总数
	go func() {
		defer wg.Done()
		total, errTotal = s.songRepo.Count(params)
	}()

	wg.Wait() // 等待两个任务都完成

	if errSongs != nil {
		return nil, fmt.Errorf("failed to list songs: %w", errSongs)
	}
	if errTotal != nil {
		return nil, fmt.Errorf("failed to count songs: %w", errTotal)
	}

	return &models.PaginatedResponseDTO{
		Total: total,
		List:  songs,
	}, nil
}
