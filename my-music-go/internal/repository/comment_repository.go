package repository

import (
	"database/sql"
	"errors"
	"my-music-go/internal/models"

	"github.com/jmoiron/sqlx"
)

type ICommentRepository interface {
	FindDTOByID(commentID int64, currentUserID *int64) (*models.CommentDTO, error)
	ListBySongID(songID int64, currentUserID *int64, params *models.ListCommentsRequestDTO) ([]models.CommentDTO, error)
	CountBySongID(songID int64) (int, error)
	ListByParentID(parentID int64, currentUserID *int64) ([]models.CommentDTO, error)
	Create(comment *models.Comment) (int64, error)
	Update(comment *models.Comment) error
	Delete(commentID int64) error
	ToggleLike(userID, commentID int64) (isLiked bool, err error)
	GetLikeCount(commentID int64) (int, error)
}

type CommentRepository struct{ db *sqlx.DB }

func NewCommentRepository(db *sqlx.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

// FindDTOByID 获取单个评论的 DTO
func (r *CommentRepository) FindDTOByID(commentID int64, currentUserID *int64) (*models.CommentDTO, error) {
	var comment models.CommentDTO
	query := `
		SELECT 
			c.*, u.username, u.name, u.avatar, ru.username as reply_to_username, ru.name as reply_to_name,
			EXISTS(SELECT 1 FROM comment_like cl WHERE cl.comment_id = c.id AND cl.user_id = ?) as is_liked_by_me
		FROM comment c
		LEFT JOIN user u ON c.user_id = u.id
		LEFT JOIN user ru ON c.reply_to_user_id = ru.id
		WHERE c.id = ?`

	err := r.db.Get(&comment, query, currentUserID, commentID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &comment, nil
}

// ListBySongID (WHERE 条件修改)
func (r *CommentRepository) ListBySongID(songID int64, currentUserID *int64, params *models.ListCommentsRequestDTO) ([]models.CommentDTO, error) {
	var comments []models.CommentDTO
	// 【修改】: 使用 is_root = TRUE 来查找根评论
	query := `
		SELECT 
			c.*, u.username, u.name, u.avatar, ru.username as reply_to_username, ru.name as reply_to_name,
			EXISTS(SELECT 1 FROM comment_like cl WHERE cl.comment_id = c.id AND cl.user_id = ?) as is_liked_by_me
		FROM comment c
		LEFT JOIN user u ON c.user_id = u.id
		LEFT JOIN user ru ON c.reply_to_user_id = ru.id
		WHERE c.song_id = ? AND c.is_root = TRUE
		ORDER BY c.like_count DESC, c.created_at DESC
		LIMIT ? OFFSET ?`

	offset := (params.Page - 1) * params.PageSize
	err := r.db.Select(&comments, query, currentUserID, songID, params.PageSize, offset)
	return comments, err
}

// CountBySongID 歌曲的根评论数量
func (r *CommentRepository) CountBySongID(songID int64) (int, error) {
	var total int
	query := "SELECT COUNT(*) FROM comment WHERE song_id = ? AND is_root = TRUE"
	err := r.db.Get(&total, query, songID)
	return total, err
}

// ListByParentID 获取评论的回复
func (r *CommentRepository) ListByParentID(parentID int64, currentUserID *int64) ([]models.CommentDTO, error) {
	var comments []models.CommentDTO
	query := `
		SELECT 
			c.*, u.username, u.name, u.avatar, ru.username as reply_to_username, ru.name as reply_to_name,
			EXISTS(SELECT 1 FROM comment_like cl WHERE cl.comment_id = c.id AND cl.user_id = ?) as is_liked_by_me
		FROM comment c
		LEFT JOIN user u ON c.user_id = u.id
		LEFT JOIN user ru ON c.reply_to_user_id = ru.id
		WHERE c.parent_id = ?
		ORDER BY c.created_at ASC`
	err := r.db.Select(&comments, query, currentUserID, parentID)
	return comments, err
}

// Create
func (r *CommentRepository) Create(comment *models.Comment) (int64, error) {
	// 在 INSERT 语句中加入 is_root 字段
	query := `
		INSERT INTO comment (user_id, song_id, text, parent_id, reply_to_user_id, is_root) 
		VALUES (:user_id, :song_id, :text, :parent_id, :reply_to_user_id, :is_root)`
	res, err := r.db.NamedExec(query, comment)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// Update ...
func (r *CommentRepository) Update(comment *models.Comment) error {
	query := "UPDATE comment SET text = :text, updated_at = :updated_at WHERE id = :id AND user_id = :user_id"
	_, err := r.db.NamedExec(query, comment)
	return err
}

// Delete 删除评论（事务）
func (r *CommentRepository) Delete(commentID int64) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// 删除此评论的所有点赞记录 (依赖外键的 ON DELETE CASCADE 也可以，但显式删除更清晰)
	if _, err := tx.Exec("DELETE FROM comment_like WHERE comment_id = ?", commentID); err != nil {
		return err
	}
	// 删除评论本身 (其子评论会因为外键的 ON DELETE SET NULL 而 parent_id 变为 NULL)
	if _, err := tx.Exec("DELETE FROM comment WHERE id = ?", commentID); err != nil {
		return err
	}

	return tx.Commit()
}

// ToggleLike 点赞/取消点赞 (事务)
func (r *CommentRepository) ToggleLike(userID, commentID int64) (isLiked bool, err error) {
	// 数据库事务
	// 以下多个操作必须要么全部成功，要么全部失败，以保证数据一致性，因此必须在事务中执行。
	tx, err := r.db.Beginx()
	if err != nil {
		return false, err
	}
	defer tx.Rollback() // 保证出错时事务会自动回滚

	var exists bool
	// 确定当前喜欢状态
	queryCheck := "SELECT EXISTS(SELECT 1 FROM comment_like WHERE user_id = ? AND comment_id = ?)"
	if err = tx.Get(&exists, queryCheck, userID, commentID); err != nil {
		return false, err
	}

	if exists {
		// 如果喜欢, 删除喜欢记录
		if _, err = tx.Exec("DELETE FROM comment_like WHERE user_id = ? AND comment_id = ?", userID, commentID); err != nil {
			return false, err
		}
		// 喜欢数量也要更新
		if _, err = tx.Exec("UPDATE comment SET like_count = like_count - 1 WHERE id = ? AND like_count > 0", commentID); err != nil {
			return false, err
		}
		isLiked = false
	} else {
		if _, err = tx.Exec("INSERT INTO comment_like (user_id, comment_id) VALUES (?, ?)", userID, commentID); err != nil {
			return false, err
		}
		if _, err = tx.Exec("UPDATE comment SET like_count = like_count + 1 WHERE id = ?", commentID); err != nil {
			return false, err
		}
		isLiked = true
	}

	return isLiked, tx.Commit() // 所有操作成功，提交事务
}

func (r *CommentRepository) GetLikeCount(commentID int64) (int, error) {
	var count int
	query := "SELECT like_count FROM comment WHERE id = ?"
	err := r.db.Get(&count, query, commentID)
	return count, err
}
