// internal/api/handlers/user_handlers.go
package handlers

import (
	"my-music-go/internal/models"
	"my-music-go/internal/services" // 引入 service
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
func (h *UserHandler) Register(c *gin.Context) {
	var req models.RegisterRequest
	// 1. 绑定并验证请求体
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名和密码不能为空"})
		return
	}

	// 2. 调用 Service 层处理业务
	err := h.userService.RegisterUser(&req)
	if err != nil {
		// 根据 Service 返回的错误类型，给予不同的 HTTP 响应
		if err.Error() == "用户名已存在" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		// 其他错误都视为服务器内部错误
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 3. 返回成功响应
	c.JSON(http.StatusCreated, gin.H{"message": "注册成功"})
}

// 暂时迅速照猫画虎完成了一个接口
func (h *UserHandler) Login(c *gin.Context) {
	var req models.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名和密码不能为空"})
		return
	}

	userID, token, err := h.userService.LoginUser(&req)
	if err != nil {
		// todo 不同的相关错误处理, 至少有三种不同的错误状态, 似乎在外部定义为枚举更好
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"token": token, "user_id": userID, "message": "登录成功"})
	return
}
