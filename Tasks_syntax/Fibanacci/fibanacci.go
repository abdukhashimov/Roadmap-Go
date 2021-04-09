package main

import (
	"fmt"
)

func main() {
	numFib := fibNum(6)
	fmt.Println(numFib)

	numFib1 := FibRec(6)
	fmt.Println(numFib1)
}

func fibNum(x int) int {
	x1, x2 := 0, 1
	o1, o2 := &x1, &x2
	var sum int
	for i := 0; i < x; i++ {
		sum = *o1 + *o2
		*o2 = *o1
		*o1 = sum
	}
	return x1
}

func FibRec(n int) int {
	if n <= 1 {
		return n
	}
	return FibRec(n-1) + FibRec(n-2)
}