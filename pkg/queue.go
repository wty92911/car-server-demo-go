package pkg

import (
	"errors"
	"log"
	"sync"
	"time"
)

type CallbackFunc func(string) bool

type BaseQueue struct {
	checkInterval time.Duration
	checkTimer    *time.Timer
	mu            sync.Mutex
	queue         []string
	indexes       map[string]int
	callbacks     map[string]CallbackFunc
}

func NewBaseQueue(checkInterval time.Duration) *BaseQueue {
	bq := &BaseQueue{
		checkInterval: checkInterval,
		queue:         make([]string, 0),
		indexes:       make(map[string]int),
		callbacks:     make(map[string]CallbackFunc),
	}
	bq.checkTimer = time.AfterFunc(bq.checkInterval, bq.checkQueue)
	return bq
}

func (bq *BaseQueue) CanDequeue(key string) (bool, error) {
	bq.mu.Lock()
	defer bq.mu.Unlock()

	cb, exists := bq.callbacks[key]
	if !exists {
		return true, nil
	}

	if cb == nil {
		return false, errors.New("callback is not callable")
	}

	return cb(key), nil
}

func (bq *BaseQueue) checkQueue() {
	if bq.Count() > 0 {
		item := bq.First()
		if ok, err := bq.CanDequeue(item); err == nil && ok {
			bq.Dequeue(item)
		} else {
			log.Fatal(err)
		}
	}
}

func (bq *BaseQueue) Dequeue(key string) {
	index := bq.IndexOf(key)
	if index == -1 {
		return
	}
	bq.queue = append(bq.queue[:index], bq.queue[index+1:]...)
	bq.RemoveCallback(key)
	bq.reloadIndexes()
}

func (bq *BaseQueue) IndexOf(key string) int {
	bq.mu.Lock()
	defer bq.mu.Unlock()

	index, exists := bq.indexes[key]
	if !exists {
		return -1
	}
	return index
}

func (bq *BaseQueue) reloadIndexes() {
	bq.mu.Lock()
	defer bq.mu.Unlock()

	bq.indexes = make(map[string]int)
	for index, item := range bq.queue {
		bq.indexes[item] = index
	}
}

func (bq *BaseQueue) Enqueue(key string, checkQueueDoneCallback CallbackFunc) {
	index := bq.IndexOf(key)
	bq.AddCallback(key, checkQueueDoneCallback)
	if index >= 0 {
		return
	}
	bq.mu.Lock()
	defer bq.mu.Unlock()
	bq.queue = append(bq.queue, key)
	bq.indexes[key] = len(bq.queue) - 1
}

func (bq *BaseQueue) GetKeyByIndex(index int) string {
	bq.mu.Lock()
	defer bq.mu.Unlock()

	if index < len(bq.queue) {
		return bq.queue[index]
	}
	return ""
}

func (bq *BaseQueue) AddCallback(key string, cb CallbackFunc) {
	bq.mu.Lock()
	defer bq.mu.Unlock()
	bq.callbacks[key] = cb
}

func (bq *BaseQueue) RemoveCallback(key string) {
	bq.mu.Lock()
	defer bq.mu.Unlock()
	delete(bq.callbacks, key)
}

func (bq *BaseQueue) GetCallback(key string) (CallbackFunc, bool) {
	bq.mu.Lock()
	defer bq.mu.Unlock()
	cb, exists := bq.callbacks[key]
	return cb, exists
}

func (bq *BaseQueue) Count() int {
	bq.mu.Lock()
	defer bq.mu.Unlock()
	return len(bq.queue)
}

func (bq *BaseQueue) First() string {
	bq.mu.Lock()
	defer bq.mu.Unlock()
	if len(bq.queue) == 0 {
		return ""
	}
	return bq.queue[0]
}
