package main

import (
	"math"
	"math/bits"
	"testing"
	"tourofgo/bitwiseOperations/utils"
)

func BenchmarkOnesCountIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		utils.OnesCount(uint64(math.MaxUint64))
	}
}

func BenchmarkOnesCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bits.OnesCount64(uint64(math.MaxUint64))
	}
}
