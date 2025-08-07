package models

import "time"

// Like 对应数据库中的 `user_likes` 表 (针对歌曲的点赞)
// 原本叫like表, 但是和保留字冲突了, 现在迁移到user_likes表
type Like struct {
	ID        int64     `db:"id" json:"id"`
	UserID    int64     `db:"user_id" json:"user_id"`
	SongID    int64     `db:"song_id" json:"song_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

// CommentLike 对应数据库中的 `comment_like` 表
type CommentLike struct {
	ID        int64     `db:"id" json:"id"`
	CommentID int64     `db:"comment_id" json:"comment_id"`
	UserID    int64     `db:"user_id" json:"user_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type ListLikedSongsRequestDTO struct {
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
	Search   string `form:"search"`
	SortBy   string `form:"sortBy"`
}
