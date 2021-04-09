package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	loop4()
}
func loop1() {
	for i := 0; i < 8; i++ {
		fmt.Println(i)
	}
}
func loop2() {
	i := 0
	for i < 8 {
		fmt.Println(i)
		i++
	}
}
func loop3() {
	for {
		fmt.Println("Alhamdulillah" + strconv.Itoa(time.Now().Second()))
		time.Sleep(1 * time.Second)
	}
}
func loop4() {
	// myarr := [3]string{"yer", "quyosh", "oy"}
	mymap:=map[int]string{
		1: "c++",
		2: "c",
		3: "c#",
	}
	for i, v := range mymap {
		fmt.Println("index:", i, "value", v)
	}
}
