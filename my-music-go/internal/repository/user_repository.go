// internal/repository/user_repository.go
package repository

import (
	"database/sql"
	"my-music-go/internal/models"

	"github.com/jmoiron/sqlx"
)

// UserRepository 提供了访问用户数据的方法
type UserRepository struct {
	// 1. 持有一个 *sqlx.DB 实例，而不是依赖全局变量
	db *sqlx.DB
}

// NewUserRepository 创建一个新的 UserRepository 实例
// 2. 构造函数接收一个 *sqlx.DB，实现依赖注入
func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

// FindByUsername 通过用户名查找用户
func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	query := "SELECT * FROM user WHERE username = ? LIMIT 1"

	// 3. 使用 db.Get 替代 QueryRow 和 Scan
	err := r.db.Get(&user, query, username)
	if err != nil {
		// 如果错误是 sql.ErrNoRows，说明用户不存在，这不是一个程序错误
		if err == sql.ErrNoRows {
			return nil, nil
		}
		// 其他错误（如数据库断开）则是一个需要处理的程序错误
		return nil, err
	}
	return &user, nil
}

// Create 创建一个新用户
func (r *UserRepository) Create(user *models.User) error {
	// 4. 使用命名查询 (Named Query) 来执行插入，更清晰、更安全
	query := `
		INSERT INTO user (username, password, name, email, avatar, status, gender) 
		VALUES (:username, :password, :name, :email, :avatar, :status, :gender)`

	// 5. 使用 db.NamedExec，它会自动匹配 user 结构体的字段和查询中的命名参数
	_, err := r.db.NamedExec(query, user)
	return err
}

// ListAllUsers 获取所有用户的列表
func (r *UserRepository) ListAllUsers() ([]models.User, error) {
	var users []models.User
	query := "SELECT * FROM user ORDER BY created_time DESC"

	err := r.db.Select(&users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}
