package main

import (
	"fmt"
	"sort"
)

func main() {
	testArrays4()
	// testSlice3()
}

func testArrays1() {
	var myarr [3]string
	myarr[0] = "GO"
	myarr[1] = "Java"
	myarr[2] = "c++"
	fmt.Println("Qatorning elementlari: ")
	fmt.Println("1: ", myarr[0])
	fmt.Println("2: ", myarr[1])
	fmt.Println("3: ", myarr[2])

}
func testArrays2() {
	myarr := [3]int{2, 5, 9}
	fmt.Println(myarr)
}
func testArrays3() {
	myarr1 := [...]int{2, 5, 9}
	myarr2 := [3]int{2, 5, 9}
	fmt.Println(myarr1 == myarr2)
}
func testArrays4() {
	myarr := [3][3]string{
		{"c++", "Go", "TS"},
		{"c", "Java", "Python"},
		{"c#", "Node", "JS"},
	}
	fmt.Println(myarr[1][1])
}
func testArrays5() {
	myarr1 := [3]int{2, 5, 9}
	myarr2 := myarr1
	fmt.Println(myarr1)
	fmt.Println(myarr2)
	myarr1[2] = 10
	fmt.Println(myarr1)
	fmt.Println(myarr2)

}
func testArrays6() {
	myarr1 := [3]int{2, 5, 9}
	myarr2 := &myarr1
	fmt.Println(myarr1)
	fmt.Println(*myarr2)
	myarr1[2] = 10
	fmt.Println(myarr1)
	fmt.Println(*myarr2)
}

// ===========================
// Slice =====================
// ===========================

func testSlice1() {
	myslice := []int{2, 4, 5}
	myslice = append(myslice, 10)
	fmt.Printf("slice'ning uzunligi: %d", len(myslice))
}

func testSlice2() {
	myarr := [7]string{"olma", "anor", "anjir", "Shaftoli", "Banan"}
	myslice := myarr[1:3]
	fmt.Println(myslice)
}
func testSlice3() {
	myslice := []int{2, 4, 5, 3, 8, 6}
	sort.Ints(myslice)
	fmt.Println(myslice)
}
