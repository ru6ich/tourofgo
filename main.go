package main

import (
	"fmt"
	"tourofgo/concurrency/practice"
)

func main() {
	safeMap := practice.NewSafeMap()
	alphabet := []string{"alpha", "beta", "gamma", "delta", "epsilon", "zeta", "eta", "tetta", "iota", "kappa"}

	// Write goroutine
	for w := 0; w < 5; w++ {
		go func(id int) {
			safeMap.Set(alphabet[id], id)
		}(w)
	}

	// Read Goroutine
	for w := 0; w < 5; w++ {
		go func(id int) {
			_, _ = safeMap.Get(alphabet[id])
		}(w)
	}

	for w := 0; w < 2; w++ {
		go func() {
			safeMap.Delete(alphabet[w])
		}()
	}

	fmt.Println(safeMap.Len())
	fmt.Println(safeMap.OpsCount())
}
