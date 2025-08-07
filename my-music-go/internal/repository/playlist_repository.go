package repository

import (
	"database/sql"
	"my-music-go/internal/models"

	"github.com/jmoiron/sqlx"
)

type IPlaylistRepository interface {
	// 歌单管理
	FindInfoByID(playlistID int64) (*models.PlaylistInfoDTO, error)
	ListByUserID(userID int64, params *models.ListPlaylistsRequestDTO) ([]models.PlaylistInfoDTO, error)
	CountByUserID(userID int64) (int, error)
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
		if err == sql.ErrNoRows {
			return nil, nil // 错误依然是留给上层进行处理
		}
		return nil, err
	}
	return &info, nil
}

// ListByUserID 查找当前用户创建的所有播放列表
func (r *PlaylistRepository) ListByUserID(userID int64, params *models.ListPlaylistsRequestDTO) ([]models.PlaylistInfoDTO, error) {
	var playlists []models.PlaylistInfoDTO
	query := `
		SELECT 
			pi.*, u.username as creator_name,
			(SELECT COUNT(*) FROM playlist_songs ps WHERE ps.playlist_id = pi.id) as song_count
		FROM playlist_info pi
		LEFT JOIN user u ON pi.user_id = u.id
		WHERE pi.user_id = ?
		ORDER BY pi.created_at DESC LIMIT ? OFFSET ?`

	offset := (params.Page - 1) * params.PageSize
	err := r.db.Select(&playlists, query, userID, params.PageSize, offset)
	return playlists, err
}

// CountByUserID 查找用户id创建的播放列表数量
func (r *PlaylistRepository) CountByUserID(userID int64) (int, error) {
	var total int
	query := "SELECT COUNT(*) FROM playlist_info WHERE user_id = ?"
	err := r.db.Get(&total, query, userID)
	return total, err
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
