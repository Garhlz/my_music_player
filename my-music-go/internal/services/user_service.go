// internal/services/user_service.go
package services

import (
	"fmt"
	"my-music-go/internal/config"
	"my-music-go/internal/models"
	"my-music-go/internal/repository"
	"my-music-go/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

// UserService 封装用户相关的业务逻辑
type UserService struct {
	userRepo repository.IUserRepository
	cfg      config.Config
}

// NewUserService 创建一个新的 UserService
func NewUserService(userRepo repository.IUserRepository, cfg config.Config) *UserService {
	return &UserService{userRepo: userRepo, cfg: cfg}
}

// RegisterUser 处理用户注册的核心逻辑
func (s *UserService) RegisterUser(req *models.RegisterRequest) error {
	existingUser, err := s.userRepo.FindByUsername(req.Username)
	if err != nil {
		// 从 repository 返回的错误现在更明确了
		// 使用 %w 包装原始错误，这样上层可以用 errors.Is/As 来解开它
		return fmt.Errorf("database query error: %w", err)
	}
	if existingUser != nil {
		return ErrUserAlreadyExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// 已经实现了不必要的字段可以不存在
	newUser := &models.User{
		Username: req.Username,
		Password: string(hashedPassword),
	}
	if req.Name != nil {
		newUser.Name = req.Name
	}
	if req.Email != nil {
		newUser.Email = req.Email
	}
	if req.Phone != nil {
		newUser.Phone = req.Phone
	}
	if req.Avatar != nil {
		newUser.Avatar = *req.Avatar
	} else {
		// 这里其实没有必要手动设置默认头像， 在前端处理即可
		newUser.Avatar = "/assets/avatars/default-user.jpg"
	}
	if req.Bio != nil {
		newUser.Bio = req.Bio
	}
	if req.Gender != nil {
		newUser.Gender = *req.Gender
	} else {
		newUser.Gender = 0 // 0是保密, 1是男, 2是女
	}
	if req.Birthday != nil {
		newUser.Birthday = req.Birthday
	}
	if req.Location != nil {
		newUser.Location = req.Location
	}

	err = s.userRepo.Create(newUser)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (s *UserService) LoginUser(req *models.LoginRequest) (int64, string, error) {
	user, err := s.userRepo.FindByUsername(req.Username)
	if err != nil {
		// 这是数据库层面的错误
		return 0, "", fmt.Errorf("database error on find user: %w", err)
	}

	if user == nil {
		return 0, "", ErrUserNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return 0, "", ErrInvalidCredentials
	}

	jwtSecret := s.cfg.JWTSecret
	token, err := utils.GenerateToken(user.ID, jwtSecret)
	if err != nil {
		return 0, "", fmt.Errorf("failed to generate token: %w", err)
	}

	// 没有错误, 就是正确的
	return user.ID, token, nil

}

func (s *UserService) GetUserProfile(userID int64) (*models.UserProfile, error) {
	profile, err := s.userRepo.FindProfileByID(userID)
	if err != nil {
		return nil, fmt.Errorf("database error on find profile: %w", err)
	}
	if profile == nil {
		return nil, ErrUserNotFound
	}
	return profile, nil
}

// GetUsernameAndName 获取用户名和昵称
func (s *UserService) GetUsernameAndName(userID int64) (*models.UsernameResponse, error) {
	resp, err := s.userRepo.FindUsernameAndNameByID(userID)
	if err != nil {
		return nil, fmt.Errorf("database error on find username: %w", err)
	}
	if resp == nil {
		return nil, ErrUserNotFound
	}
	return resp, nil
}

func (s *UserService) UpdateUserProfile(userID int64, req *models.UpdateUserRequest) error {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return fmt.Errorf("database error on find user for update: %w", err)
	}
	if user == nil {
		return ErrUserNotFound
	}

	// 使用 DTO 中的数据更新 user 模型
	if req.Name != nil {
		user.Name = req.Name
	}
	if req.Avatar != nil {
		user.Avatar = *req.Avatar
	}
	if req.Email != nil {
		user.Email = req.Email
	}
	if req.Phone != nil {
		user.Phone = req.Phone
	}
	if req.Gender != nil {
		user.Gender = *req.Gender
	}
	if req.Birthday != nil {
		user.Birthday = req.Birthday
	}
	if req.Location != nil {
		user.Location = req.Location
	}
	if req.Bio != nil {
		user.Bio = req.Bio
	}

	// 调用 repository 的 Update 方法
	return s.userRepo.Update(user)
}

func (s *UserService) GetFollowStatus(currentUserID, targetUserID int64) (*models.FollowStatusResponse, error) {
	// 确认两个用户存在
	resp0, err := s.userRepo.FindUsernameAndNameByID(currentUserID)
	if err != nil {
		return nil, fmt.Errorf("database error on find username: %w", err)
	}
	if resp0 == nil {
		return nil, ErrUserNotFound
	}

	resp1, err := s.userRepo.FindUsernameAndNameByID(targetUserID)
	if err != nil {
		return nil, fmt.Errorf("database error on find username: %w", err)
	}
	if resp1 == nil {
		return nil, ErrUserNotFound
	}

	resp, err := s.userRepo.GetFollowStatus(currentUserID, targetUserID)
	if err != nil {
		return nil, fmt.Errorf("database error on find username: %w", err)
	}
	if resp == nil {
		return nil, ErrUserNotFound
	}

	return resp, err
}

func (s *UserService) ToggleFollow(currentUserID, targetUserID int64) error {
	// 确认两个用户存在
	resp0, err := s.userRepo.FindUsernameAndNameByID(currentUserID)
	if err != nil {
		return fmt.Errorf("database error on find username: %w", err)
	}
	if resp0 == nil {
		return ErrUserNotFound
	}

	resp1, err := s.userRepo.FindUsernameAndNameByID(targetUserID)
	if err != nil {
		return fmt.Errorf("database error on find username: %w", err)
	}
	if resp1 == nil {
		return ErrUserNotFound
	}

	err1 := s.userRepo.ToggleFollow(currentUserID, targetUserID)
	if err1 != nil {
		return fmt.Errorf("database error on find username: %w", err1)
	}

	return nil
}
