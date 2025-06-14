package utils

import "math"

//Проверить четность числа
func IsEven(n int) bool {
	return n&1 == 0
}

//Включить бит по указанному индексу
func BitEnable(num, k int) int {
	return num | int(math.Pow(2, float64(k)))
}

//Поменять числа с помощью xor
func SwapXor(x, y int) (int, int) {
	x = x ^ y
	y = x ^ y
	x = x ^ y
	return x, y
}

//Подсчитать количество единиц и сравнить с реализацией math/bits.OnesCount64
func OnesCount(x uint64) int {
	count := 0
	for x != 0 {
		if x%2 == 1 {
			count++
		}
		x = x / 2
	}
	return count
}
