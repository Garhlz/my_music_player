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

// ---- DTOs for Playlist APIs ----

// ListPlaylistsRequestDTO 用于绑定获取歌单列表的查询参数
type ListPlaylistsRequestDTO struct {
	Page     int `form:"page"`
	PageSize int `form:"pageSize"`
}

// PlaylistInfoDTO 用于歌单列表的响应，包含了歌曲数量和创建者名字
type PlaylistInfoDTO struct {
	PlaylistInfo
	CreatorName string `db:"creator_name" json:"creator_name"`
	SongCount   int    `db:"song_count" json:"song_count"`
}

// PlaylistDetailDTO 用于获取歌单详情的完整响应
type PlaylistDetailDTO struct {
	PlaylistInfoDTO
	Songs *PaginatedResponseDTO `json:"songs"` // 复用分页结构
}

// CreatePlaylistRequestDTO 用于绑定创建歌单的请求体
type CreatePlaylistRequestDTO struct {
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description"`
	IsPublic    *bool   `json:"is_public"`
}

// UpdatePlaylistRequestDTO 用于绑定更新歌单的请求体
type UpdatePlaylistRequestDTO struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	IsPublic    *bool   `json:"is_public"`
	Cover       *string `json:"cover"`
}

// AddSongToPlaylistRequestDTO 用于绑定添加歌曲到歌单的请求体
type AddSongToPlaylistRequestDTO struct {
	SongID int64 `json:"song_id" binding:"required"`
}
