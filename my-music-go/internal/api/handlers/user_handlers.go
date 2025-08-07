// internal/api/handlers/user_handlers.go
package handlers

import (
	"errors"
	"log"
	"my-music-go/internal/models"
	"my-music-go/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserHandler 包含用户相关的处理器
type UserHandler struct {
	userService *services.UserService
}

// NewUserHandler 创建一个新的 UserHandler, 在此处进行依赖注入
func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// Register 处理用户注册的 HTTP 请求
// @Summary      用户注册
// @Description  根据提供的用户名和密码等信息，创建一个新的用户账户。
// @Tags         用户认证 (Auth)
// @Accept       json
// @Produce      json
// @Param        registerRequest body models.RegisterRequest true "用户注册信息"
// @Success      201 {object} map[string]string "{"message": "注册成功"}"
// @Failure      400 {object} map[string]string "{"error": "无效的请求参数"}"
// @Failure      409 {object} map[string]string "{"error": "该用户名已被注册"}"
// @Failure      500 {object} map[string]string "{"error": "注册失败，请稍后重试"}"
// @Router       /auth/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数"})
		return
	}

	err := h.userService.RegisterUser(&req)
	if err != nil {
		if errors.Is(err, services.ErrUserAlreadyExists) {
			c.JSON(http.StatusConflict, gin.H{"error": "该用户名已被注册"})
		} else {
			log.Printf("注册用户失败: %+v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "注册失败，请稍后重试"})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "注册成功"})
}

// Login 处理用户登录的 HTTP 请求
// @Summary      用户登录
// @Description  用户使用用户名和密码登录，成功后返回 JWT Token 和用户信息。
// @Tags         用户认证 (Auth)
// @Accept       json
// @Produce      json
// @Param        loginRequest body models.LoginRequest true "用户登录凭证"
// @Success      200 {object} map[string]interface{} "{"message": "登录成功", "token": "...", "user_id": 1}"
// @Failure      400 {object} map[string]string "{"error": "用户名和密码不能为空"}"
// @Failure      401 {object} map[string]string "{"error": "密码不正确"}"
// @Failure      404 {object} map[string]string "{"error": "用户不存在"}"
// @Failure      500 {object} map[string]string "{"error": "登录时发生未知错误"}"
// @Router       /auth/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名和密码不能为空"})
		return
	}

	userID, token, err := h.userService.LoginUser(&req)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrUserNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		case errors.Is(err, services.ErrInvalidCredentials):
			c.JSON(http.StatusUnauthorized, gin.H{"error": "密码不正确"})
		default:
			log.Printf("用户登录失败: %+v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "登录时发生未知错误"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "user_id": userID, "message": "登录成功"})
}

// GetUserProfile 处理获取用户完整公开信息的请求
// @Summary      获取用户个人资料
// @Description  根据用户ID获取用户的完整公开信息，包括关注数和粉丝数。
// @Tags         用户信息 (User)
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "用户 ID"
// @Success      200  {object}  models.UserProfile
// @Failure      400  {object}  map[string]string "{"error": "无效的用户ID格式"}"
// @Failure      401  {object}  map[string]string "{"error": "需要认证"}"
// @Failure      404  {object}  map[string]string "{"error": "用户未找到"}"
// @Failure      500  {object}  map[string]string "{"error": "获取用户信息失败，请稍后重试"}"
// @Router       /users/{id} [get]
func (h *UserHandler) GetUserProfile(c *gin.Context) {
	userID, err := GetIDFromParam(c, "id")
	if err != nil {
		return
	} // 已经在辅助函数内处理了返回值了

	profile, err := h.userService.GetUserProfile(userID)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "用户未找到"})
		} else {
			log.Printf("获取用户 Profile 失败: %+v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "获取用户信息失败，请稍后重试"})
		}
		return
	}
	c.JSON(http.StatusOK, profile)
}

// GetUsernameAndName 处理只获取用户名的请求
// @Summary      获取用户名和昵称
// @Description  根据用户ID获取用户的公开用户名和昵称。
// @Tags         用户信息 (User)
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "用户 ID"
// @Success      200  {object}  models.UsernameResponse
// @Failure      400  {object}  map[string]string "{"error": "无效的用户ID格式"}"
// @Failure      401  {object}  map[string]string "{"error": "需要认证"}"
// @Failure      404  {object}  map[string]string "{"error": "用户未找到"}"
// @Failure      500  {object}  map[string]string "{"error": "获取用户名失败，请稍后重试"}"
// @Router       /users/{id}/name [get]
func (h *UserHandler) GetUsernameAndName(c *gin.Context) {
	userID, err := GetIDFromParam(c, "id")
	if err != nil {
		return
	}

	usernameInfo, err := h.userService.GetUsernameAndName(userID)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "用户未找到"})
		} else {
			log.Printf("获取用户名失败: %+v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "获取用户名失败，请稍后重试"})
		}
		return
	}

	c.JSON(http.StatusOK, usernameInfo)
}

// UpdateMyProfile 处理更新当前用户信息的请求
// @Summary      更新当前登录用户信息
// @Description  更新当前认证用户（通过JWT Token识别）的个人资料。
// @Tags         用户信息 (User)
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        updateRequest body models.UpdateUserRequest true "要更新的用户信息"
// @Success      200  {object}  map[string]string "{"message": "用户信息更新成功"}"
// @Failure      400  {object}  map[string]string "{"error": "请求参数格式错误"}"
// @Failure      401  {object}  map[string]string "{"error": "需要认证"}"
// @Failure      404  {object}  map[string]string "{"error": "尝试更新的用户不存在"}"
// @Failure      500  {object}  map[string]string "{"error": "更新用户信息失败，请稍后重试"}"
// @Router       /users/{id} [put]
func (h *UserHandler) UpdateMyProfile(c *gin.Context) {
	// 从 JWT context 获取 userID
	authUserID := c.MustGet("userID").(int64)

	var req models.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数格式错误"})
		return
	}

	err := h.userService.UpdateUserProfile(authUserID, &req)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "尝试更新的用户不存在"})
		} else {
			log.Printf("更新用户信息失败: %+v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "更新用户信息失败，请稍后重试"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "用户信息更新成功"})
}
