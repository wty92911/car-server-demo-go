package model

import (
	"time"
)

type QueueState int

const (
	Done QueueState = iota
	Wait
)

type StartProjectParams struct {
	UserId               string `json:"UserId" binding:"required"`
	UserIp               string `json:"UserIp"`
	ProjectId            string `json:"ProjectId" binding:"required"`
	ApplicationId        string `json:"ApplicationId"`
	ApplicationVersionId string `json:"ApplicationVersionId"`
	ClientSession        string `json:"ClientSession" binding:"required"`
	RequestId            string `json:"RequestId"`
	Sign                 string `json:"Sign"`
}

type StopProjectParams struct {
	UserId    string `json:"UserId" binding:"required"`
	RequestId string `json:"RequestId"`
	Sign      string `json:"Sign"`
}

type EnqueueParams struct {
	UserId               string `json:"UserId" binding:"required"`
	UserIp               string `json:"UserIp"`
	ProjectId            string `json:"ProjectId" binding:"required"`
	ApplicationId        string `json:"ApplicationId"`
	ApplicationVersionId string `json:"ApplicationVersionId"`
	RequestId            string `json:"RequestId"`
	Sign                 string `json:"Sign"`
	TimeStamp            time.Time
	State                QueueState
}

type DequeueParams struct {
	UserId    string `json:"UserId" binding:"required"`
	RequestId string `json:"RequestId"`
	Sign      string `json:"Sign"`
}

type Response struct {
	Code int `json:"Code"`
}

type EnqueueResponse struct {
	Index     int    `json:"Index"`
	UserId    string `json:"UserId"`
	ProjectId string `json:"ProjectId"`
}
