package main

import (
	"fmt"
	"sync"
	"tourofgo/concurrency/practice"
)

func main() {
	wg := sync.WaitGroup{}
	even := make([]int, 0)
	odd := make([]int, 0)
	ch := make(chan int, 100)
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			go practice.Compute(n, ch)
		}(i)
	}
	wg.Wait()
	close(ch)
	for v := range ch {
		if v%2 == 0 {
			even = append(even, v)
		}
		if v%2 != 0 {
			odd = append(odd, v)
		}
	}
	fmt.Println(even)
	fmt.Println(odd)
}
