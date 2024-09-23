package service

import (
	car "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/car/v20220110"
	"github.com/wty92911/car-server-demo-go/model"
	"github.com/wty92911/car-server-demo-go/pkg"
	"time"
)

var carClient = pkg.NewCarClient()

func ApplyConcurrent(params *model.StartProjectParams, userIp string) (*car.ApplyConcurrentResponse, error) {
	req := car.NewApplyConcurrentRequest()
	req.ProjectId = &params.ProjectId
	req.ApplicationId = &params.ApplicationId
	req.ApplicationVersionId = &params.ApplicationVersionId
	req.UserIp = &userIp
	return carClient.ApplyConcurrent(req)

}

func CreateSession(params *model.StartProjectParams, userIp string) (*car.CreateSessionResponse, error) {
	req := car.NewCreateSessionRequest()
	req.UserId = &params.UserId
	req.UserIp = &userIp
	req.ClientSession = &params.ClientSession
	return carClient.CreateSession(req)
}

func DestroySession(userId string) (*car.DestroySessionResponse, error) {
	req := car.NewDestroySessionRequest()
	req.UserId = &userId
	return carClient.DestroySession(req)
}

var waitQueue = make(map[string]model.QueueItem)

func Enqueue(params model.EnqueueParams, userIp string) (model.Response, error) {
	if _, exists := waitQueue[params.UserId]; exists {
		// 更新现有用户的时间戳和其他信息
		waitQueue[params.UserId] = model.QueueItem{
			UserId:               params.UserId,
			ProjectId:            params.ProjectId,
			ApplicationId:        params.ApplicationId,
			ApplicationVersionId: params.ApplicationVersionId,
			UserIp:               userIp,
			TimeStamp:            time.Now(),
			State:                "Wait",
		}
	} else {
		// 新增用户到队列
		waitQueue[params.UserId] = model.QueueItem{
			UserId:               params.UserId,
			ProjectId:            params.ProjectId,
			ApplicationId:        params.ApplicationId,
			ApplicationVersionId: params.ApplicationVersionId,
			UserIp:               userIp,
			TimeStamp:            time.Now(),
			State:                "Wait",
		}
	}
	return model.Response{Code: 0}, nil
}

func Dequeue(userId string) {
	delete(waitQueue, userId)
}
