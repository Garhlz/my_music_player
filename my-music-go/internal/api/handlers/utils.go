package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetIDFromParam 从 URL 路径中解析 ID
func GetIDFromParam(c *gin.Context, paramName string) (int64, error) {
	idStr := c.Param(paramName)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID格式: " + paramName})
		return 0, err
	}
	return id, nil
}
