package model

import (
	"time"
)

type StartProjectParams struct {
	UserId               string `json:"UserId" binding:"required"`
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
	ProjectId            string `json:"ProjectId" binding:"required"`
	ApplicationId        string `json:"ApplicationId"`
	ApplicationVersionId string `json:"ApplicationVersionId"`
	RequestId            string `json:"RequestId"`
	Sign                 string `json:"Sign"`
}

type DequeueParams struct {
	UserId    string `json:"UserId" binding:"required"`
	RequestId string `json:"RequestId"`
	Sign      string `json:"Sign"`
}

type QueueItem struct {
	UserId               string
	ProjectId            string
	ApplicationId        string
	ApplicationVersionId string
	UserIp               string
	TimeStamp            time.Time
	State                string
}

type Response struct {
	Code int `json:"Code"`
}
