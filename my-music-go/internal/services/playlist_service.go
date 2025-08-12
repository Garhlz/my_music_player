package services

import (
	"database/sql"
	"errors"
	"fmt"
	"my-music-go/internal/models"
	"my-music-go/internal/repository"
	"sync"
	"time"
)

// service的结构体中, 一般都放的是接口而不是直接的类
type PlaylistService struct {
	playlistRepo repository.IPlaylistRepository
	songService  *SongService
}

func NewPlaylistService(playlistRepo repository.IPlaylistRepository, songService *SongService) *PlaylistService {
	return &PlaylistService{playlistRepo: playlistRepo, songService: songService}
}

// checkPlaylistOwner 是一个可复用的权限检查辅助函数
func (s *PlaylistService) checkPlaylistOwner(userID, playlistID int64) (*models.PlaylistInfoDTO, error) {
	playlist, err := s.playlistRepo.FindInfoByID(playlistID)
	if err != nil {
		return nil, fmt.Errorf("database error on find playlist: %w", err)
	}
	if playlist == nil {
		return nil, ErrPlaylistNotFound
	}
	// 没有权限就不可以处理(用户id != 歌单创建者id)
	if playlist.UserID != userID {
		return nil, ErrForbidden
	}
	return playlist, nil
}

// ListUserPlaylists 获取指定用户的歌单列表, 又带有常规的分页查询内容
// 又有并行处理查询分页歌单和数量的逻辑
func (s *PlaylistService) ListUserPlaylists(userID int64, params *models.ListPlaylistsRequestDTO) (*models.PaginatedResponseDTO, error) {
	if params.Page <= 0 {
		params.Page = 1
	}
	if params.PageSize <= 0 {
		params.PageSize = 10
	}

	var wg sync.WaitGroup
	var playlists []models.PlaylistInfoDTO
	var total int
	var errPlaylists, errTotal error

	wg.Add(2)
	go func() {
		defer wg.Done()
		playlists, errPlaylists = s.playlistRepo.ListByUserID(userID, params)
	}()
	go func() {
		defer wg.Done()
		total, errTotal = s.playlistRepo.CountByUserID(userID, params)
	}()
	wg.Wait()

	if errPlaylists != nil {
		return nil, fmt.Errorf("failed to list playlists: %w", errPlaylists)
	}
	if errTotal != nil {
		return nil, fmt.Errorf("failed to count playlists: %w", errTotal)
	}

	return &models.PaginatedResponseDTO{Total: total, List: playlists}, nil
}

// CreatePlaylist 为指定用户创建歌单
func (s *PlaylistService) CreatePlaylist(userID int64, req *models.CreatePlaylistRequestDTO) (int64, error) {
	defaultCover := "/assets/covers/playlists/default.jpg"
	// todo 修改封面逻辑, 因为此处连文件上传逻辑都没有
	playlist := &models.PlaylistInfo{
		Name:        req.Name,
		Description: req.Description,
		UserID:      userID,
		IsPublic:    *req.IsPublic,
		Cover:       &defaultCover, // 默认封面
	}
	return s.playlistRepo.Create(playlist)
}

// UpdatePlaylist 更新歌单信息
func (s *PlaylistService) UpdatePlaylist(userID, playlistID int64, req *models.UpdatePlaylistRequestDTO) error {
	playlist, err := s.checkPlaylistOwner(userID, playlistID)
	if err != nil {
		return err
	}

	if req.Name != nil {
		playlist.Name = *req.Name
	}
	if req.Description != nil {
		playlist.Description = req.Description
	}
	if req.IsPublic != nil {
		playlist.IsPublic = *req.IsPublic
	}
	if req.Cover != nil {
		playlist.Cover = req.Cover
	}
	playlist.UpdatedAt = time.Now()

	return s.playlistRepo.Update(&playlist.PlaylistInfo)
}

// DeletePlaylist 删除歌单
func (s *PlaylistService) DeletePlaylist(userID, playlistID int64) error {
	if _, err := s.checkPlaylistOwner(userID, playlistID); err != nil {
		return err
	}
	return s.playlistRepo.Delete(playlistID)
}

// AddSongToPlaylist 添加歌曲到歌单
func (s *PlaylistService) AddSongToPlaylist(userID, playlistID, songID int64) error {
	if _, err := s.checkPlaylistOwner(userID, playlistID); err != nil {
		return err
	}

	var wg sync.WaitGroup
	var song *models.SongDetailDTO
	var songExists, alreadyInPlaylist bool
	var errSong, errCheck error
	wg.Add(2)
	go func() {
		defer wg.Done()
		song, errSong = s.songService.GetSongDetail(songID)
		songExists = song != nil
	}()
	go func() {
		defer wg.Done()
		alreadyInPlaylist, errCheck = s.playlistRepo.IsSongInPlaylist(playlistID, songID)
	}()
	wg.Wait()

	if !songExists {
		return ErrSongNotFound
	}
	if errSong != nil {
		return fmt.Errorf("db error finding song details: %w", errCheck)
	}
	if errCheck != nil {
		return fmt.Errorf("db error checking if song in playlist: %w", errCheck)
	}
	if alreadyInPlaylist {
		return ErrSongAlreadyInPlaylist
	}

	return s.playlistRepo.AddSong(playlistID, songID, userID)
}

// RemoveSongFromPlaylist 从歌单移除歌曲
func (s *PlaylistService) RemoveSongFromPlaylist(userID, playlistID, songID int64) error {
	if _, err := s.checkPlaylistOwner(userID, playlistID); err != nil {
		return err
	}
	rowsAffected, err := s.playlistRepo.RemoveSong(playlistID, songID)
	if err != nil {
		return fmt.Errorf("db error on remove song: %w", err)
	}
	if rowsAffected == 0 {
		return ErrSongNotInPlaylist
	}
	return nil
}

// 新建一个简单的get方法
func (s *PlaylistService) GetPlaylist(playlistID int64) (*models.PlaylistInfoDTO, error) {
	playlistInfo, err := s.playlistRepo.FindInfoByID(playlistID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrPlaylistNotFound
		}
		return nil, fmt.Errorf("failed to find playlist by id: %w", err)
	}
	return playlistInfo, nil
}
