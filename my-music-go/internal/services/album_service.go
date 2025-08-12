package services

import (
	"database/sql"
	"errors"
	"fmt"
	"my-music-go/internal/models"
	"my-music-go/internal/repository"
	"sync"
)

// SongService 是否依赖? 是的, 需要知道歌曲数量
type AlbumService struct {
	albumRepo repository.IAlbumRepository // 注意这里是接口, 而且不是指针

}

// NewAlbumService 构造函数
func NewAlbumService(albumRepo repository.IAlbumRepository) *AlbumService {
	return &AlbumService{
		albumRepo: albumRepo,
	}
}

func (s *AlbumService) ListAlbums(albumParams *models.ListAlbumsRequestDTO) (*models.PaginatedResponseDTO, error) {
	// 设置分页默认值
	if albumParams.Page <= 0 {
		albumParams.Page = 1
	}
	if albumParams.PageSize <= 0 {
		albumParams.PageSize = 10
	}

	var wg sync.WaitGroup
	var albums []models.AlbumListDTO
	var total int
	var errAlbums, errTotal error

	wg.Add(2) // 我们有两个并发任务

	// 任务1：获取专辑列表
	go func() {
		defer wg.Done()
		albums, errAlbums = s.albumRepo.List(albumParams)
	}()

	// 任务2：获取总数
	go func() {
		defer wg.Done()
		total, errTotal = s.albumRepo.Count(albumParams)
	}()

	wg.Wait() // 等待两个任务都完成

	if errAlbums != nil {
		return nil, fmt.Errorf("failed to list albums: %w", errAlbums)
	}
	if errTotal != nil {
		return nil, fmt.Errorf("failed to count albums: %w", errTotal)
	}

	return &models.PaginatedResponseDTO{
		Total: total,
		List:  albums,
	}, nil
}

func (s *AlbumService) GetAlbum(albumID int64) (*models.Album, error) {
	albumInfo, err := s.albumRepo.FindByID(albumID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrAlbumNotFound // 保持错误类型的封装
		}
		return nil, fmt.Errorf("failed to find album by id: %w", err)
	}
	return albumInfo, nil
}
