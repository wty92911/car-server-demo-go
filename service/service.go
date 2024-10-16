package service

import (
	car "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/car/v20220110"
	"github.com/wty92911/car-server-demo-go/model"
	"github.com/wty92911/car-server-demo-go/pkg"
	"log"
	"sync"
	"time"
)

const (
	enqueueTimeout     = 30 * time.Second
	queueCheckInterval = 1 * time.Second
)

var queue = pkg.NewBaseQueue(queueCheckInterval)
var waitQueueLock = sync.Mutex{}
var waitQueue = make(map[string]*model.EnqueueParams)
var carClient = pkg.NewCarClient()

func ApplyConcurrent(params *model.StartProjectParams) (*car.ApplyConcurrentResponse, error) {
	req := car.NewApplyConcurrentRequest()
	req.UserId = &params.UserId
	req.ProjectId = &params.ProjectId
	req.ApplicationId = &params.ApplicationId
	req.ApplicationVersionId = &params.ApplicationVersionId
	req.UserIp = &params.UserIp
	return carClient.ApplyConcurrent(req)

}

func CreateSession(params *model.StartProjectParams) (*car.CreateSessionResponse, error) {
	req := car.NewCreateSessionRequest()
	req.UserId = &params.UserId
	req.UserIp = &params.UserIp
	req.ClientSession = &params.ClientSession
	return carClient.CreateSession(req)
}

func DestroySession(userId string) (*car.DestroySessionResponse, error) {
	req := car.NewDestroySessionRequest()
	req.UserId = &userId
	return carClient.DestroySession(req)
}

func Enqueue(params *model.EnqueueParams) (*model.EnqueueResponse, error) {
	log.Println("Enqueue")
	waitQueueLock.Lock()
	if _, ok := waitQueue[params.UserId]; ok {
		waitQueueLock.Unlock()
		return &model.EnqueueResponse{
			Index:     queue.IndexOf(params.UserId),
			UserId:    params.UserId,
			ProjectId: params.ProjectId,
		}, nil
	}
	waitQueueLock.Unlock()
	log.Println("Enqueue1")
	// 尝试申请并发
	applyParams := &model.StartProjectParams{
		UserId:               params.UserId,
		ProjectId:            params.ProjectId,
		ApplicationId:        params.ApplicationId,
		ApplicationVersionId: params.ApplicationVersionId,
		UserIp:               params.UserIp,
	}
	_, err := ApplyConcurrent(applyParams)
	if err == nil {
		return &model.EnqueueResponse{
			Index:     -1,
			UserId:    params.UserId,
			ProjectId: params.ProjectId,
		}, nil
	}
	log.Println("Enqueue2")
	// 申请失败，加入队列
	waitQueueLock.Lock()
	params.TimeStamp = time.Now()
	waitQueue[params.UserId] = params
	waitQueueLock.Unlock()
	log.Println("Enqueue3")
	queue.Enqueue(params.UserId, shouldDequeue)
	rsp := &model.EnqueueResponse{
		Index:     queue.IndexOf(params.UserId),
		UserId:    params.UserId,
		ProjectId: params.ProjectId,
	}
	log.Println("Enqueue4")
	return rsp, nil
}

// shouldDequeue 判断是否可以删除队列元素
func shouldDequeue(userId string) bool {
	waitQueueLock.Lock()
	defer waitQueueLock.Unlock()
	item := waitQueue[userId]
	// 如果超过 30s 写在队列里，删除队列元素
	if time.Since(item.TimeStamp) > enqueueTimeout {
		delete(waitQueue, userId)
		return true
	}

	// 再次尝试申请并发
	params := &model.StartProjectParams{
		UserId:               item.UserId,
		ProjectId:            item.ProjectId,
		ApplicationId:        item.ApplicationId,
		ApplicationVersionId: item.ApplicationVersionId,
		UserIp:               item.UserIp,
	}
	_, err := ApplyConcurrent(params)
	if err == nil {
		item.State = model.Done
		delete(waitQueue, userId)
		return true
	}
	item.State = model.Wait
	return false
}
func Dequeue(userId string) {
	queue.Dequeue(userId)
	waitQueueLock.Lock()
	defer waitQueueLock.Unlock()
	delete(waitQueue, userId)
}
