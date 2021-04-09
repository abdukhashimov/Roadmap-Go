package main

import (
	"fmt"
)

func main() {

	pm := Palindrome("aziza aziza")
	fmt.Println(pm)

}

func Palindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] == input[len(input)-i-1] {
			return true
		}
	}
	return false
}

