package repository

import (
	"fmt"
	"my-music-go/internal/models"
	"strings"

	"github.com/jmoiron/sqlx"
)

type ILikeRepository interface {
	List(userID int64, params *models.ListLikedSongsRequestDTO) ([]models.SongDetailDTO, error)
	Count(userID int64, params *models.ListLikedSongsRequestDTO) (int, error)
	IsLiked(userID, songID int64) (bool, error)
	Create(userID, songID int64) error
	Delete(userID, songID int64) (int64, error)
}

type LikeRepository struct {
	db *sqlx.DB
}

func NewLikeRepository(db *sqlx.DB) *LikeRepository {
	return &LikeRepository{db: db}
}

// buildLikedSongsConditions 辅助函数，构建 WHERE 条件
func buildLikedSongsConditions(params *models.ListLikedSongsRequestDTO) (string, []interface{}) {
	var conditions []string
	var args []interface{}
	if params.Search != "" {
		// 注意这里的 JOIN 逻辑，搜索 artist.name
		conditions = append(conditions, "(s.name LIKE ? OR a.name LIKE ?)")
		likePattern := fmt.Sprintf("%%%s%%", params.Search)
		args = append(args, likePattern, likePattern)
	}
	if len(conditions) > 0 {
		return " AND " + strings.Join(conditions, " AND "), args
	}
	return "", args
}

func (r *LikeRepository) List(userID int64, params *models.ListLikedSongsRequestDTO) ([]models.SongDetailDTO, error) {
	var songs []models.SongDetailDTO
	baseQuery := `
		SELECT 
			s.*, 
			a.name as artist_name, a.id as artist_id,
			al.name as album_name, al.id as album_id, al.cover as album_cover
		FROM user_likes l
		JOIN song s ON l.song_id = s.id
		LEFT JOIN artist a ON s.author_id = a.id
		LEFT JOIN album al ON s.album_id = al.id
		WHERE l.user_id = ?`

	whereClause, args := buildLikedSongsConditions(params)

	finalArgs := append([]interface{}{userID}, args...)

	query := baseQuery + whereClause

	orderBy := "ORDER BY l.created_at DESC"

	allowedSorts := map[string]string{
		"latest":   "l.created_at DESC",
		"oldest":   "l.created_at ASC",
		"name":     "s.name ASC",
		"duration": "s.duration ASC",
	}

	if sort, ok := allowedSorts[params.SortBy]; ok {
		orderBy = "ORDER BY " + sort
	}

	query += " " + orderBy

	query += " LIMIT ? OFFSET ?"
	offset := (params.Page - 1) * params.PageSize
	finalArgs = append(finalArgs, params.PageSize, offset)

	err := r.db.Select(&songs, query, finalArgs...)
	return songs, err
}

func (r *LikeRepository) Count(userID int64, params *models.ListLikedSongsRequestDTO) (int, error) {
	var total int
	baseQuery := `
		SELECT COUNT(l.id) FROM user_likes l
		JOIN song s ON l.song_id = s.id
		LEFT JOIN artist a ON s.author_id = a.id
		WHERE l.user_id = ?`

	whereClause, args := buildLikedSongsConditions(params)
	finalArgs := append([]interface{}{userID}, args...)

	query := baseQuery + " " + whereClause

	err := r.db.Get(&total, query, finalArgs...)
	return total, err
}

func (r *LikeRepository) IsLiked(userID, songID int64) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM user_likes WHERE user_id = ? AND song_id = ?)"
	err := r.db.Get(&exists, query, userID, songID)
	return exists, err
}

func (r *LikeRepository) Create(userID, songID int64) error {
	query := "INSERT INTO user_likes (user_id, song_id) VALUES (?, ?)"
	_, err := r.db.Exec(query, userID, songID)
	return err
}

func (r *LikeRepository) Delete(userID, songID int64) (int64, error) {
	query := "DELETE FROM user_likes WHERE user_id = ? AND song_id = ?"
	res, err := r.db.Exec(query, userID, songID)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
