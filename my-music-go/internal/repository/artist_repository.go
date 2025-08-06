package repository

import (
	"database/sql"
	"errors"
	"my-music-go/internal/models"

	"github.com/jmoiron/sqlx"
)

type IArtistRepository interface {
	FindByID(artistID int64) (*models.Artist, error)
}

type ArtistRepository struct {
	db *sqlx.DB
}

func NewArtistRepository(db *sqlx.DB) *ArtistRepository {
	return &ArtistRepository{db: db}
}

func (r *ArtistRepository) FindByID(artistID int64) (*models.Artist, error) {
	var artist models.Artist
	query := "SELECT * FROM artist WHERE id = ?"

	err := r.db.Get(&artist, query, artistID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // 未找到
		}
		return nil, err // 数据库错误
	}
	return &artist, nil
}
