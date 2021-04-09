package main

import (
	"fmt"
)

func main() {
	arr := []string{"aziz", "aziz", "aj", "awer", "aj"}
	duble := duble(arr)
	fmt.Println(duble)

}

func duble(arr []string) []string {
	check := make(map[string]int)
	res := make([]string, 0)
	for _, val := range arr {
		check[val] = 1
	}
	for key := range check {
		res = append(res, key)
	}
	return res
}
