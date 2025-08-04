package models

import "time"

// Song 对应数据库中的 `song` 表
type Song struct {
	ID         int64     `db:"id" json:"id"`
	Name       string    `db:"name" json:"name"`
	AuthorID   int64     `db:"author_id" json:"author_id"`
	AlbumID    int64     `db:"album_id" json:"album_id"`
	URL        string    `db:"url" json:"url"`
	Cover      string    `db:"cover" json:"cover"`
	Duration   int       `db:"duration" json:"duration"`
	PlayCount  int       `db:"play_count" json:"play_count"`
	LikeCount  int       `db:"like_count" json:"like_count"`
	IsPublic   bool      `db:"is_public" json:"is_public"`
	CreateTime time.Time `db:"create_time" json:"create_time"`
}