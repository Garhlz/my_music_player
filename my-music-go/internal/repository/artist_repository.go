package repository

import (
	"my-music-go/internal/models"

	"github.com/jmoiron/sqlx"
)

type IArtistRepository interface {
	FindByID(artistID int64) (*models.Artist, error)
	GetAlbumCount(artistID int64) (*models.AlbumCount, error)
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
		return nil, err
	}
	return &artist, nil
}

func (r *ArtistRepository) GetAlbumCount(artistID int64) (*models.AlbumCount, error) {
	var albumCount models.AlbumCount
	query := `
		SELECT COUNT(*) as album_count
		FROM artist_album aa
		LEFT JOIN artist ar ON ar.id = aa.artist_id
		WHERE ar.id = ?
	`
	err := r.db.Get(&albumCount, query, artistID)

	if err != nil {
		return nil, err
	}

	return &albumCount, nil
}
