package practice

import (
	"fmt"
	"sync"
)

func Compute(num int, c chan int) {
	res := num * num
	c <- res
}

// concurent squaring and sorting by evenness using wg and channels
func Execute() {
	wg := sync.WaitGroup{}
	even := make([]int, 0)
	odd := make([]int, 0)
	ch := make(chan int, 100)
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			Compute(n, ch)
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
