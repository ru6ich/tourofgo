package practice

import (
	"math/rand"
	"time"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	if n == 2 {
		return true
	}

	if n%2 == 0 {
		return false
	}

	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func shuffle(nums []int) []int {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	for i := len(nums) - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}

func Generator(n int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		nums := make([]int, n)
		for i := 0; i < n; i++ {
			nums[i] = i
		}

		nums = shuffle(nums)

		for _, v := range nums {
			out <- v
		}
	}()

	return out
}

func Filter(nums <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for v := range nums {
			if isPrime(v) {
				out <- v
			}
		}
	}()
	return out
}

func Summer(filtered <-chan int) int {
	sum := 0

	for v := range filtered {
		sum += v
	}

	return sum
}
