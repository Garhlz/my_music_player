package models

import "time"

// Artist 对应数据库中的 `artist` 表
type Artist struct {
	ID          int64     `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Sex         *string   `db:"sex" json:"sex,omitempty"`
	Avatar      *string   `db:"avatar" json:"avatar,omitempty"`
	Description *string   `db:"description" json:"description,omitempty"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}
