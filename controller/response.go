package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ok(c *gin.Context, requestId string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"Code":      Success,
		"Msg":       "ok",
		"RequestId": requestId,
		"Data":      data,
	})
}

func Err(c *gin.Context, code ErrCode, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"Code": code,
		"Msg":  err.Error(),
	})
}

func Rsp(c *gin.Context, code ErrCode, requestId string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"Code":      code,
		"Msg":       "ok",
		"RequestId": requestId,
		"Data":      data,
	})
}
