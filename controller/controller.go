package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	secretID  = "secretID"
	secretKey = "secretKey"
)

func SetSecret(id, key string) {
	secretID = id
	secretKey = key
}

// StartProject 启动应用
func StartProject(c *gin.Context) {
	// 获取并处理请求参数
	// ...

	// 返回响应
	c.JSON(http.StatusOK, gin.H{
		"Code":      0,
		"Msg":       "Success",
		"RequestId": "req123",
		"SessionDescribe": gin.H{
			"ServerSession": "session123",
			"RequestId":     "apiReq123",
		},
	})
}

// StopProject 结束应用
func StopProject(c *gin.Context) {
	// 获取并处理请求参数
	// ...

	// 返回响应
	c.JSON(http.StatusOK, gin.H{
		"Code":      0,
		"Msg":       "Success",
		"RequestId": "req123",
	})
}

// Enqueue 用户加入队列
func Enqueue(c *gin.Context) {
	// 获取并处理请求参数
	// ...

	// 返回响应
	c.JSON(http.StatusOK, gin.H{
		"Code":      0,
		"Msg":       "Success",
		"RequestId": "req123",
		"Data": gin.H{
			"Index":     1,
			"UserId":    "user123",
			"ProjectId": "cap-123",
		},
	})
}

// Dequeue 用户退出队列
func Dequeue(c *gin.Context) {
	// 获取并处理请求参数
	// ...

	// 返回响应
	c.JSON(http.StatusOK, gin.H{
		"Code":      0,
		"Msg":       "Success",
		"RequestId": "req123",
	})
}
