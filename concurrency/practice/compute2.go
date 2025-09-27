package practice

import (
	"fmt"
	"sync"
)

// Fan Out
func Compute2(id int, nums <-chan int, res chan<- int) {
	for num := range nums {
		res <- num * num
	}
}

func Execute2() {

	nums := make(chan int, 10)
	res := make(chan int, 10)

	even := make([]int, 0)
	odd := make([]int, 0)

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		nums <- i
	}

	close(nums)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			go Compute2(n, nums, res)
		}(i)
	}

	wg.Wait()
	close(res)

	for v := range res {
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
