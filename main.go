package main

import (
	"fmt"
	"tourofgo/bitwiseOperations/flags"
	"tourofgo/bitwiseOperations/pixels"
	"tourofgo/bitwiseOperations/utils"
)

func taskIfEven(n int) {
	if utils.IsEven(n) {
		fmt.Println("task IfEven: Even")
	} else {
		fmt.Println("task IfEven: Odd")
	}
}

func taskBitEnable(num, k int) {
	fmt.Printf("Decimal: %d enable %d bit = %d\n", num, k, utils.BitEnable(num, k))
	fmt.Printf("Binary: %b enable %d bit = %b\n", num, k, utils.BitEnable(num, k))
}

func taskSwapXor(x, y int) {
	fmt.Printf("before swap(decimal): x = %d, y = %d\n", x, y)
	x, y = utils.SwapXor(x, y)
	fmt.Printf("after swap(decimal): x = %d, y = %d\n", x, y)
}

func taskOnesCount(x uint64) {
	fmt.Printf("Ones count in %d (%b) = %d\n", x, x, utils.OnesCount(x))

}

func main() {
	fmt.Println("----------------testing utils------------------")

	taskIfEven(7)

	taskBitEnable(89, 5)

	taskSwapXor(11, 12)

	taskOnesCount(29)

	fmt.Println("----------------testing flags------------------")

	s := []flags.Flags{flags.Read, flags.Write, flags.Execute, flags.Delete, flags.Share}
	var perms flags.Flags
	perms.Set(flags.Read)
	perms.Set(flags.Write)
	perms.Set(flags.Share)
	perms.Toggle()
	perms.Clear()
	for _, val := range s {
		if perms.Has(val) {
			fmt.Printf("user has %v permission\n", val)
		} else {
			fmt.Printf("user hasn't %v permission\n", val)
		}
	}
	fmt.Printf("perms = %08b\n", perms)

	fmt.Println("----------------testing pixels------------------")
	var pixel pixels.Pixel = 0x80FF9933
	a := pixels.Alpha(pixel)
	r := pixels.Red(pixel)
	g := pixels.Green(pixel)
	b := pixels.Blue(pixel)
	fmt.Printf("Alpha = %d, Red = %d, Green = %d, Blue = %d\n", a, r, g, b)
}
