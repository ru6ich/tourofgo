package practice

func Compute(num int, c chan int) {
	res := num * num
	c <- res
}
