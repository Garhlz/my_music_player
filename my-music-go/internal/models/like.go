package models

import "time"

// Like 对应数据库中的 `like` 表 (针对歌曲的点赞)
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

// UserLike 对应数据库中的 `user_likes` 表
// 注意：这个表的结构和功能与 `like` 表几乎完全一样，可能是冗余的。
// 在这里依然为你创建，方便你后续整理。
type UserLike struct {
	ID        int64     `db:"id" json:"id"`
	UserID    int64     `db:"user_id" json:"user_id"`
	SongID    int64     `db:"song_id" json:"song_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}