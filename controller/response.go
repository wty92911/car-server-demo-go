package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Ok(c *gin.Context, requestId string, session *string) {
	c.JSON(http.StatusOK, gin.H{
		"Code":      Success,
		"Msg":       "ok",
		"RequestId": requestId,
		"SessionDescribe": map[string]*string{
			"ServerSession": session,
		},
	})
}

func Err(c *gin.Context, code ErrCode, err error) {
	log.Printf("err %v code %v", err, code)
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
