package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Error 失败数据处理
func Error(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

func OK(c *gin.Context, data interface{}, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  msg,
		"data": data,
	})
}
