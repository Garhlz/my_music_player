// internal/services/user_service.go
package services

import (
	"errors"
	"my-music-go/internal/models"
	"my-music-go/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

// UserService 封装用户相关的业务逻辑
type UserService struct {
	userRepo *repository.UserRepository
}

// NewUserService 创建一个新的 UserService
func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// RegisterUser 处理用户注册的核心逻辑
func (s *UserService) RegisterUser(req *models.RegisterRequest) error {
	existingUser, err := s.userRepo.FindByUsername(req.Username)
	if err != nil {
		// 从 repository 返回的错误现在更明确了
		return errors.New("数据库查询出错: " + err.Error())
	}
	if existingUser != nil {
		return errors.New("用户名已存在")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("密码加密失败")
	}

	newUser := &models.User{
		Username: req.Username,
		Password: string(hashedPassword),
		Name:     req.Name,
		Email:    req.Email,
		Avatar:   "/assets/avatars/default-user.jpg", // 在 Service 层设置默认值
		Status:   1,
		Gender:   0,
	}

	err = s.userRepo.Create(newUser)
	if err != nil {
		return errors.New("注册失败: " + err.Error())
	}

	return nil
}


func (s *UserService) LoginUser(req *models.LoginRequest) (int64, string, error) {
	user, err := s.userRepo.FindByUsername(req.Username)
	if err != nil {
		return 0, "", errors.New("数据库出错: " + err.Error())
	}

	if user == nil {
		return 0, "", errors.New("用户未找到: " + err.Error())
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return 0, "", errors.New("密码错误: " + err.Error())
	}

	token := "i am a cute token" // todo jwt令牌生成, 调库

	return user.ID, token, nil

}
