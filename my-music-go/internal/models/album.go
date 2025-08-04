package models

import "time"

// Album 对应数据库中的 `album` 表
type Album struct {
	ID          int64      `db:"id" json:"id"`
	Name        string     `db:"name" json:"name"`
	Cover       string     `db:"cover" json:"cover,omitempty"`
	ReleaseDate *time.Time `db:"release_date" json:"release_date,omitempty"`
	Description *string    `db:"description" json:"description,omitempty"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
}

// ArtistAlbum 对应数据库中的 `artist_album` 关联表 (M2M)
type ArtistAlbum struct {
	ArtistID int64 `db:"artist_id" json:"artist_id"`
	AlbumID  int64 `db:"album_id" json:"album_id"`
}
