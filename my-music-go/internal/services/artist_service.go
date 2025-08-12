package services

import (
	"database/sql"
	"errors"
	"fmt"
	"my-music-go/internal/models"
	"my-music-go/internal/repository"
)

type ArtistService struct {
	artistRepo repository.IArtistRepository
}

func NewArtistService(artistRepo repository.IArtistRepository) *ArtistService {
	return &ArtistService{
		artistRepo: artistRepo,
	}
}

// GetArtist 是新的、职责单一的方法
func (s *ArtistService) GetArtist(artistID int64) (*models.Artist, error) {
	artistInfo, err := s.artistRepo.FindByID(artistID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrArtistNotFound // 保持错误类型的封装
		}
		return nil, fmt.Errorf("failed to find artist by id: %w", err)
	}
	return artistInfo, nil
}
