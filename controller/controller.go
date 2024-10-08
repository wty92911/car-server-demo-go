package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wty92911/car-server-demo-go/model"
	"github.com/wty92911/car-server-demo-go/service"
	"net/http"
)

type ErrCode int

const (
	// Success represents a successful request.
	Success ErrCode = 0

	// ErrCodeSignValidationError represents a sign validation error.
	ErrCodeSignValidationError ErrCode = 10000

	// ErrCodeInvalidParams represents a missing necessary parameter error.
	ErrCodeInvalidParams ErrCode = 10001

	// ErrCodeQueueInProcess represents a queue in process error.
	ErrCodeQueueInProcess ErrCode = 10100

	// ErrCodeQueueCompleted represents a queue completed status.
	ErrCodeQueueCompleted ErrCode = 10101

	// ErrCodeCreateSessionFailed represents a failure to create a cloud rendering session.
	ErrCodeCreateSessionFailed ErrCode = 10200

	// ErrCodeReleaseSessionFailed represents a failure to release a cloud rendering session.
	ErrCodeReleaseSessionFailed ErrCode = 10201

	// ErrCodeConcurrencyFailure represents a concurrency application failure.
	ErrCodeConcurrencyFailure ErrCode = 10202

	// ErrCodeInternalError represents a generic internal error.
	ErrCodeInternalError ErrCode = 10002
)

func StartProject(c *gin.Context) {
	var params model.StartProjectParams
	if err := c.ShouldBindJSON(&params); err != nil {
		Err(c, ErrCodeInvalidParams, err)
		return
	}

	userIp := c.ClientIP()
	_, err := service.ApplyConcurrent(&params, userIp)
	if err != nil {
		Err(c, ErrCodeConcurrencyFailure, err)
		return
	}

	rsp, err := service.CreateSession(&params, userIp)
	if err != nil {
		Err(c, ErrCodeCreateSessionFailed, err)
		return
	}
	Ok(c, *rsp.Response.RequestId, nil)
}

func StopProject(c *gin.Context) {
	var params model.StopProjectParams
	if err := c.ShouldBindJSON(&params); err != nil {
		Err(c, ErrCodeInvalidParams, err)
		return
	}

	ret, err := service.DestroySession(params.UserId)
	if err != nil {
		Err(c, ErrCodeReleaseSessionFailed, err)
		return
	}
	Ok(c, *ret.Response.RequestId, nil)
}

func Enqueue(c *gin.Context) {
	var params model.EnqueueParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIp := c.ClientIP()
	ret, err := service.Enqueue(params, userIp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ret)
		return
	}

	c.JSON(http.StatusOK, ret)
}

func Dequeue(c *gin.Context) {
	var params model.DequeueParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service.Dequeue(params.UserId)
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
