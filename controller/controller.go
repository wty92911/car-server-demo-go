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

	params.UserIp = c.ClientIP()
	_, err := service.ApplyConcurrent(&params)
	if err != nil {
		Err(c, ErrCodeConcurrencyFailure, err)
		return
	}

	_, err = service.CreateSession(&params)
	if err != nil {
		Err(c, ErrCodeCreateSessionFailed, err)
		return
	}
	Ok(c, params.RequestId, nil)
}

func StopProject(c *gin.Context) {
	var params model.StopProjectParams
	if err := c.ShouldBindJSON(&params); err != nil {
		Err(c, ErrCodeInvalidParams, err)
		return
	}

	_, err := service.DestroySession(params.UserId)
	if err != nil {
		Err(c, ErrCodeReleaseSessionFailed, err)
		return
	}
	Ok(c, params.RequestId, nil)
}

func Enqueue(c *gin.Context) {
	var params model.EnqueueParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	params.UserIp = c.ClientIP()
	rsp, err := service.Enqueue(&params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, rsp)
		return
	}
	Ok(c, params.RequestId, rsp)
}

func Dequeue(c *gin.Context) {
	var params model.DequeueParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service.Dequeue(params.UserId)
	Ok(c, params.RequestId, nil)
}
