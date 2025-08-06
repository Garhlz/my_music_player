package repository

import (
	"database/sql"
	"fmt"
	"my-music-go/internal/models"
	"strings"

	"github.com/jmoiron/sqlx"
)

// ISongRepository 定义接口
type ISongRepository interface {
	FindDetailByID(songID int64) (*models.SongDetailDTO, error)
	List(params *models.ListSongsRequestDTO) ([]models.SongDetailDTO, error)
	Count(params *models.ListSongsRequestDTO) (int, error)
}

// SongRepository 实现
type SongRepository struct {
	db *sqlx.DB
}

// 返回的不是接口...工厂方法
func NewSongRepository(db *sqlx.DB) *SongRepository {
	return &SongRepository{db: db}
}

func (r *SongRepository) FindDetailByID(songID int64) (*models.SongDetailDTO, error) {
	var songDetail models.SongDetailDTO
	query := `
		SELECT 
			s.*, a.name as artist_name, a.id as artist_id,
			al.name as album_name, al.id as album_id, al.cover as album_cover
		FROM song s
		LEFT JOIN artist a ON s.author_id = a.id
		LEFT JOIN album al ON s.album_id = al.id
		WHERE s.id = ?`

	err := r.db.Get(&songDetail, query, songID)
	if err != nil {
		// 采用和 UserRepository 中完全一致的错误处理模式
		if err == sql.ErrNoRows {
			return nil, nil // 返回 nil, nil 表示“未找到”，交由 Service 层处理
		}
		return nil, err // 其他数据库错误，直接返回
	}
	return &songDetail, nil
}

// List 和 Count 方法是重点
func (r *SongRepository) List(params *models.ListSongsRequestDTO) ([]models.SongDetailDTO, error) {
	var songs []models.SongDetailDTO

	// 基础查询语句
	baseQuery := `
		SELECT s.*, a.name as artist_name, al.name as album_name
		FROM song s
		LEFT JOIN artist a ON s.author_id = a.id
		LEFT JOIN album al ON s.album_id = al.id`

	// 动态构建 WHERE 子句和参数
	var conditions []string
	var args []interface{}

	conditions = append(conditions, "s.is_public = TRUE")

	if params.Search != "" {
		conditions = append(conditions, "(s.name LIKE ? OR a.name LIKE ?)")
		likePattern := fmt.Sprintf("%%%s%%", params.Search)
		args = append(args, likePattern, likePattern)
	}

	// 组合最终的查询语句
	query := baseQuery
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	// 添加排序
	orderBy := "ORDER BY s.create_time DESC" // 默认排序
	allowedSorts := map[string]string{
		"latest":     "s.create_time DESC",
		"play_count": "s.play_count DESC",
		"like_count": "s.like_count DESC",
	}
	if sort, ok := allowedSorts[params.SortBy]; ok {
		orderBy = "ORDER BY " + sort
	}
	query += " " + orderBy

	// 添加分页
	query += " LIMIT ? OFFSET ?"
	offset := (params.Page - 1) * params.PageSize
	args = append(args, params.PageSize, offset)

	err := r.db.Select(&songs, query, args...)
	return songs, err
}

func (r *SongRepository) Count(params *models.ListSongsRequestDTO) (int, error) {
	var total int
	baseQuery := "SELECT COUNT(s.id) FROM song s LEFT JOIN artist a ON s.author_id = a.id"

	var conditions []string
	var args []interface{}

	conditions = append(conditions, "s.is_public = TRUE")

	if params.Search != "" {
		conditions = append(conditions, "(s.name LIKE ? OR a.name LIKE ?)")
		likePattern := fmt.Sprintf("%%%s%%", params.Search)
		args = append(args, likePattern, likePattern)
	}

	query := baseQuery
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	err := r.db.Get(&total, query, args...)
	return total, err
}
