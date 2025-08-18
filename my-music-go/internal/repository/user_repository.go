// internal/repository/user_repository.go
package repository

// TODO
// IUserRepository 定义了用户仓库需要实现的所有方法
// 这使得我们的 Service 层可以依赖这个接口，而不是具体的实现
import (
	"database/sql"
	"my-music-go/internal/models"

	"github.com/jmoiron/sqlx"
)

type IUserRepository interface {
	FindByUsername(username string) (*models.User, error)
	FindByID(userID int64) (*models.User, error)
	Create(user *models.User) error
	ListAllUsers() ([]models.User, error)
	FindUsernameAndNameByID(userID int64) (*models.UsernameResponse, error)
	FindProfileByID(userID int64) (*models.UserProfile, error)
	Update(user *models.User) error
}

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

func (r *UserRepository) FindByID(userID int64) (*models.User, error) {
	var user models.User
	query := "SELECT * FROM user WHERE id = ? LIMIT 1"

	// 3. 使用 db.Get 替代 QueryRow 和 Scan
	err := r.db.Get(&user, query, userID)
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
		INSERT INTO user (username, password, name, email, phone, avatar, bio, gender, birthday, location) 
		VALUES (:username, :password, :name, :email, :phone, :avatar, :bio, :gender, :birthday, :location)`

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

func (r *UserRepository) FindUsernameAndNameByID(userID int64) (*models.UsernameResponse, error) {
	var resp models.UsernameResponse
	query := `SELECT username, name FROM user WHERE id = ?`
	err := r.db.Get(&resp, query, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &resp, nil
}

func (r *UserRepository) FindProfileByID(userID int64) (*models.UserProfile, error) {
	var profile models.UserProfile
	query := `
		SELECT 
			u.*,
			(SELECT COUNT(*) FROM follow WHERE follower_id = u.id) as followings,
			(SELECT COUNT(*) FROM follow WHERE following_id = u.id) as followers
		FROM user u
		WHERE u.id = ?`

	err := r.db.Get(&profile, query, userID)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // 使用 nil, nil 来表示未找到
		}
		return nil, err
	}
	profile.Password = ""
	return &profile, nil
}

// Update 更新用户信息
func (r *UserRepository) Update(user *models.User) error {
	// 更新了部分字段, 其中 name 相当于 nickname
	query := `
		UPDATE user SET
			name = :name, avatar = :avatar, email = :email, phone = :phone, 
			gender = :gender, birthday = :birthday, location = :location, bio = :bio
		WHERE id = :id`

	_, err := r.db.NamedExec(query, user)
	return err
}
