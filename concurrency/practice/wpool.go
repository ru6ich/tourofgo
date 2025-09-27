package practice

import (
	"fmt"
	"sync"
	"time"
)

func Execute6() {
	jobs := make(chan int)
	results := make(chan int, 100)
	wg := sync.WaitGroup{}

	for w := 0; w < 5; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case v, ok := <-jobs:
					if !ok {
						return
					}
					results <- v * v
				case <-time.After(1 * time.Second):
					return
				}
			}
		}()
	}

	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for v := range results {
		fmt.Println(v)
	}
}
