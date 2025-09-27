package practice

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int64
	mu    sync.Mutex
}

func (counter *Counter) Inc() {
	counter.mu.Lock()
	counter.count++
	counter.mu.Unlock()
	// atomic.AddInt64(&counter.count, 1)
}

func (counter *Counter) Counted() int64 {
	counter.mu.Lock()
	defer counter.mu.Unlock()
	return counter.count
	// return atomic.LoadInt64(&counter.count)
}

func Execute5() {
	wg := sync.WaitGroup{}
	counter := Counter{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}
	wg.Wait()
	fmt.Println(counter.Counted())
}
