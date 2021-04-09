package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	odd_even(12345)
	odd_even2(12345)
}

func odd_even(x int) {
	var odd, even int
	str := strconv.Itoa(x)
	strArr := strings.Split(str, "")
	intArr := []int{}
	for _, val := range strArr {
		j, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println(err)
		}
		intArr = append(intArr, j)
	}
	for i, v := range intArr {
		if i%2 == 0 {
			odd += v
		} else {
			even += v
		}
	}
	fmt.Println(odd, even)
}

func odd_even2(n int) {
	var odd, even int
	i := 0
	for ; n > 0; i++ {
		digit := n % 10
		n /= 10

		if i%2 == 0 {
			odd += digit
		} else {
			even += digit
		}
	}
	fmt.Println(odd, even)
}
