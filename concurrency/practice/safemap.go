package practice

import (
	"sync"
	"sync/atomic"
)

type SafeMap struct {
	data    map[string]int
	counter int64
	mu      sync.RWMutex
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]int),
	}
}

func (safeMap *SafeMap) Set(key string, val int) {
	atomic.AddInt64(&safeMap.counter, 1)
	safeMap.mu.Lock()
	safeMap.data[key] = val
	safeMap.mu.Unlock()
}

func (safeMap *SafeMap) Get(key string) (int, bool) {
	atomic.AddInt64(&safeMap.counter, 1)
	safeMap.mu.RLock()
	defer safeMap.mu.RUnlock()
	val, ok := safeMap.data[key]
	return val, ok
}

func (safeMap *SafeMap) Delete(key string) {
	atomic.AddInt64(&safeMap.counter, 1)
	safeMap.mu.Lock()
	delete(safeMap.data, key)
	safeMap.mu.Unlock()
}

func (safeMap *SafeMap) Len() int {
	atomic.AddInt64(&safeMap.counter, 1)
	safeMap.mu.RLock()
	defer safeMap.mu.RUnlock()
	return len(safeMap.data)
}

func (safeMap *SafeMap) OpsCount() int64 {
	safeMap.mu.Lock()
	defer safeMap.mu.Unlock()
	return safeMap.counter
}
