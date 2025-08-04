package models

import "time"

// PlaylistInfo 对应数据库中的 `playlist_info` 表
type PlaylistInfo struct {
	ID          int64     `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	UserID      int64     `db:"user_id" json:"user_id"`
	Cover       *string   `db:"cover" json:"cover,omitempty"`
	Description *string   `db:"description" json:"description,omitempty"`
	IsPublic    bool      `db:"is_public" json:"is_public"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

// PlaylistSong 对应数据库中的 `playlist_songs` 关联表
type PlaylistSong struct {
	ID         int64     `db:"id" json:"id"`
	PlaylistID int64     `db:"playlist_id" json:"playlist_id"`
	SongID     int64     `db:"song_id" json:"song_id"`
	UserID     int64     `db:"user_id" json:"user_id"`
	AddedAt    time.Time `db:"added_at" json:"added_at"`
}