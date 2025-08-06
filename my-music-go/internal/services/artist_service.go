package services

import (
	"database/sql"
	"errors"
	"fmt"
	"my-music-go/internal/models"
	"my-music-go/internal/repository"
	"sync"
)

// ArtistService 依赖 ArtistRepository 和 SongService
type ArtistService struct {
	artistRepo  repository.IArtistRepository
	songService *SongService // 依赖另一个 Service
}

// NewArtistService 构造函数
func NewArtistService(artistRepo repository.IArtistRepository, songService *SongService) *ArtistService {
	return &ArtistService{
		artistRepo:  artistRepo,
		songService: songService,
	}
}

// GetArtistDetail 是核心业务方法
func (s *ArtistService) GetArtistDetail(artistID int64, songParams *models.ListSongsRequestDTO) (*models.ArtistDetailResponseDTO, error) {
	var wg sync.WaitGroup

	// 准备变量来接收并发任务的结果
	var artistInfo *models.Artist
	var paginatedSongs *models.PaginatedResponseDTO
	var errArtist, errSongs error

	wg.Add(2) // 启动两个并发任务

	// 任务1: 获取艺术家信息
	go func() {
		defer wg.Done()
		artistInfo, errArtist = s.artistRepo.FindByID(artistID)
	}()

	// 任务2: 获取该艺术家的歌曲列表
	go func() {
		defer wg.Done()
		// 将 artistID 设置到歌曲查询参数中，以复用 ListSongs
		songParams.ArtistID = &artistID
		paginatedSongs, errSongs = s.songService.ListSongs(songParams)
	}()

	wg.Wait() // 等待两个任务全部完成

	// 统一检查错误
	if errArtist != nil {
		// 如果是未找到的特定错误，我们单独处理
		if errors.Is(errArtist, sql.ErrNoRows) {
			return nil, ErrArtistNotFound // 假设你在 errors.go 中定义了这个
		}
		return nil, fmt.Errorf("failed to get artist info: %w", errArtist)
	}
	if errSongs != nil {
		return nil, fmt.Errorf("failed to get artist's songs: %w", errSongs)
	}

	// 组合最终的响应DTO
	response := &models.ArtistDetailResponseDTO{
		ArtistInfo: artistInfo,
		Songs:      paginatedSongs,
	}

	return response, nil
}
