package main

import (
	"fmt"
	"tourofgo/concurrency/practice"
)

func main() {
	nums := practice.Generator(10)
	primes := practice.Filter(nums)
	fmt.Println(practice.Summer(primes))
}
