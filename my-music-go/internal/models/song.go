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

// ---- DTOs for Song APIs ----

// SongDetailDTO 用于获取单曲详情接口的返回，包含了关联的艺术家和专辑信息
type SongDetailDTO struct {
	Song               // 结构体嵌入, 因为没有字段名, 相当于直接展开, 体现了组合的思想
	ArtistName *string `db:"artist_name" json:"artist_name"` // 使用指针以防 JOIN 结果为 NULL
	ArtistID   *int64  `db:"artist_id" json:"artist_id"`
	AlbumName  *string `db:"album_name" json:"album_name"`
	AlbumID    *int64  `db:"album_id" json:"album_id"`
	AlbumCover *string `db:"album_cover" json:"album_cover"`
}

// ListSongsRequestDTO 用于绑定获取歌曲列表接口的查询参数
// `form` 标签用于 Gin 绑定GET请求的路径查询参数, 而json标签用于绑定POST请求的JSON请求体
type ListSongsRequestDTO struct {
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
	Search   string `form:"search"`
	SortBy   string `form:"sortBy"`

	ArtistID *int64 `form:"artistId"`
	AlbumID  *int64 `form:"albumId"`
}

// PaginatedResponseDTO 是一个通用的分页响应结构体
type PaginatedResponseDTO struct {
	Total int         `json:"total"`
	List  interface{} `json:"list"` // 使用 interface{} 来适应任何类型的列表数据
}
