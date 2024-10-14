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
	time.AfterFunc(bq.checkInterval, bq.checkQueue)
	return bq
}

func (bq *BaseQueue) CanDequeue(key string) (bool, error) {
	log.Printf("checking can dequeue: %s\n", key)
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
	log.Println("checking queue")
	bq.mu.Lock()
	for i := range bq.queue {
		log.Printf("queue: %s\n", bq.queue[i])
	}
	bq.mu.Unlock()

	if bq.Count() > 0 {
		item := bq.First()
		if ok, err := bq.CanDequeue(item); err == nil && ok {
			bq.Dequeue(item)
		}
	}
	time.AfterFunc(bq.checkInterval, bq.checkQueue)
}

func (bq *BaseQueue) Dequeue(key string) {
	log.Printf("dequeueing: %s\n\n", key)
	index := bq.IndexOf(key)
	if index == -1 {
		return
	}
	bq.mu.Lock()
	bq.queue = append(bq.queue[:index], bq.queue[index+1:]...)
	bq.mu.Unlock()
	bq.RemoveCallback(key)
	bq.reloadIndexes()
}

func (bq *BaseQueue) IndexOf(key string) int {
	log.Printf("checking index of %s\n", key)
	bq.mu.Lock()
	defer bq.mu.Unlock()

	index, exists := bq.indexes[key]
	if !exists {
		return -1
	}
	return index
}

func (bq *BaseQueue) reloadIndexes() {
	log.Printf("reloading indexes...\n")
	bq.mu.Lock()
	defer bq.mu.Unlock()

	bq.indexes = make(map[string]int)
	for index, item := range bq.queue {
		bq.indexes[item] = index
	}
}

func (bq *BaseQueue) Enqueue(key string, checkQueueDoneCallback CallbackFunc) {
	log.Printf("enqueing: %s\n\n", key)
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

func (bq *BaseQueue) AddCallback(key string, cb CallbackFunc) {
	log.Printf("adding callback: %s\n\n", key)
	bq.mu.Lock()
	defer bq.mu.Unlock()
	bq.callbacks[key] = cb
}

func (bq *BaseQueue) RemoveCallback(key string) {
	log.Printf("removing callback: %s\n\n", key)
	bq.mu.Lock()
	defer bq.mu.Unlock()
	delete(bq.callbacks, key)
}

func (bq *BaseQueue) Count() int {
	log.Printf("checking count...\n")
	bq.mu.Lock()
	defer bq.mu.Unlock()
	return len(bq.queue)
}

func (bq *BaseQueue) First() string {
	log.Printf("checking first...\n")
	bq.mu.Lock()
	defer bq.mu.Unlock()
	if len(bq.queue) == 0 {
		return ""
	}
	return bq.queue[0]
}
