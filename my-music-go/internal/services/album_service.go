package services

import (
	"fmt"
	"my-music-go/internal/models"
	"my-music-go/internal/repository"
	"sync"
)

// SongService 是否依赖? 是的, 需要知道歌曲数量
type AlbumService struct {
	albumRepo   repository.IAlbumRepository // 注意这里是接口, 而且不是指针
	songService *SongService                // 依赖另一个 Service
}

// NewAlbumService 构造函数
func NewAlbumService(albumRepo repository.IAlbumRepository, songService *SongService) *AlbumService {
	return &AlbumService{
		albumRepo:   albumRepo,
		songService: songService,
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

func (s *AlbumService) GetAlbumDetail(albumID int64, songParams *models.ListSongsRequestDTO) (*models.AlbumDetailDTO, error) {
	var wg sync.WaitGroup

	var album *models.Album
	var artists []models.Artist
	var paginatedSongs *models.PaginatedResponseDTO
	var errAlbum, errArtists, errSongs error

	wg.Add(3)

	go func() {
		defer wg.Done()
		album, errAlbum = s.albumRepo.FindByID(albumID)
	}()

	go func() {
		defer wg.Done()
		artists, errArtists = s.albumRepo.FindArtistsByAlbumID(albumID)
	}()

	// 复用 songService, 获取专辑的歌曲列表
	go func() {
		defer wg.Done()
		songParams.AlbumID = &albumID
		paginatedSongs, errSongs = s.songService.ListSongs(songParams)
	}()

	wg.Wait()

	// 统一检查错误
	if errAlbum != nil || errArtists != nil || errSongs != nil {
		// 这里可以根据具体错误返回，或者统一返回一个通用错误
		return nil, fmt.Errorf("failed to get album detail. albumErr:%w, artistsErr:%w, songsErr:%w", errAlbum, errArtists, errSongs)
	}

	if album == nil {
		return nil, ErrAlbumNotFound // 使用正确的哨兵错误
	}
	// 组合最终的响应DTO
	response := &models.AlbumDetailDTO{
		Album:   *album, // 注意这里解引用
		Artists: artists,
		Songs:   paginatedSongs,
	}

	return response, nil
}
