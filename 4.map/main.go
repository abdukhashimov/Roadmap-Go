// map[KeyType]ValueType

package main

import "fmt"

func main() {
	statuses := make(map[string]int)
	statuses["Akmal"] = 0
	statuses["Akmal"] = 1
	statuses["Mamura"] = 2
	statuses["Alisher"] = 3
	statuses["Vali"] = 4
	// var ismlarStatus = statuses["Vali"]
	fmt.Println(statuses)
	// delete(statuses, "Mamura")

	// arr := []string{"aziz","aziz","aj","awer","aj"}
	// check := make(map[string]int)
	// res :=[]string{}
	// for i, val := range arr {
	// 	// fmt.Println(i ,val)
	// 	check[val] = i
	// }
	// fmt.Println(check)
	// // fmt.Println(res)
	// for k, v:= range check {
	// 	fmt.Println(k, v)
	// 	res = append(res,k)
	// }
	// fmt.Println(res)

}
