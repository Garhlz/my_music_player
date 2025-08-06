package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"my-music-go/internal/models"
	"strings"

	"github.com/jmoiron/sqlx"
)

type IAlbumRepository interface {
	List(params *models.ListAlbumsRequestDTO) ([]models.AlbumListDTO, error)
	Count(params *models.ListAlbumsRequestDTO) (int, error)
	FindByID(albumID int64) (*models.Album, error)
	FindArtistsByAlbumID(albumID int64) ([]models.Artist, error)
}

type AlbumRepository struct {
	db *sqlx.DB
}

func NewAlbumRepository(db *sqlx.DB) *AlbumRepository {
	return &AlbumRepository{db: db}
}

// buildAlbumListConditions 是一个辅助函数，用于构建 List 和 Count 的 WHERE 条件
func buildAlbumListConditions(params *models.ListAlbumsRequestDTO) (string, []interface{}) {
	var conditions []string
	var args []interface{}
	if params.Search != "" {
		conditions = append(conditions, "(al.name LIKE ? OR ar.name LIKE ?)")
		likePattern := fmt.Sprintf("%%%s%%", params.Search)
		args = append(args, likePattern, likePattern)
	}

	if len(conditions) == 0 {
		return "", nil
	}
	return "WHERE " + strings.Join(conditions, " AND "), args
}

func (r *AlbumRepository) List(params *models.ListAlbumsRequestDTO) ([]models.AlbumListDTO, error) {
	var albums []models.AlbumListDTO

	baseQuery := `
		SELECT 
			al.*, 
			GROUP_CONCAT(DISTINCT ar.name SEPARATOR ', ') AS artist_names,
			COUNT(DISTINCT s.id) AS song_count
		FROM album al
		LEFT JOIN artist_album aa ON al.id = aa.album_id
		LEFT JOIN artist ar ON aa.artist_id = ar.id
		LEFT JOIN song s ON al.id = s.album_id`

	whereClause, args := buildAlbumListConditions(params)
	query := baseQuery + " " + whereClause

	// GROUP BY 是聚合查询的关键，必须要有
	query += " GROUP BY al.id"

	// 排序逻辑优化
	orderBy := "ORDER BY al.release_date DESC, al.id DESC" // 默认按发行日期
	allowedSorts := map[string]string{
		"latest":       "al.created_at DESC",
		"release_date": "al.release_date DESC",
	}
	if sort, ok := allowedSorts[params.SortBy]; ok {
		orderBy = "ORDER BY " + sort
	}
	query += " " + orderBy

	// 分页
	query += " LIMIT ? OFFSET ?"
	offset := (params.Page - 1) * params.PageSize
	args = append(args, params.PageSize, offset)

	err := r.db.Select(&albums, query, args...)
	return albums, err
}

func (r *AlbumRepository) Count(params *models.ListAlbumsRequestDTO) (int, error) {
	var total int
	baseQuery := `
		SELECT COUNT(DISTINCT al.id) 
		FROM album al
		LEFT JOIN artist_album aa ON al.id = aa.album_id
		LEFT JOIN artist ar ON aa.artist_id = ar.id`

	whereClause, args := buildAlbumListConditions(params)
	query := baseQuery + " " + whereClause

	err := r.db.Get(&total, query, args...)
	return total, err
}

// FindByID 获取专辑信息
func (r *AlbumRepository) FindByID(albumID int64) (*models.Album, error) {
	var album models.Album
	query := "SELECT * FROM album WHERE id = ?"
	err := r.db.Get(&album, query, albumID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &album, nil
}

// FindArtistsByAlbumID 获取专辑的所有艺术家
func (r *AlbumRepository) FindArtistsByAlbumID(albumID int64) ([]models.Artist, error) {
	var artists []models.Artist
	query := `
		SELECT ar.* FROM artist ar
		JOIN artist_album aa ON ar.id = aa.artist_id
		WHERE aa.album_id = ?`
	err := r.db.Select(&artists, query, albumID)
	return artists, err
}
