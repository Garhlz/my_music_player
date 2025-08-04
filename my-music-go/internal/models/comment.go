package models

import (
	"database/sql"
	"time"
)

// Comment 对应数据库中的 `comment` 表
type Comment struct {
	ID            int64         `db:"id" json:"id"`
	UserID        int64         `db:"user_id" json:"user_id"`
	SongID        int64         `db:"song_id" json:"song_id"`
	Text          string        `db:"text" json:"text"`
	CreatedAt     time.Time     `db:"created_at" json:"created_at"`
	ParentID      sql.NullInt64 `db:"parent_id" json:"parent_id,omitempty"` // 使用 sql.NullInt64 处理可能为 NULL 的外键
	ReplyToUserID sql.NullInt64 `db:"reply_to_user_id" json:"reply_to_user_id,omitempty"`
	LikeCount     int           `db:"like_count" json:"like_count"`
	UpdatedAt     time.Time     `db:"updated_at" json:"updated_at"`
}