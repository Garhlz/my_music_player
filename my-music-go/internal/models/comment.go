package models

import (
	"time"
)

type Comment struct {
	ID            int64     `db:"id" json:"id"`
	UserID        int64     `db:"user_id" json:"user_id"`
	SongID        int64     `db:"song_id" json:"song_id"`
	Text          string    `db:"text" json:"text"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	ParentID      *int64    `db:"parent_id" json:"parent_id,omitempty"`               // 使用指针类型处理 NULL
	ReplyToUserID *int64    `db:"reply_to_user_id" json:"reply_to_user_id,omitempty"` // 使用指针
	LikeCount     int       `db:"like_count" json:"like_count"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
	IsRoot        bool      `db:"is_root" json:"is_root"` // 新增字段
}

// ---- DTOs for Comment APIs ----

// CommentDTO 用于API响应，包含了用户信息和状态
type CommentDTO struct {
	Comment
	Username        string       `db:"username" json:"username"`
	Name            string       `db:"name" json:"name"` // 新增， 用于展示
	Avatar          string       `db:"avatar" json:"avatar"`
	ReplyToUsername *string      `db:"reply_to_username" json:"reply_to_username,omitempty"`
	ReplyToName     *string      `db:"reply_to_name" json:"reply_to_name,omitempty"`
	IsLikedByMe     bool         `db:"is_liked_by_me" json:"is_liked_by_me"`
	Replies         []CommentDTO `json:"replies,omitempty"` // 用于嵌套回复
}

// ListCommentsRequestDTO 用于绑定获取评论列表的查询参数
type ListCommentsRequestDTO struct {
	Page     int `form:"page" default:"1"`
	PageSize int `form:"pageSize" default:"10"`
}

// CreateCommentRequestDTO 用于绑定发表评论的请求体
type CreateCommentRequestDTO struct {
	Text          string `json:"text" binding:"required,max=1000"`
	ParentID      *int64 `json:"parent_id"`
	ReplyToUserID *int64 `json:"reply_to_user_id"`
}

// UpdateCommentRequestDTO 用于绑定更新评论的请求体
type UpdateCommentRequestDTO struct {
	Text string `json:"text" binding:"required,max=1000"`
}

// LikeCommentResponseDTO 定义点赞评论的响应
type LikeCommentResponseDTO struct {
	LikeCount   int  `json:"like_count"`
	IsLikedByMe bool `json:"is_liked_by_me"`
}

// LikeCommentResponseDTO 定义获取评论赞数的响应
type CommentLikeCountResponseDTO struct {
	LikeCount int `json:"like_count"`
}
