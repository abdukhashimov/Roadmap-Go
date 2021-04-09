package main

import (
	"fmt"
	"time"
)

func main() {
	testSwitch3()
}
func testIf() {
	i := 7
	if i == 7 {
		fmt.Println("yetti!")
	}
}
func testIfElse() {
	points := 100000
	if points < 500 {
		fmt.Println("Silver")
	} else if points < 10000 {
		fmt.Println("Gold")
	} else {
		fmt.Println("Platinum")
	}
}

func getTime() {
	weekday := time.Now().Day()
	fmt.Println(weekday)
}
func testSwitch1() {
	weekday := 1
	switch weekday {
	case 1:
		fmt.Println("Dushanba")
	case 2:
		fmt.Println("Seshanba")
	case 3:
		fmt.Println("Chorshanba")
	case 4:
		fmt.Println("Payshanba")
	case 5:
		fmt.Println("Juma")
	case 6:
		fmt.Println("Shanba")
	case 7:
		fmt.Println("Yakshanba")
	default:
		fmt.Println("Xato")
	}
}
func testSwitch2() {
	greeting := 1
	switch {
	case greeting == 1:
		fmt.Println("Dushanba")
	case greeting == 2:
		fmt.Println("Seshanba")
	case greeting == 3:
		fmt.Println("Chorshanba")
	case greeting == 4:
		fmt.Println("Payshanba")
	case greeting == 5:
		fmt.Println("Juma")
	case greeting == 6:
		fmt.Println("Shanba")
	case greeting == 7:
		fmt.Println("Yakshanba")
	default:
		fmt.Println("Xato")
	}
}

func testSwitch3() {
	var userChoice string = "ikki"
	switch userChoice {
	case "bir":
		fmt.Println("c#")
	case "ikki", "uch":
		fmt.Println("go")
	case "tort", "besh":
		fmt.Println("Java")
	}
}
