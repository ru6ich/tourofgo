package bitwiseOperations

import (
	"fmt"
	"math/bits"
)

func Operations() {
	var a, b uint8 = 0b1100_1010, 0b0110_1100

	fmt.Printf("%08b & %08b = %08b\n", a, b, a&b)
	fmt.Printf("%08b | %08b = %08b\n", a, b, a|b)
	fmt.Printf("%08b ^ %08b = %08b\n", a, b, a^b)
	fmt.Printf("^%08b = %08b\n", a, ^a)
	fmt.Printf("%08b &^ %08b = %08b\n", a, b, a&^b)
	fmt.Printf("%08b << 3 = %08b\n", a, a<<3)
	fmt.Printf("%08b >> 2 = %08b\n", a, a>>2)

	fmt.Println("Popcount a = ", bits.OnesCount8(a))
}
