package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"my-music-go/internal/models"
	"strings"

	"github.com/jmoiron/sqlx"
)

type IPlaylistRepository interface {
	// 歌单管理
	FindInfoByID(playlistID int64) (*models.PlaylistInfoDTO, error)
	ListByUserID(userID int64, params *models.ListPlaylistsRequestDTO, publicOnly bool) ([]models.PlaylistInfoDTO, error)
	CountByUserID(userID int64, params *models.ListPlaylistsRequestDTO, publicOnly bool) (int, error)
	Create(playlist *models.PlaylistInfo) (int64, error)
	Update(playlist *models.PlaylistInfo) error
	Delete(playlistID int64) error

	// 歌单中的歌曲管理
	AddSong(playlistID, songID, userID int64) error
	RemoveSong(playlistID, songID int64) (int64, error)
	IsSongInPlaylist(playlistID, songID int64) (bool, error)
}

type PlaylistRepository struct {
	db *sqlx.DB
}

func NewPlaylistRepository(db *sqlx.DB) *PlaylistRepository {
	return &PlaylistRepository{db: db}
}

// FindInfoByID 查找当前playlistID对应的列表详细信息
func (r *PlaylistRepository) FindInfoByID(playlistID int64) (*models.PlaylistInfoDTO, error) {
	var info models.PlaylistInfoDTO
	query := `
		SELECT 
			pi.*, u.username as creator_name,
			(SELECT COUNT(*) FROM playlist_songs ps WHERE ps.playlist_id = pi.id) as song_count
		FROM playlist_info pi
		LEFT JOIN user u ON pi.user_id = u.id
		WHERE pi.id = ?`

	err := r.db.Get(&info, query, playlistID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // 错误依然是留给上层进行处理
		}
		return nil, err
	}
	return &info, nil
}

func buildPlaylistConditions(userID int64, params *models.ListPlaylistsRequestDTO, publicOnly bool) (string, []interface{}) {
	var conditions []string
	var args []interface{}

	// 1. 用户ID条件
	conditions = append(conditions, "pi.user_id = ?")
	args = append(args, userID)

	// 2. 搜索条件
	if params.Search != "" {
		conditions = append(conditions, "pi.name LIKE ?") // 只有一个占位符
		likePattern := "%" + params.Search + "%"
		args = append(args, likePattern) // 只添加一个参数
	}

	// 3. 是否仅公开
	if publicOnly {
		conditions = append(conditions, "is_public = 1")
	}
	// 组合所有条件
	if len(conditions) > 0 {
		return "WHERE " + strings.Join(conditions, " AND "), args
	}

	return "", args
}

func (r *PlaylistRepository) ListByUserID(userID int64, params *models.ListPlaylistsRequestDTO, publicOnly bool) ([]models.PlaylistInfoDTO, error) {
	var playlists []models.PlaylistInfoDTO

	// 1. 定义基础查询语句，不包含 WHERE 及之后的子句
	baseQuery := `
		SELECT 
			pi.*, u.username as creator_name,
			(SELECT COUNT(*) FROM playlist_songs ps WHERE ps.playlist_id = pi.id) as song_count
		FROM playlist_info pi
		LEFT JOIN user u ON pi.user_id = u.id`

	// 2. 使用辅助函数构建 WHERE 子句和参数
	whereClause, args := buildPlaylistConditions(userID, params, publicOnly)

	// 3. 构建完整的 SQL 语句
	// 注意每个子句前面的空格
	query := baseQuery + " " + whereClause
	query += " ORDER BY pi.created_at DESC" // 这里可以根据 params.SortBy 进一步做动态排序
	query += " LIMIT ? OFFSET ?"

	// 4. 将分页参数添加到 args 列表的末尾
	offset := (params.Page - 1) * params.PageSize
	finalArgs := append(args, params.PageSize, offset)

	// 5. 执行查询，传入构建好的 query 和 finalArgs
	err := r.db.Select(&playlists, query, finalArgs...)
	if err != nil {
		return nil, fmt.Errorf("failed to list playlists by user id: %w", err)
	}

	return playlists, nil
}

// 修改函数签名，让它能接收 params
func (r *PlaylistRepository) CountByUserID(userID int64, params *models.ListPlaylistsRequestDTO, publicOnly bool) (int, error) {
	var total int

	// 1. 基础查询语句
	baseQuery := "SELECT COUNT(pi.id) FROM playlist_info pi"

	// 2. 使用同一个辅助函数构建 WHERE 子句和参数
	whereClause, args := buildPlaylistConditions(userID, params, publicOnly)

	// 3. 构建完整的 SQL 语句
	query := baseQuery + " " + whereClause

	// 4. 执行查询
	err := r.db.Get(&total, query, args...)
	if err != nil {
		return 0, fmt.Errorf("failed to count playlists by user id: %w", err)
	}

	return total, nil
}

// Create 创建新的播放列表
func (r *PlaylistRepository) Create(playlist *models.PlaylistInfo) (int64, error) {
	query := `
		INSERT INTO playlist_info (name, description, user_id, is_public, cover) 
		VALUES (?, ?, ?, ?, ?)`
	res, err := r.db.Exec(query, playlist.Name, playlist.Description, playlist.UserID, playlist.IsPublic, playlist.Cover)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// Update 更新播放列表信息
func (r *PlaylistRepository) Update(playlist *models.PlaylistInfo) error {
	query := `
		UPDATE playlist_info SET 
			name = :name, description = :description, is_public = :is_public, 
			cover = :cover, updated_at = :updated_at
		WHERE id = :id AND user_id = :user_id`
	_, err := r.db.NamedExec(query, playlist)
	return err
}

// Delete 开启事务, 先删除关联表中的所有歌曲记录, 在删除歌单本身
// todo 这里还需要深究一下, 不是很熟悉
func (r *PlaylistRepository) Delete(playlistID int64) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback() // 保证出错时回滚

	// 先删除关联表中的歌曲
	if _, err := tx.Exec("DELETE FROM playlist_songs WHERE playlist_id = ?", playlistID); err != nil {
		return err
	}
	// 再删除歌单本身
	if _, err := tx.Exec("DELETE FROM playlist_info WHERE id = ?", playlistID); err != nil {
		return err
	}

	return tx.Commit() // 提交事务
}

func (r *PlaylistRepository) AddSong(playlistID, songID, userID int64) error {
	// user_id 字段用于记录是谁添加的，这在你的原表设计中存在
	query := "INSERT INTO playlist_songs (playlist_id, song_id, user_id) VALUES (?, ?, ?)"
	_, err := r.db.Exec(query, playlistID, songID, userID)
	return err
}

func (r *PlaylistRepository) RemoveSong(playlistID, songID int64) (int64, error) {
	query := "DELETE FROM playlist_songs WHERE playlist_id = ? AND song_id = ?"
	res, err := r.db.Exec(query, playlistID, songID)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (r *PlaylistRepository) IsSongInPlaylist(playlistID, songID int64) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM playlist_songs WHERE playlist_id = ? AND song_id = ?)"
	err := r.db.Get(&exists, query, playlistID, songID)
	return exists, err
}
